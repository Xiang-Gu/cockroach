// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// Code generated by generate_visitor.go. DO NOT EDIT.

package scop

import "context"

// MutationOp is an operation which can be visited by MutationVisitor.
type MutationOp interface {
	Op
	Visit(context.Context, MutationVisitor) error
}

// MutationVisitor is a visitor for MutationOp operations.
type MutationVisitor interface {
	NotImplemented(context.Context, NotImplemented) error
	MakeAbsentTempIndexDeleteOnly(context.Context, MakeAbsentTempIndexDeleteOnly) error
	MakeAbsentIndexBackfilling(context.Context, MakeAbsentIndexBackfilling) error
	SetAddedIndexPartialPredicate(context.Context, SetAddedIndexPartialPredicate) error
	MakeDeleteOnlyIndexWriteOnly(context.Context, MakeDeleteOnlyIndexWriteOnly) error
	MakeBackfilledIndexMerging(context.Context, MakeBackfilledIndexMerging) error
	MakeMergedIndexWriteOnly(context.Context, MakeMergedIndexWriteOnly) error
	MakeBackfillingIndexDeleteOnly(context.Context, MakeBackfillingIndexDeleteOnly) error
	MakeValidatedSecondaryIndexPublic(context.Context, MakeValidatedSecondaryIndexPublic) error
	MakeValidatedPrimaryIndexPublic(context.Context, MakeValidatedPrimaryIndexPublic) error
	MakePublicPrimaryIndexWriteOnly(context.Context, MakePublicPrimaryIndexWriteOnly) error
	CreateGCJobForTable(context.Context, CreateGCJobForTable) error
	CreateGCJobForDatabase(context.Context, CreateGCJobForDatabase) error
	CreateGCJobForIndex(context.Context, CreateGCJobForIndex) error
	MarkDescriptorAsPublic(context.Context, MarkDescriptorAsPublic) error
	MarkDescriptorAsSyntheticallyDropped(context.Context, MarkDescriptorAsSyntheticallyDropped) error
	MarkDescriptorAsDropped(context.Context, MarkDescriptorAsDropped) error
	DrainDescriptorName(context.Context, DrainDescriptorName) error
	MakeDeleteOnlyColumnWriteOnly(context.Context, MakeDeleteOnlyColumnWriteOnly) error
	MakePublicSecondaryIndexWriteOnly(context.Context, MakePublicSecondaryIndexWriteOnly) error
	MakeWriteOnlyIndexDeleteOnly(context.Context, MakeWriteOnlyIndexDeleteOnly) error
	RemoveDroppedIndexPartialPredicate(context.Context, RemoveDroppedIndexPartialPredicate) error
	MakeIndexAbsent(context.Context, MakeIndexAbsent) error
	MakeAbsentColumnDeleteOnly(context.Context, MakeAbsentColumnDeleteOnly) error
	SetAddedColumnType(context.Context, SetAddedColumnType) error
	MakeWriteOnlyColumnPublic(context.Context, MakeWriteOnlyColumnPublic) error
	MakePublicColumnWriteOnly(context.Context, MakePublicColumnWriteOnly) error
	MakeWriteOnlyColumnDeleteOnly(context.Context, MakeWriteOnlyColumnDeleteOnly) error
	RemoveDroppedColumnType(context.Context, RemoveDroppedColumnType) error
	MakeDeleteOnlyColumnAbsent(context.Context, MakeDeleteOnlyColumnAbsent) error
	RemoveOwnerBackReferenceInSequence(context.Context, RemoveOwnerBackReferenceInSequence) error
	RemoveSequenceOwner(context.Context, RemoveSequenceOwner) error
	RemoveCheckConstraint(context.Context, RemoveCheckConstraint) error
	MakeAbsentCheckConstraintWriteOnly(context.Context, MakeAbsentCheckConstraintWriteOnly) error
	MakePublicCheckConstraintValidated(context.Context, MakePublicCheckConstraintValidated) error
	MakeValidatedCheckConstraintPublic(context.Context, MakeValidatedCheckConstraintPublic) error
	MakeAbsentForeignKeyConstraintWriteOnly(context.Context, MakeAbsentForeignKeyConstraintWriteOnly) error
	MakeValidatedForeignKeyConstraintPublic(context.Context, MakeValidatedForeignKeyConstraintPublic) error
	MakePublicForeignKeyConstraintValidated(context.Context, MakePublicForeignKeyConstraintValidated) error
	RemoveForeignKeyConstraint(context.Context, RemoveForeignKeyConstraint) error
	RemoveForeignKeyBackReference(context.Context, RemoveForeignKeyBackReference) error
	MakeAbsentUniqueWithoutIndexConstraintWriteOnly(context.Context, MakeAbsentUniqueWithoutIndexConstraintWriteOnly) error
	MakeValidatedUniqueWithoutIndexConstraintPublic(context.Context, MakeValidatedUniqueWithoutIndexConstraintPublic) error
	MakePublicUniqueWithoutIndexConstraintValidated(context.Context, MakePublicUniqueWithoutIndexConstraintValidated) error
	RemoveUniqueWithoutIndexConstraint(context.Context, RemoveUniqueWithoutIndexConstraint) error
	RemoveSchemaParent(context.Context, RemoveSchemaParent) error
	AddIndexPartitionInfo(context.Context, AddIndexPartitionInfo) error
	LogEvent(context.Context, LogEvent) error
	AddColumnFamily(context.Context, AddColumnFamily) error
	AddColumnDefaultExpression(context.Context, AddColumnDefaultExpression) error
	RemoveColumnDefaultExpression(context.Context, RemoveColumnDefaultExpression) error
	AddColumnOnUpdateExpression(context.Context, AddColumnOnUpdateExpression) error
	RemoveColumnOnUpdateExpression(context.Context, RemoveColumnOnUpdateExpression) error
	UpdateTableBackReferencesInTypes(context.Context, UpdateTableBackReferencesInTypes) error
	RemoveBackReferenceInTypes(context.Context, RemoveBackReferenceInTypes) error
	UpdateBackReferencesInSequences(context.Context, UpdateBackReferencesInSequences) error
	RemoveViewBackReferencesInRelations(context.Context, RemoveViewBackReferencesInRelations) error
	SetColumnName(context.Context, SetColumnName) error
	SetIndexName(context.Context, SetIndexName) error
	SetConstraintName(context.Context, SetConstraintName) error
	DeleteDescriptor(context.Context, DeleteDescriptor) error
	RemoveJobStateFromDescriptor(context.Context, RemoveJobStateFromDescriptor) error
	SetJobStateOnDescriptor(context.Context, SetJobStateOnDescriptor) error
	UpdateSchemaChangerJob(context.Context, UpdateSchemaChangerJob) error
	CreateSchemaChangerJob(context.Context, CreateSchemaChangerJob) error
	UpsertTableComment(context.Context, UpsertTableComment) error
	RemoveAllTableComments(context.Context, RemoveAllTableComments) error
	RemoveTableComment(context.Context, RemoveTableComment) error
	UpsertDatabaseComment(context.Context, UpsertDatabaseComment) error
	RemoveDatabaseComment(context.Context, RemoveDatabaseComment) error
	UpsertSchemaComment(context.Context, UpsertSchemaComment) error
	RemoveSchemaComment(context.Context, RemoveSchemaComment) error
	UpsertIndexComment(context.Context, UpsertIndexComment) error
	RemoveIndexComment(context.Context, RemoveIndexComment) error
	UpsertColumnComment(context.Context, UpsertColumnComment) error
	RemoveColumnComment(context.Context, RemoveColumnComment) error
	UpsertConstraintComment(context.Context, UpsertConstraintComment) error
	RemoveConstraintComment(context.Context, RemoveConstraintComment) error
	RemoveDatabaseRoleSettings(context.Context, RemoveDatabaseRoleSettings) error
	RemoveUserPrivileges(context.Context, RemoveUserPrivileges) error
	DeleteSchedule(context.Context, DeleteSchedule) error
	RefreshStats(context.Context, RefreshStats) error
	AddColumnToIndex(context.Context, AddColumnToIndex) error
	RemoveColumnFromIndex(context.Context, RemoveColumnFromIndex) error
}

// Visit is part of the MutationOp interface.
func (op NotImplemented) Visit(ctx context.Context, v MutationVisitor) error {
	return v.NotImplemented(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeAbsentTempIndexDeleteOnly) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeAbsentTempIndexDeleteOnly(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeAbsentIndexBackfilling) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeAbsentIndexBackfilling(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op SetAddedIndexPartialPredicate) Visit(ctx context.Context, v MutationVisitor) error {
	return v.SetAddedIndexPartialPredicate(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeDeleteOnlyIndexWriteOnly) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeDeleteOnlyIndexWriteOnly(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeBackfilledIndexMerging) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeBackfilledIndexMerging(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeMergedIndexWriteOnly) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeMergedIndexWriteOnly(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeBackfillingIndexDeleteOnly) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeBackfillingIndexDeleteOnly(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeValidatedSecondaryIndexPublic) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeValidatedSecondaryIndexPublic(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeValidatedPrimaryIndexPublic) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeValidatedPrimaryIndexPublic(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakePublicPrimaryIndexWriteOnly) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakePublicPrimaryIndexWriteOnly(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op CreateGCJobForTable) Visit(ctx context.Context, v MutationVisitor) error {
	return v.CreateGCJobForTable(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op CreateGCJobForDatabase) Visit(ctx context.Context, v MutationVisitor) error {
	return v.CreateGCJobForDatabase(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op CreateGCJobForIndex) Visit(ctx context.Context, v MutationVisitor) error {
	return v.CreateGCJobForIndex(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MarkDescriptorAsPublic) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MarkDescriptorAsPublic(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MarkDescriptorAsSyntheticallyDropped) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MarkDescriptorAsSyntheticallyDropped(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MarkDescriptorAsDropped) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MarkDescriptorAsDropped(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op DrainDescriptorName) Visit(ctx context.Context, v MutationVisitor) error {
	return v.DrainDescriptorName(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeDeleteOnlyColumnWriteOnly) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeDeleteOnlyColumnWriteOnly(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakePublicSecondaryIndexWriteOnly) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakePublicSecondaryIndexWriteOnly(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeWriteOnlyIndexDeleteOnly) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeWriteOnlyIndexDeleteOnly(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveDroppedIndexPartialPredicate) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveDroppedIndexPartialPredicate(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeIndexAbsent) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeIndexAbsent(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeAbsentColumnDeleteOnly) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeAbsentColumnDeleteOnly(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op SetAddedColumnType) Visit(ctx context.Context, v MutationVisitor) error {
	return v.SetAddedColumnType(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeWriteOnlyColumnPublic) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeWriteOnlyColumnPublic(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakePublicColumnWriteOnly) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakePublicColumnWriteOnly(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeWriteOnlyColumnDeleteOnly) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeWriteOnlyColumnDeleteOnly(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveDroppedColumnType) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveDroppedColumnType(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeDeleteOnlyColumnAbsent) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeDeleteOnlyColumnAbsent(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveOwnerBackReferenceInSequence) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveOwnerBackReferenceInSequence(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveSequenceOwner) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveSequenceOwner(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveCheckConstraint) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveCheckConstraint(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeAbsentCheckConstraintWriteOnly) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeAbsentCheckConstraintWriteOnly(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakePublicCheckConstraintValidated) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakePublicCheckConstraintValidated(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeValidatedCheckConstraintPublic) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeValidatedCheckConstraintPublic(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeAbsentForeignKeyConstraintWriteOnly) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeAbsentForeignKeyConstraintWriteOnly(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeValidatedForeignKeyConstraintPublic) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeValidatedForeignKeyConstraintPublic(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakePublicForeignKeyConstraintValidated) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakePublicForeignKeyConstraintValidated(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveForeignKeyConstraint) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveForeignKeyConstraint(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveForeignKeyBackReference) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveForeignKeyBackReference(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeAbsentUniqueWithoutIndexConstraintWriteOnly) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeAbsentUniqueWithoutIndexConstraintWriteOnly(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakeValidatedUniqueWithoutIndexConstraintPublic) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakeValidatedUniqueWithoutIndexConstraintPublic(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op MakePublicUniqueWithoutIndexConstraintValidated) Visit(ctx context.Context, v MutationVisitor) error {
	return v.MakePublicUniqueWithoutIndexConstraintValidated(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveUniqueWithoutIndexConstraint) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveUniqueWithoutIndexConstraint(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveSchemaParent) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveSchemaParent(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op AddIndexPartitionInfo) Visit(ctx context.Context, v MutationVisitor) error {
	return v.AddIndexPartitionInfo(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op LogEvent) Visit(ctx context.Context, v MutationVisitor) error {
	return v.LogEvent(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op AddColumnFamily) Visit(ctx context.Context, v MutationVisitor) error {
	return v.AddColumnFamily(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op AddColumnDefaultExpression) Visit(ctx context.Context, v MutationVisitor) error {
	return v.AddColumnDefaultExpression(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveColumnDefaultExpression) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveColumnDefaultExpression(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op AddColumnOnUpdateExpression) Visit(ctx context.Context, v MutationVisitor) error {
	return v.AddColumnOnUpdateExpression(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveColumnOnUpdateExpression) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveColumnOnUpdateExpression(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op UpdateTableBackReferencesInTypes) Visit(ctx context.Context, v MutationVisitor) error {
	return v.UpdateTableBackReferencesInTypes(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveBackReferenceInTypes) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveBackReferenceInTypes(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op UpdateBackReferencesInSequences) Visit(ctx context.Context, v MutationVisitor) error {
	return v.UpdateBackReferencesInSequences(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveViewBackReferencesInRelations) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveViewBackReferencesInRelations(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op SetColumnName) Visit(ctx context.Context, v MutationVisitor) error {
	return v.SetColumnName(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op SetIndexName) Visit(ctx context.Context, v MutationVisitor) error {
	return v.SetIndexName(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op SetConstraintName) Visit(ctx context.Context, v MutationVisitor) error {
	return v.SetConstraintName(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op DeleteDescriptor) Visit(ctx context.Context, v MutationVisitor) error {
	return v.DeleteDescriptor(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveJobStateFromDescriptor) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveJobStateFromDescriptor(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op SetJobStateOnDescriptor) Visit(ctx context.Context, v MutationVisitor) error {
	return v.SetJobStateOnDescriptor(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op UpdateSchemaChangerJob) Visit(ctx context.Context, v MutationVisitor) error {
	return v.UpdateSchemaChangerJob(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op CreateSchemaChangerJob) Visit(ctx context.Context, v MutationVisitor) error {
	return v.CreateSchemaChangerJob(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op UpsertTableComment) Visit(ctx context.Context, v MutationVisitor) error {
	return v.UpsertTableComment(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveAllTableComments) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveAllTableComments(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveTableComment) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveTableComment(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op UpsertDatabaseComment) Visit(ctx context.Context, v MutationVisitor) error {
	return v.UpsertDatabaseComment(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveDatabaseComment) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveDatabaseComment(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op UpsertSchemaComment) Visit(ctx context.Context, v MutationVisitor) error {
	return v.UpsertSchemaComment(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveSchemaComment) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveSchemaComment(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op UpsertIndexComment) Visit(ctx context.Context, v MutationVisitor) error {
	return v.UpsertIndexComment(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveIndexComment) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveIndexComment(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op UpsertColumnComment) Visit(ctx context.Context, v MutationVisitor) error {
	return v.UpsertColumnComment(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveColumnComment) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveColumnComment(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op UpsertConstraintComment) Visit(ctx context.Context, v MutationVisitor) error {
	return v.UpsertConstraintComment(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveConstraintComment) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveConstraintComment(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveDatabaseRoleSettings) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveDatabaseRoleSettings(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveUserPrivileges) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveUserPrivileges(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op DeleteSchedule) Visit(ctx context.Context, v MutationVisitor) error {
	return v.DeleteSchedule(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RefreshStats) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RefreshStats(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op AddColumnToIndex) Visit(ctx context.Context, v MutationVisitor) error {
	return v.AddColumnToIndex(ctx, op)
}

// Visit is part of the MutationOp interface.
func (op RemoveColumnFromIndex) Visit(ctx context.Context, v MutationVisitor) error {
	return v.RemoveColumnFromIndex(ctx, op)
}
