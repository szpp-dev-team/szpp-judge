// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/predicate"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcaseresult"
)

// TestcaseResultDelete is the builder for deleting a TestcaseResult entity.
type TestcaseResultDelete struct {
	config
	hooks    []Hook
	mutation *TestcaseResultMutation
}

// Where appends a list predicates to the TestcaseResultDelete builder.
func (trd *TestcaseResultDelete) Where(ps ...predicate.TestcaseResult) *TestcaseResultDelete {
	trd.mutation.Where(ps...)
	return trd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (trd *TestcaseResultDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, trd.sqlExec, trd.mutation, trd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (trd *TestcaseResultDelete) ExecX(ctx context.Context) int {
	n, err := trd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (trd *TestcaseResultDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(testcaseresult.Table, sqlgraph.NewFieldSpec(testcaseresult.FieldID, field.TypeInt))
	if ps := trd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, trd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	trd.mutation.done = true
	return affected, err
}

// TestcaseResultDeleteOne is the builder for deleting a single TestcaseResult entity.
type TestcaseResultDeleteOne struct {
	trd *TestcaseResultDelete
}

// Where appends a list predicates to the TestcaseResultDelete builder.
func (trdo *TestcaseResultDeleteOne) Where(ps ...predicate.TestcaseResult) *TestcaseResultDeleteOne {
	trdo.trd.mutation.Where(ps...)
	return trdo
}

// Exec executes the deletion query.
func (trdo *TestcaseResultDeleteOne) Exec(ctx context.Context) error {
	n, err := trdo.trd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{testcaseresult.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (trdo *TestcaseResultDeleteOne) ExecX(ctx context.Context) {
	if err := trdo.Exec(ctx); err != nil {
		panic(err)
	}
}