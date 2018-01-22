// Copyright 2016 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package distsqlrun

import (
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
	"github.com/cockroachdb/cockroach/pkg/storage/engine"
)

// sorter sorts the input rows according to the specified ordering.
type sorterBase struct {
	processorBase

	evalCtx *tree.EvalContext

	input    RowSource
	ordering sqlbase.ColumnOrdering
	matchLen uint32
	// count is the maximum number of rows that the sorter will push to the
	// ProcOutputHelper. 0 if the sorter should sort and push all the rows from
	// the input.
	count int64
	// tempStorage is used to store rows when the working set is larger than can
	// be stored in memory.
	tempStorage engine.Engine
}

func newSorterBase(
	flowCtx *FlowCtx, spec *SorterSpec, input RowSource, post *PostProcessSpec, output RowReceiver,
) (*sorterBase, error) {
	count := int64(0)
	if post.Limit != 0 {
		// The sorter needs to produce Offset + Limit rows. The ProcOutputHelper
		// will discard the first Offset ones.
		count = int64(post.Limit) + int64(post.Offset)
	}

	s := &sorterBase{
		input:       input,
		ordering:    convertToColumnOrdering(spec.OutputOrdering),
		matchLen:    spec.OrderingMatchLen,
		count:       count,
		tempStorage: flowCtx.TempStorage,
		evalCtx:     flowCtx.NewEvalCtx(),
	}
	if err := s.init(post, input.OutputTypes(), flowCtx, s.evalCtx, output); err != nil {
		return nil, err
	}
	return s, nil
}

func newSorter(
	flowCtx *FlowCtx, spec *SorterSpec, input RowSource, post *PostProcessSpec, output RowReceiver,
) (Processor, error) {
	s, err := newSorterBase(flowCtx, spec, input, post, output)
	if err != nil {
		return nil, err
	}

	if s.matchLen == 0 {
		if s.count == 0 {
			// No specified ordering match length and unspecified limit; no
			// optimizations are possible so we simply load all rows into memory and
			// sort all values in-place. It has a worst-case time complexity of
			// O(n*log(n)) and a worst-case space complexity of O(n).
			return newSortAllStrategy(s), nil
		}
		// No specified ordering match length but specified limit; we can optimize
		// our sort procedure by maintaining a max-heap populated with only the
		// smallest k rows seen. It has a worst-case time complexity of
		// O(n*log(k)) and a worst-case space complexity of O(k).
		return newSortTopKStrategy(s, s.count), nil
	}
	// Ordering match length is specified. We will be able to use existing
	// ordering in order to avoid loading all the rows into memory. If we're
	// scanning an index with a prefix matching an ordering prefix, we can only
	// accumulate values for equal fields in this prefix, sort the accumulated
	// chunk and then output.
	// TODO(irfansharif): Add optimization for case where both ordering match
	// length and limit is specified.
	return newSortChunksStrategy(s), nil

}
