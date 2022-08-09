// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package opgen

import (
	"github.com/cockroachdb/cockroach/pkg/sql/schemachanger/scop"
	"github.com/cockroachdb/cockroach/pkg/sql/schemachanger/scpb"
)

func init() {
	opRegistry.register((*scpb.CheckConstraint)(nil),
		toPublic(
			scpb.Status_ABSENT,
			to(scpb.Status_VALIDATING,
				emit(func(this *scpb.CheckConstraint) *scop.AddCheckConstraint {
					return &scop.AddCheckConstraint{
						TableID:      this.TableID,
						ConstraintID: this.ConstraintID,
						ColumnIDs:    this.ColumnIDs,
						Expression:   this.Expression,
					}
				}),
			),
			to(scpb.Status_VALIDATED,
				emit(func(this *scpb.CheckConstraint) *scop.ValidateCheckConstraint {
					return &scop.ValidateCheckConstraint{
						TableID:      this.TableID,
						ConstraintID: this.ConstraintID,
					}
				}),
			),
			to(scpb.Status_PUBLIC,
				emit(func(this *scpb.CheckConstraint) *scop.MakeAddedCheckConstraintPublic {
					return &scop.MakeAddedCheckConstraintPublic{
						TableID:      this.TableID,
						ConstraintID: this.ConstraintID,
					}
				}),
			),
		),
		toAbsent(
			scpb.Status_PUBLIC,
			to(scpb.Status_VALIDATED,
				emit(func(this *scpb.CheckConstraint) *scop.MakeDroppedCheckConstraintValidated {
					return &scop.MakeDroppedCheckConstraintValidated{
						TableID:      this.TableID,
						ConstraintID: this.ConstraintID,
					}
				})),
			equiv(scpb.Status_VALIDATING),
			to(scpb.Status_ABSENT,
				// TODO(postamar): remove revertibility constraint when possible
				revertible(false),
				emit(func(this *scpb.CheckConstraint) *scop.RemoveCheckConstraint {
					return &scop.RemoveCheckConstraint{
						TableID:      this.TableID,
						ConstraintID: this.ConstraintID,
					}
				}),
				emit(func(this *scpb.CheckConstraint) *scop.UpdateTableBackReferencesInTypes {
					if len(this.UsesTypeIDs) == 0 {
						return nil
					}
					return &scop.UpdateTableBackReferencesInTypes{
						TypeIDs:               this.UsesTypeIDs,
						BackReferencedTableID: this.TableID,
					}
				}),
				emit(func(this *scpb.CheckConstraint) *scop.UpdateBackReferencesInSequences {
					if len(this.UsesSequenceIDs) == 0 {
						return nil
					}
					return &scop.UpdateBackReferencesInSequences{
						SequenceIDs:           this.UsesSequenceIDs,
						BackReferencedTableID: this.TableID,
					}
				}),
			),
		),
	)
}
