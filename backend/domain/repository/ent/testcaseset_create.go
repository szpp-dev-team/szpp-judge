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

// TestcaseSetCreate is the builder for creating a TestcaseSet entity.
type TestcaseSetCreate struct {
	config
	mutation *TestcaseSetMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (tsc *TestcaseSetCreate) SetName(s string) *TestcaseSetCreate {
	tsc.mutation.SetName(s)
	return tsc
}

// SetScoreRatio sets the "score_ratio" field.
func (tsc *TestcaseSetCreate) SetScoreRatio(i int) *TestcaseSetCreate {
	tsc.mutation.SetScoreRatio(i)
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
	if _, ok := tsc.mutation.ScoreRatio(); !ok {
		return &ValidationError{Name: "score_ratio", err: errors.New(`ent: missing required field "TestcaseSet.score_ratio"`)}
	}
	if _, ok := tsc.mutation.IsSample(); !ok {
		return &ValidationError{Name: "is_sample", err: errors.New(`ent: missing required field "TestcaseSet.is_sample"`)}
	}
	if _, ok := tsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TestcaseSet.created_at"`)}
	}
	if _, ok := tsc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task", err: errors.New(`ent: missing required edge "TestcaseSet.task"`)}
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
	_spec.OnConflict = tsc.conflict
	if id, ok := tsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tsc.mutation.Name(); ok {
		_spec.SetField(testcaseset.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tsc.mutation.ScoreRatio(); ok {
		_spec.SetField(testcaseset.FieldScoreRatio, field.TypeInt, value)
		_node.ScoreRatio = value
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TestcaseSet.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TestcaseSetUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (tsc *TestcaseSetCreate) OnConflict(opts ...sql.ConflictOption) *TestcaseSetUpsertOne {
	tsc.conflict = opts
	return &TestcaseSetUpsertOne{
		create: tsc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TestcaseSet.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tsc *TestcaseSetCreate) OnConflictColumns(columns ...string) *TestcaseSetUpsertOne {
	tsc.conflict = append(tsc.conflict, sql.ConflictColumns(columns...))
	return &TestcaseSetUpsertOne{
		create: tsc,
	}
}

type (
	// TestcaseSetUpsertOne is the builder for "upsert"-ing
	//  one TestcaseSet node.
	TestcaseSetUpsertOne struct {
		create *TestcaseSetCreate
	}

	// TestcaseSetUpsert is the "OnConflict" setter.
	TestcaseSetUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *TestcaseSetUpsert) SetName(v string) *TestcaseSetUpsert {
	u.Set(testcaseset.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TestcaseSetUpsert) UpdateName() *TestcaseSetUpsert {
	u.SetExcluded(testcaseset.FieldName)
	return u
}

// SetScoreRatio sets the "score_ratio" field.
func (u *TestcaseSetUpsert) SetScoreRatio(v int) *TestcaseSetUpsert {
	u.Set(testcaseset.FieldScoreRatio, v)
	return u
}

// UpdateScoreRatio sets the "score_ratio" field to the value that was provided on create.
func (u *TestcaseSetUpsert) UpdateScoreRatio() *TestcaseSetUpsert {
	u.SetExcluded(testcaseset.FieldScoreRatio)
	return u
}

// AddScoreRatio adds v to the "score_ratio" field.
func (u *TestcaseSetUpsert) AddScoreRatio(v int) *TestcaseSetUpsert {
	u.Add(testcaseset.FieldScoreRatio, v)
	return u
}

// SetIsSample sets the "is_sample" field.
func (u *TestcaseSetUpsert) SetIsSample(v bool) *TestcaseSetUpsert {
	u.Set(testcaseset.FieldIsSample, v)
	return u
}

// UpdateIsSample sets the "is_sample" field to the value that was provided on create.
func (u *TestcaseSetUpsert) UpdateIsSample() *TestcaseSetUpsert {
	u.SetExcluded(testcaseset.FieldIsSample)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TestcaseSetUpsert) SetCreatedAt(v time.Time) *TestcaseSetUpsert {
	u.Set(testcaseset.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TestcaseSetUpsert) UpdateCreatedAt() *TestcaseSetUpsert {
	u.SetExcluded(testcaseset.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TestcaseSetUpsert) SetUpdatedAt(v time.Time) *TestcaseSetUpsert {
	u.Set(testcaseset.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TestcaseSetUpsert) UpdateUpdatedAt() *TestcaseSetUpsert {
	u.SetExcluded(testcaseset.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *TestcaseSetUpsert) ClearUpdatedAt() *TestcaseSetUpsert {
	u.SetNull(testcaseset.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TestcaseSet.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(testcaseset.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TestcaseSetUpsertOne) UpdateNewValues() *TestcaseSetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(testcaseset.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TestcaseSet.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TestcaseSetUpsertOne) Ignore() *TestcaseSetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TestcaseSetUpsertOne) DoNothing() *TestcaseSetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TestcaseSetCreate.OnConflict
// documentation for more info.
func (u *TestcaseSetUpsertOne) Update(set func(*TestcaseSetUpsert)) *TestcaseSetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TestcaseSetUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *TestcaseSetUpsertOne) SetName(v string) *TestcaseSetUpsertOne {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TestcaseSetUpsertOne) UpdateName() *TestcaseSetUpsertOne {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.UpdateName()
	})
}

// SetScoreRatio sets the "score_ratio" field.
func (u *TestcaseSetUpsertOne) SetScoreRatio(v int) *TestcaseSetUpsertOne {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.SetScoreRatio(v)
	})
}

// AddScoreRatio adds v to the "score_ratio" field.
func (u *TestcaseSetUpsertOne) AddScoreRatio(v int) *TestcaseSetUpsertOne {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.AddScoreRatio(v)
	})
}

// UpdateScoreRatio sets the "score_ratio" field to the value that was provided on create.
func (u *TestcaseSetUpsertOne) UpdateScoreRatio() *TestcaseSetUpsertOne {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.UpdateScoreRatio()
	})
}

// SetIsSample sets the "is_sample" field.
func (u *TestcaseSetUpsertOne) SetIsSample(v bool) *TestcaseSetUpsertOne {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.SetIsSample(v)
	})
}

// UpdateIsSample sets the "is_sample" field to the value that was provided on create.
func (u *TestcaseSetUpsertOne) UpdateIsSample() *TestcaseSetUpsertOne {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.UpdateIsSample()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TestcaseSetUpsertOne) SetCreatedAt(v time.Time) *TestcaseSetUpsertOne {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TestcaseSetUpsertOne) UpdateCreatedAt() *TestcaseSetUpsertOne {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TestcaseSetUpsertOne) SetUpdatedAt(v time.Time) *TestcaseSetUpsertOne {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TestcaseSetUpsertOne) UpdateUpdatedAt() *TestcaseSetUpsertOne {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *TestcaseSetUpsertOne) ClearUpdatedAt() *TestcaseSetUpsertOne {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.ClearUpdatedAt()
	})
}

// Exec executes the query.
func (u *TestcaseSetUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TestcaseSetCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TestcaseSetUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TestcaseSetUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TestcaseSetUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TestcaseSetCreateBulk is the builder for creating many TestcaseSet entities in bulk.
type TestcaseSetCreateBulk struct {
	config
	builders []*TestcaseSetCreate
	conflict []sql.ConflictOption
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
					spec.OnConflict = tscb.conflict
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TestcaseSet.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TestcaseSetUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (tscb *TestcaseSetCreateBulk) OnConflict(opts ...sql.ConflictOption) *TestcaseSetUpsertBulk {
	tscb.conflict = opts
	return &TestcaseSetUpsertBulk{
		create: tscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TestcaseSet.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tscb *TestcaseSetCreateBulk) OnConflictColumns(columns ...string) *TestcaseSetUpsertBulk {
	tscb.conflict = append(tscb.conflict, sql.ConflictColumns(columns...))
	return &TestcaseSetUpsertBulk{
		create: tscb,
	}
}

// TestcaseSetUpsertBulk is the builder for "upsert"-ing
// a bulk of TestcaseSet nodes.
type TestcaseSetUpsertBulk struct {
	create *TestcaseSetCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TestcaseSet.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(testcaseset.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TestcaseSetUpsertBulk) UpdateNewValues() *TestcaseSetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(testcaseset.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TestcaseSet.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TestcaseSetUpsertBulk) Ignore() *TestcaseSetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TestcaseSetUpsertBulk) DoNothing() *TestcaseSetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TestcaseSetCreateBulk.OnConflict
// documentation for more info.
func (u *TestcaseSetUpsertBulk) Update(set func(*TestcaseSetUpsert)) *TestcaseSetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TestcaseSetUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *TestcaseSetUpsertBulk) SetName(v string) *TestcaseSetUpsertBulk {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TestcaseSetUpsertBulk) UpdateName() *TestcaseSetUpsertBulk {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.UpdateName()
	})
}

// SetScoreRatio sets the "score_ratio" field.
func (u *TestcaseSetUpsertBulk) SetScoreRatio(v int) *TestcaseSetUpsertBulk {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.SetScoreRatio(v)
	})
}

// AddScoreRatio adds v to the "score_ratio" field.
func (u *TestcaseSetUpsertBulk) AddScoreRatio(v int) *TestcaseSetUpsertBulk {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.AddScoreRatio(v)
	})
}

// UpdateScoreRatio sets the "score_ratio" field to the value that was provided on create.
func (u *TestcaseSetUpsertBulk) UpdateScoreRatio() *TestcaseSetUpsertBulk {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.UpdateScoreRatio()
	})
}

// SetIsSample sets the "is_sample" field.
func (u *TestcaseSetUpsertBulk) SetIsSample(v bool) *TestcaseSetUpsertBulk {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.SetIsSample(v)
	})
}

// UpdateIsSample sets the "is_sample" field to the value that was provided on create.
func (u *TestcaseSetUpsertBulk) UpdateIsSample() *TestcaseSetUpsertBulk {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.UpdateIsSample()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TestcaseSetUpsertBulk) SetCreatedAt(v time.Time) *TestcaseSetUpsertBulk {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TestcaseSetUpsertBulk) UpdateCreatedAt() *TestcaseSetUpsertBulk {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TestcaseSetUpsertBulk) SetUpdatedAt(v time.Time) *TestcaseSetUpsertBulk {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TestcaseSetUpsertBulk) UpdateUpdatedAt() *TestcaseSetUpsertBulk {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *TestcaseSetUpsertBulk) ClearUpdatedAt() *TestcaseSetUpsertBulk {
	return u.Update(func(s *TestcaseSetUpsert) {
		s.ClearUpdatedAt()
	})
}

// Exec executes the query.
func (u *TestcaseSetUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TestcaseSetCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TestcaseSetCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TestcaseSetUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
