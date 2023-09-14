// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcase"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcaseset"
)

// TestcaseCreate is the builder for creating a Testcase entity.
type TestcaseCreate struct {
	config
	mutation *TestcaseMutation
	hooks    []Hook
	conflict []sql.ConflictOption
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

// SetTaskID sets the "task" edge to the Task entity by ID.
func (tc *TestcaseCreate) SetTaskID(id int) *TestcaseCreate {
	tc.mutation.SetTaskID(id)
	return tc
}

// SetTask sets the "task" edge to the Task entity.
func (tc *TestcaseCreate) SetTask(t *Task) *TestcaseCreate {
	return tc.SetTaskID(t.ID)
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
	if _, ok := tc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task", err: errors.New(`ent: missing required edge "Testcase.task"`)}
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
	_spec.OnConflict = tc.conflict
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
	if nodes := tc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   testcase.TaskTable,
			Columns: []string{testcase.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.task_testcases = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Testcase.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TestcaseUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (tc *TestcaseCreate) OnConflict(opts ...sql.ConflictOption) *TestcaseUpsertOne {
	tc.conflict = opts
	return &TestcaseUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Testcase.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TestcaseCreate) OnConflictColumns(columns ...string) *TestcaseUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TestcaseUpsertOne{
		create: tc,
	}
}

type (
	// TestcaseUpsertOne is the builder for "upsert"-ing
	//  one Testcase node.
	TestcaseUpsertOne struct {
		create *TestcaseCreate
	}

	// TestcaseUpsert is the "OnConflict" setter.
	TestcaseUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *TestcaseUpsert) SetName(v string) *TestcaseUpsert {
	u.Set(testcase.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TestcaseUpsert) UpdateName() *TestcaseUpsert {
	u.SetExcluded(testcase.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *TestcaseUpsert) SetDescription(v string) *TestcaseUpsert {
	u.Set(testcase.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TestcaseUpsert) UpdateDescription() *TestcaseUpsert {
	u.SetExcluded(testcase.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *TestcaseUpsert) ClearDescription() *TestcaseUpsert {
	u.SetNull(testcase.FieldDescription)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TestcaseUpsert) SetCreatedAt(v time.Time) *TestcaseUpsert {
	u.Set(testcase.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TestcaseUpsert) UpdateCreatedAt() *TestcaseUpsert {
	u.SetExcluded(testcase.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TestcaseUpsert) SetUpdatedAt(v time.Time) *TestcaseUpsert {
	u.Set(testcase.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TestcaseUpsert) UpdateUpdatedAt() *TestcaseUpsert {
	u.SetExcluded(testcase.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *TestcaseUpsert) ClearUpdatedAt() *TestcaseUpsert {
	u.SetNull(testcase.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Testcase.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(testcase.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TestcaseUpsertOne) UpdateNewValues() *TestcaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(testcase.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Testcase.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TestcaseUpsertOne) Ignore() *TestcaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TestcaseUpsertOne) DoNothing() *TestcaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TestcaseCreate.OnConflict
// documentation for more info.
func (u *TestcaseUpsertOne) Update(set func(*TestcaseUpsert)) *TestcaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TestcaseUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *TestcaseUpsertOne) SetName(v string) *TestcaseUpsertOne {
	return u.Update(func(s *TestcaseUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TestcaseUpsertOne) UpdateName() *TestcaseUpsertOne {
	return u.Update(func(s *TestcaseUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *TestcaseUpsertOne) SetDescription(v string) *TestcaseUpsertOne {
	return u.Update(func(s *TestcaseUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TestcaseUpsertOne) UpdateDescription() *TestcaseUpsertOne {
	return u.Update(func(s *TestcaseUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *TestcaseUpsertOne) ClearDescription() *TestcaseUpsertOne {
	return u.Update(func(s *TestcaseUpsert) {
		s.ClearDescription()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TestcaseUpsertOne) SetCreatedAt(v time.Time) *TestcaseUpsertOne {
	return u.Update(func(s *TestcaseUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TestcaseUpsertOne) UpdateCreatedAt() *TestcaseUpsertOne {
	return u.Update(func(s *TestcaseUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TestcaseUpsertOne) SetUpdatedAt(v time.Time) *TestcaseUpsertOne {
	return u.Update(func(s *TestcaseUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TestcaseUpsertOne) UpdateUpdatedAt() *TestcaseUpsertOne {
	return u.Update(func(s *TestcaseUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *TestcaseUpsertOne) ClearUpdatedAt() *TestcaseUpsertOne {
	return u.Update(func(s *TestcaseUpsert) {
		s.ClearUpdatedAt()
	})
}

// Exec executes the query.
func (u *TestcaseUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TestcaseCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TestcaseUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TestcaseUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TestcaseUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TestcaseCreateBulk is the builder for creating many Testcase entities in bulk.
type TestcaseCreateBulk struct {
	config
	builders []*TestcaseCreate
	conflict []sql.ConflictOption
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
					spec.OnConflict = tcb.conflict
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Testcase.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TestcaseUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (tcb *TestcaseCreateBulk) OnConflict(opts ...sql.ConflictOption) *TestcaseUpsertBulk {
	tcb.conflict = opts
	return &TestcaseUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Testcase.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TestcaseCreateBulk) OnConflictColumns(columns ...string) *TestcaseUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TestcaseUpsertBulk{
		create: tcb,
	}
}

// TestcaseUpsertBulk is the builder for "upsert"-ing
// a bulk of Testcase nodes.
type TestcaseUpsertBulk struct {
	create *TestcaseCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Testcase.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(testcase.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TestcaseUpsertBulk) UpdateNewValues() *TestcaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(testcase.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Testcase.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TestcaseUpsertBulk) Ignore() *TestcaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TestcaseUpsertBulk) DoNothing() *TestcaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TestcaseCreateBulk.OnConflict
// documentation for more info.
func (u *TestcaseUpsertBulk) Update(set func(*TestcaseUpsert)) *TestcaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TestcaseUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *TestcaseUpsertBulk) SetName(v string) *TestcaseUpsertBulk {
	return u.Update(func(s *TestcaseUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TestcaseUpsertBulk) UpdateName() *TestcaseUpsertBulk {
	return u.Update(func(s *TestcaseUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *TestcaseUpsertBulk) SetDescription(v string) *TestcaseUpsertBulk {
	return u.Update(func(s *TestcaseUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TestcaseUpsertBulk) UpdateDescription() *TestcaseUpsertBulk {
	return u.Update(func(s *TestcaseUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *TestcaseUpsertBulk) ClearDescription() *TestcaseUpsertBulk {
	return u.Update(func(s *TestcaseUpsert) {
		s.ClearDescription()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TestcaseUpsertBulk) SetCreatedAt(v time.Time) *TestcaseUpsertBulk {
	return u.Update(func(s *TestcaseUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TestcaseUpsertBulk) UpdateCreatedAt() *TestcaseUpsertBulk {
	return u.Update(func(s *TestcaseUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TestcaseUpsertBulk) SetUpdatedAt(v time.Time) *TestcaseUpsertBulk {
	return u.Update(func(s *TestcaseUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TestcaseUpsertBulk) UpdateUpdatedAt() *TestcaseUpsertBulk {
	return u.Update(func(s *TestcaseUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *TestcaseUpsertBulk) ClearUpdatedAt() *TestcaseUpsertBulk {
	return u.Update(func(s *TestcaseUpsert) {
		s.ClearUpdatedAt()
	})
}

// Exec executes the query.
func (u *TestcaseUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TestcaseCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TestcaseCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TestcaseUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
