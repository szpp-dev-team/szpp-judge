// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcase"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcaseset"
)

// TestcaseCreate is the builder for creating a Testcase entity.
type TestcaseCreate struct {
	config
	mutation *TestcaseMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *TestcaseCreate) SetName(s string) *TestcaseCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetDescription sets the "description" field.
func (tc *TestcaseCreate) SetDescription(s string) *TestcaseCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tc *TestcaseCreate) SetNillableDescription(s *string) *TestcaseCreate {
	if s != nil {
		tc.SetDescription(*s)
	}
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TestcaseCreate) SetCreatedAt(t time.Time) *TestcaseCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TestcaseCreate) SetUpdatedAt(t time.Time) *TestcaseCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TestcaseCreate) SetNillableUpdatedAt(t *time.Time) *TestcaseCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TestcaseCreate) SetID(i int) *TestcaseCreate {
	tc.mutation.SetID(i)
	return tc
}

// AddTestcaseSetIDs adds the "testcase_sets" edge to the TestcaseSet entity by IDs.
func (tc *TestcaseCreate) AddTestcaseSetIDs(ids ...int) *TestcaseCreate {
	tc.mutation.AddTestcaseSetIDs(ids...)
	return tc
}

// AddTestcaseSets adds the "testcase_sets" edges to the TestcaseSet entity.
func (tc *TestcaseCreate) AddTestcaseSets(t ...*TestcaseSet) *TestcaseCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTestcaseSetIDs(ids...)
}

// Mutation returns the TestcaseMutation object of the builder.
func (tc *TestcaseCreate) Mutation() *TestcaseMutation {
	return tc.mutation
}

// Save creates the Testcase in the database.
func (tc *TestcaseCreate) Save(ctx context.Context) (*Testcase, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TestcaseCreate) SaveX(ctx context.Context) *Testcase {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TestcaseCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TestcaseCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TestcaseCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Testcase.name"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Testcase.created_at"`)}
	}
	return nil
}

func (tc *TestcaseCreate) sqlSave(ctx context.Context) (*Testcase, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TestcaseCreate) createSpec() (*Testcase, *sqlgraph.CreateSpec) {
	var (
		_node = &Testcase{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(testcase.Table, sqlgraph.NewFieldSpec(testcase.FieldID, field.TypeInt))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(testcase.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(testcase.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(testcase.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(testcase.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if nodes := tc.mutation.TestcaseSetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   testcase.TestcaseSetsTable,
			Columns: testcase.TestcaseSetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcaseset.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TestcaseCreateBulk is the builder for creating many Testcase entities in bulk.
type TestcaseCreateBulk struct {
	config
	builders []*TestcaseCreate
}

// Save creates the Testcase entities in the database.
func (tcb *TestcaseCreateBulk) Save(ctx context.Context) ([]*Testcase, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Testcase, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TestcaseMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TestcaseCreateBulk) SaveX(ctx context.Context) []*Testcase {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TestcaseCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TestcaseCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
