// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcase"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcaseset"
)

// TestcaseSetCreate is the builder for creating a TestcaseSet entity.
type TestcaseSetCreate struct {
	config
	mutation *TestcaseSetMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tsc *TestcaseSetCreate) SetName(s string) *TestcaseSetCreate {
	tsc.mutation.SetName(s)
	return tsc
}

// SetScore sets the "score" field.
func (tsc *TestcaseSetCreate) SetScore(i int) *TestcaseSetCreate {
	tsc.mutation.SetScore(i)
	return tsc
}

// SetIsSample sets the "is_sample" field.
func (tsc *TestcaseSetCreate) SetIsSample(b bool) *TestcaseSetCreate {
	tsc.mutation.SetIsSample(b)
	return tsc
}

// SetCreatedAt sets the "created_at" field.
func (tsc *TestcaseSetCreate) SetCreatedAt(t time.Time) *TestcaseSetCreate {
	tsc.mutation.SetCreatedAt(t)
	return tsc
}

// SetUpdatedAt sets the "updated_at" field.
func (tsc *TestcaseSetCreate) SetUpdatedAt(t time.Time) *TestcaseSetCreate {
	tsc.mutation.SetUpdatedAt(t)
	return tsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tsc *TestcaseSetCreate) SetNillableUpdatedAt(t *time.Time) *TestcaseSetCreate {
	if t != nil {
		tsc.SetUpdatedAt(*t)
	}
	return tsc
}

// SetID sets the "id" field.
func (tsc *TestcaseSetCreate) SetID(i int) *TestcaseSetCreate {
	tsc.mutation.SetID(i)
	return tsc
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (tsc *TestcaseSetCreate) SetTaskID(id int) *TestcaseSetCreate {
	tsc.mutation.SetTaskID(id)
	return tsc
}

// SetNillableTaskID sets the "task" edge to the Task entity by ID if the given value is not nil.
func (tsc *TestcaseSetCreate) SetNillableTaskID(id *int) *TestcaseSetCreate {
	if id != nil {
		tsc = tsc.SetTaskID(*id)
	}
	return tsc
}

// SetTask sets the "task" edge to the Task entity.
func (tsc *TestcaseSetCreate) SetTask(t *Task) *TestcaseSetCreate {
	return tsc.SetTaskID(t.ID)
}

// AddTestcaseIDs adds the "testcases" edge to the Testcase entity by IDs.
func (tsc *TestcaseSetCreate) AddTestcaseIDs(ids ...int) *TestcaseSetCreate {
	tsc.mutation.AddTestcaseIDs(ids...)
	return tsc
}

// AddTestcases adds the "testcases" edges to the Testcase entity.
func (tsc *TestcaseSetCreate) AddTestcases(t ...*Testcase) *TestcaseSetCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tsc.AddTestcaseIDs(ids...)
}

// Mutation returns the TestcaseSetMutation object of the builder.
func (tsc *TestcaseSetCreate) Mutation() *TestcaseSetMutation {
	return tsc.mutation
}

// Save creates the TestcaseSet in the database.
func (tsc *TestcaseSetCreate) Save(ctx context.Context) (*TestcaseSet, error) {
	return withHooks(ctx, tsc.sqlSave, tsc.mutation, tsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tsc *TestcaseSetCreate) SaveX(ctx context.Context) *TestcaseSet {
	v, err := tsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tsc *TestcaseSetCreate) Exec(ctx context.Context) error {
	_, err := tsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsc *TestcaseSetCreate) ExecX(ctx context.Context) {
	if err := tsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tsc *TestcaseSetCreate) check() error {
	if _, ok := tsc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "TestcaseSet.name"`)}
	}
	if _, ok := tsc.mutation.Score(); !ok {
		return &ValidationError{Name: "score", err: errors.New(`ent: missing required field "TestcaseSet.score"`)}
	}
	if _, ok := tsc.mutation.IsSample(); !ok {
		return &ValidationError{Name: "is_sample", err: errors.New(`ent: missing required field "TestcaseSet.is_sample"`)}
	}
	if _, ok := tsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TestcaseSet.created_at"`)}
	}
	return nil
}

func (tsc *TestcaseSetCreate) sqlSave(ctx context.Context) (*TestcaseSet, error) {
	if err := tsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	tsc.mutation.id = &_node.ID
	tsc.mutation.done = true
	return _node, nil
}

func (tsc *TestcaseSetCreate) createSpec() (*TestcaseSet, *sqlgraph.CreateSpec) {
	var (
		_node = &TestcaseSet{config: tsc.config}
		_spec = sqlgraph.NewCreateSpec(testcaseset.Table, sqlgraph.NewFieldSpec(testcaseset.FieldID, field.TypeInt))
	)
	if id, ok := tsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tsc.mutation.Name(); ok {
		_spec.SetField(testcaseset.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tsc.mutation.Score(); ok {
		_spec.SetField(testcaseset.FieldScore, field.TypeInt, value)
		_node.Score = value
	}
	if value, ok := tsc.mutation.IsSample(); ok {
		_spec.SetField(testcaseset.FieldIsSample, field.TypeBool, value)
		_node.IsSample = value
	}
	if value, ok := tsc.mutation.CreatedAt(); ok {
		_spec.SetField(testcaseset.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tsc.mutation.UpdatedAt(); ok {
		_spec.SetField(testcaseset.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if nodes := tsc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   testcaseset.TaskTable,
			Columns: []string{testcaseset.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.task_testcase_sets = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tsc.mutation.TestcasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   testcaseset.TestcasesTable,
			Columns: testcaseset.TestcasesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcase.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TestcaseSetCreateBulk is the builder for creating many TestcaseSet entities in bulk.
type TestcaseSetCreateBulk struct {
	config
	builders []*TestcaseSetCreate
}

// Save creates the TestcaseSet entities in the database.
func (tscb *TestcaseSetCreateBulk) Save(ctx context.Context) ([]*TestcaseSet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tscb.builders))
	nodes := make([]*TestcaseSet, len(tscb.builders))
	mutators := make([]Mutator, len(tscb.builders))
	for i := range tscb.builders {
		func(i int, root context.Context) {
			builder := tscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TestcaseSetMutation)
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
					_, err = mutators[i+1].Mutate(root, tscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tscb *TestcaseSetCreateBulk) SaveX(ctx context.Context) []*TestcaseSet {
	v, err := tscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tscb *TestcaseSetCreateBulk) Exec(ctx context.Context) error {
	_, err := tscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tscb *TestcaseSetCreateBulk) ExecX(ctx context.Context) {
	if err := tscb.Exec(ctx); err != nil {
		panic(err)
	}
}
