// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contestuser"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/predicate"
)

// ContestUserDelete is the builder for deleting a ContestUser entity.
type ContestUserDelete struct {
	config
	hooks    []Hook
	mutation *ContestUserMutation
}

// Where appends a list predicates to the ContestUserDelete builder.
func (cud *ContestUserDelete) Where(ps ...predicate.ContestUser) *ContestUserDelete {
	cud.mutation.Where(ps...)
	return cud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cud *ContestUserDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cud.sqlExec, cud.mutation, cud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cud *ContestUserDelete) ExecX(ctx context.Context) int {
	n, err := cud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cud *ContestUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(contestuser.Table, sqlgraph.NewFieldSpec(contestuser.FieldID, field.TypeInt))
	if ps := cud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cud.mutation.done = true
	return affected, err
}

// ContestUserDeleteOne is the builder for deleting a single ContestUser entity.
type ContestUserDeleteOne struct {
	cud *ContestUserDelete
}

// Where appends a list predicates to the ContestUserDelete builder.
func (cudo *ContestUserDeleteOne) Where(ps ...predicate.ContestUser) *ContestUserDeleteOne {
	cudo.cud.mutation.Where(ps...)
	return cudo
}

// Exec executes the deletion query.
func (cudo *ContestUserDeleteOne) Exec(ctx context.Context) error {
	n, err := cudo.cud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{contestuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cudo *ContestUserDeleteOne) ExecX(ctx context.Context) {
	if err := cudo.Exec(ctx); err != nil {
		panic(err)
	}
}