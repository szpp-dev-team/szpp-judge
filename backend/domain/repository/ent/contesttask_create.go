// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contest"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contesttask"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
)

// ContestTaskCreate is the builder for creating a ContestTask entity.
type ContestTaskCreate struct {
	config
	mutation *ContestTaskMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetScore sets the "score" field.
func (ctc *ContestTaskCreate) SetScore(i int) *ContestTaskCreate {
	ctc.mutation.SetScore(i)
	return ctc
}

// SetOrder sets the "order" field.
func (ctc *ContestTaskCreate) SetOrder(i int) *ContestTaskCreate {
	ctc.mutation.SetOrder(i)
	return ctc
}

// SetContestID sets the "contest_id" field.
func (ctc *ContestTaskCreate) SetContestID(i int) *ContestTaskCreate {
	ctc.mutation.SetContestID(i)
	return ctc
}

// SetTaskID sets the "task_id" field.
func (ctc *ContestTaskCreate) SetTaskID(i int) *ContestTaskCreate {
	ctc.mutation.SetTaskID(i)
	return ctc
}

// SetID sets the "id" field.
func (ctc *ContestTaskCreate) SetID(i int) *ContestTaskCreate {
	ctc.mutation.SetID(i)
	return ctc
}

// SetContest sets the "contest" edge to the Contest entity.
func (ctc *ContestTaskCreate) SetContest(c *Contest) *ContestTaskCreate {
	return ctc.SetContestID(c.ID)
}

// SetTask sets the "task" edge to the Task entity.
func (ctc *ContestTaskCreate) SetTask(t *Task) *ContestTaskCreate {
	return ctc.SetTaskID(t.ID)
}

// Mutation returns the ContestTaskMutation object of the builder.
func (ctc *ContestTaskCreate) Mutation() *ContestTaskMutation {
	return ctc.mutation
}

// Save creates the ContestTask in the database.
func (ctc *ContestTaskCreate) Save(ctx context.Context) (*ContestTask, error) {
	return withHooks(ctx, ctc.sqlSave, ctc.mutation, ctc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ctc *ContestTaskCreate) SaveX(ctx context.Context) *ContestTask {
	v, err := ctc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctc *ContestTaskCreate) Exec(ctx context.Context) error {
	_, err := ctc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctc *ContestTaskCreate) ExecX(ctx context.Context) {
	if err := ctc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctc *ContestTaskCreate) check() error {
	if _, ok := ctc.mutation.Score(); !ok {
		return &ValidationError{Name: "score", err: errors.New(`ent: missing required field "ContestTask.score"`)}
	}
	if _, ok := ctc.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "ContestTask.order"`)}
	}
	if _, ok := ctc.mutation.ContestID(); !ok {
		return &ValidationError{Name: "contest_id", err: errors.New(`ent: missing required field "ContestTask.contest_id"`)}
	}
	if _, ok := ctc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task_id", err: errors.New(`ent: missing required field "ContestTask.task_id"`)}
	}
	if _, ok := ctc.mutation.ContestID(); !ok {
		return &ValidationError{Name: "contest", err: errors.New(`ent: missing required edge "ContestTask.contest"`)}
	}
	if _, ok := ctc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task", err: errors.New(`ent: missing required edge "ContestTask.task"`)}
	}
	return nil
}

func (ctc *ContestTaskCreate) sqlSave(ctx context.Context) (*ContestTask, error) {
	if err := ctc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ctc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ctc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	ctc.mutation.id = &_node.ID
	ctc.mutation.done = true
	return _node, nil
}

func (ctc *ContestTaskCreate) createSpec() (*ContestTask, *sqlgraph.CreateSpec) {
	var (
		_node = &ContestTask{config: ctc.config}
		_spec = sqlgraph.NewCreateSpec(contesttask.Table, sqlgraph.NewFieldSpec(contesttask.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ctc.conflict
	if id, ok := ctc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ctc.mutation.Score(); ok {
		_spec.SetField(contesttask.FieldScore, field.TypeInt, value)
		_node.Score = value
	}
	if value, ok := ctc.mutation.Order(); ok {
		_spec.SetField(contesttask.FieldOrder, field.TypeInt, value)
		_node.Order = value
	}
	if nodes := ctc.mutation.ContestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   contesttask.ContestTable,
			Columns: []string{contesttask.ContestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ContestID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ctc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   contesttask.TaskTable,
			Columns: []string{contesttask.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TaskID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ContestTask.Create().
//		SetScore(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ContestTaskUpsert) {
//			SetScore(v+v).
//		}).
//		Exec(ctx)
func (ctc *ContestTaskCreate) OnConflict(opts ...sql.ConflictOption) *ContestTaskUpsertOne {
	ctc.conflict = opts
	return &ContestTaskUpsertOne{
		create: ctc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ContestTask.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ctc *ContestTaskCreate) OnConflictColumns(columns ...string) *ContestTaskUpsertOne {
	ctc.conflict = append(ctc.conflict, sql.ConflictColumns(columns...))
	return &ContestTaskUpsertOne{
		create: ctc,
	}
}

type (
	// ContestTaskUpsertOne is the builder for "upsert"-ing
	//  one ContestTask node.
	ContestTaskUpsertOne struct {
		create *ContestTaskCreate
	}

	// ContestTaskUpsert is the "OnConflict" setter.
	ContestTaskUpsert struct {
		*sql.UpdateSet
	}
)

// SetScore sets the "score" field.
func (u *ContestTaskUpsert) SetScore(v int) *ContestTaskUpsert {
	u.Set(contesttask.FieldScore, v)
	return u
}

// UpdateScore sets the "score" field to the value that was provided on create.
func (u *ContestTaskUpsert) UpdateScore() *ContestTaskUpsert {
	u.SetExcluded(contesttask.FieldScore)
	return u
}

// AddScore adds v to the "score" field.
func (u *ContestTaskUpsert) AddScore(v int) *ContestTaskUpsert {
	u.Add(contesttask.FieldScore, v)
	return u
}

// SetOrder sets the "order" field.
func (u *ContestTaskUpsert) SetOrder(v int) *ContestTaskUpsert {
	u.Set(contesttask.FieldOrder, v)
	return u
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *ContestTaskUpsert) UpdateOrder() *ContestTaskUpsert {
	u.SetExcluded(contesttask.FieldOrder)
	return u
}

// AddOrder adds v to the "order" field.
func (u *ContestTaskUpsert) AddOrder(v int) *ContestTaskUpsert {
	u.Add(contesttask.FieldOrder, v)
	return u
}

// SetContestID sets the "contest_id" field.
func (u *ContestTaskUpsert) SetContestID(v int) *ContestTaskUpsert {
	u.Set(contesttask.FieldContestID, v)
	return u
}

// UpdateContestID sets the "contest_id" field to the value that was provided on create.
func (u *ContestTaskUpsert) UpdateContestID() *ContestTaskUpsert {
	u.SetExcluded(contesttask.FieldContestID)
	return u
}

// SetTaskID sets the "task_id" field.
func (u *ContestTaskUpsert) SetTaskID(v int) *ContestTaskUpsert {
	u.Set(contesttask.FieldTaskID, v)
	return u
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *ContestTaskUpsert) UpdateTaskID() *ContestTaskUpsert {
	u.SetExcluded(contesttask.FieldTaskID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ContestTask.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(contesttask.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ContestTaskUpsertOne) UpdateNewValues() *ContestTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(contesttask.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ContestTask.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ContestTaskUpsertOne) Ignore() *ContestTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ContestTaskUpsertOne) DoNothing() *ContestTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ContestTaskCreate.OnConflict
// documentation for more info.
func (u *ContestTaskUpsertOne) Update(set func(*ContestTaskUpsert)) *ContestTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ContestTaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetScore sets the "score" field.
func (u *ContestTaskUpsertOne) SetScore(v int) *ContestTaskUpsertOne {
	return u.Update(func(s *ContestTaskUpsert) {
		s.SetScore(v)
	})
}

// AddScore adds v to the "score" field.
func (u *ContestTaskUpsertOne) AddScore(v int) *ContestTaskUpsertOne {
	return u.Update(func(s *ContestTaskUpsert) {
		s.AddScore(v)
	})
}

// UpdateScore sets the "score" field to the value that was provided on create.
func (u *ContestTaskUpsertOne) UpdateScore() *ContestTaskUpsertOne {
	return u.Update(func(s *ContestTaskUpsert) {
		s.UpdateScore()
	})
}

// SetOrder sets the "order" field.
func (u *ContestTaskUpsertOne) SetOrder(v int) *ContestTaskUpsertOne {
	return u.Update(func(s *ContestTaskUpsert) {
		s.SetOrder(v)
	})
}

// AddOrder adds v to the "order" field.
func (u *ContestTaskUpsertOne) AddOrder(v int) *ContestTaskUpsertOne {
	return u.Update(func(s *ContestTaskUpsert) {
		s.AddOrder(v)
	})
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *ContestTaskUpsertOne) UpdateOrder() *ContestTaskUpsertOne {
	return u.Update(func(s *ContestTaskUpsert) {
		s.UpdateOrder()
	})
}

// SetContestID sets the "contest_id" field.
func (u *ContestTaskUpsertOne) SetContestID(v int) *ContestTaskUpsertOne {
	return u.Update(func(s *ContestTaskUpsert) {
		s.SetContestID(v)
	})
}

// UpdateContestID sets the "contest_id" field to the value that was provided on create.
func (u *ContestTaskUpsertOne) UpdateContestID() *ContestTaskUpsertOne {
	return u.Update(func(s *ContestTaskUpsert) {
		s.UpdateContestID()
	})
}

// SetTaskID sets the "task_id" field.
func (u *ContestTaskUpsertOne) SetTaskID(v int) *ContestTaskUpsertOne {
	return u.Update(func(s *ContestTaskUpsert) {
		s.SetTaskID(v)
	})
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *ContestTaskUpsertOne) UpdateTaskID() *ContestTaskUpsertOne {
	return u.Update(func(s *ContestTaskUpsert) {
		s.UpdateTaskID()
	})
}

// Exec executes the query.
func (u *ContestTaskUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ContestTaskCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ContestTaskUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ContestTaskUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ContestTaskUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ContestTaskCreateBulk is the builder for creating many ContestTask entities in bulk.
type ContestTaskCreateBulk struct {
	config
	builders []*ContestTaskCreate
	conflict []sql.ConflictOption
}

// Save creates the ContestTask entities in the database.
func (ctcb *ContestTaskCreateBulk) Save(ctx context.Context) ([]*ContestTask, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ctcb.builders))
	nodes := make([]*ContestTask, len(ctcb.builders))
	mutators := make([]Mutator, len(ctcb.builders))
	for i := range ctcb.builders {
		func(i int, root context.Context) {
			builder := ctcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContestTaskMutation)
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
					_, err = mutators[i+1].Mutate(root, ctcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ctcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ctcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ctcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ctcb *ContestTaskCreateBulk) SaveX(ctx context.Context) []*ContestTask {
	v, err := ctcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctcb *ContestTaskCreateBulk) Exec(ctx context.Context) error {
	_, err := ctcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctcb *ContestTaskCreateBulk) ExecX(ctx context.Context) {
	if err := ctcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ContestTask.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ContestTaskUpsert) {
//			SetScore(v+v).
//		}).
//		Exec(ctx)
func (ctcb *ContestTaskCreateBulk) OnConflict(opts ...sql.ConflictOption) *ContestTaskUpsertBulk {
	ctcb.conflict = opts
	return &ContestTaskUpsertBulk{
		create: ctcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ContestTask.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ctcb *ContestTaskCreateBulk) OnConflictColumns(columns ...string) *ContestTaskUpsertBulk {
	ctcb.conflict = append(ctcb.conflict, sql.ConflictColumns(columns...))
	return &ContestTaskUpsertBulk{
		create: ctcb,
	}
}

// ContestTaskUpsertBulk is the builder for "upsert"-ing
// a bulk of ContestTask nodes.
type ContestTaskUpsertBulk struct {
	create *ContestTaskCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ContestTask.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(contesttask.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ContestTaskUpsertBulk) UpdateNewValues() *ContestTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(contesttask.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ContestTask.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ContestTaskUpsertBulk) Ignore() *ContestTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ContestTaskUpsertBulk) DoNothing() *ContestTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ContestTaskCreateBulk.OnConflict
// documentation for more info.
func (u *ContestTaskUpsertBulk) Update(set func(*ContestTaskUpsert)) *ContestTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ContestTaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetScore sets the "score" field.
func (u *ContestTaskUpsertBulk) SetScore(v int) *ContestTaskUpsertBulk {
	return u.Update(func(s *ContestTaskUpsert) {
		s.SetScore(v)
	})
}

// AddScore adds v to the "score" field.
func (u *ContestTaskUpsertBulk) AddScore(v int) *ContestTaskUpsertBulk {
	return u.Update(func(s *ContestTaskUpsert) {
		s.AddScore(v)
	})
}

// UpdateScore sets the "score" field to the value that was provided on create.
func (u *ContestTaskUpsertBulk) UpdateScore() *ContestTaskUpsertBulk {
	return u.Update(func(s *ContestTaskUpsert) {
		s.UpdateScore()
	})
}

// SetOrder sets the "order" field.
func (u *ContestTaskUpsertBulk) SetOrder(v int) *ContestTaskUpsertBulk {
	return u.Update(func(s *ContestTaskUpsert) {
		s.SetOrder(v)
	})
}

// AddOrder adds v to the "order" field.
func (u *ContestTaskUpsertBulk) AddOrder(v int) *ContestTaskUpsertBulk {
	return u.Update(func(s *ContestTaskUpsert) {
		s.AddOrder(v)
	})
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *ContestTaskUpsertBulk) UpdateOrder() *ContestTaskUpsertBulk {
	return u.Update(func(s *ContestTaskUpsert) {
		s.UpdateOrder()
	})
}

// SetContestID sets the "contest_id" field.
func (u *ContestTaskUpsertBulk) SetContestID(v int) *ContestTaskUpsertBulk {
	return u.Update(func(s *ContestTaskUpsert) {
		s.SetContestID(v)
	})
}

// UpdateContestID sets the "contest_id" field to the value that was provided on create.
func (u *ContestTaskUpsertBulk) UpdateContestID() *ContestTaskUpsertBulk {
	return u.Update(func(s *ContestTaskUpsert) {
		s.UpdateContestID()
	})
}

// SetTaskID sets the "task_id" field.
func (u *ContestTaskUpsertBulk) SetTaskID(v int) *ContestTaskUpsertBulk {
	return u.Update(func(s *ContestTaskUpsert) {
		s.SetTaskID(v)
	})
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *ContestTaskUpsertBulk) UpdateTaskID() *ContestTaskUpsertBulk {
	return u.Update(func(s *ContestTaskUpsert) {
		s.UpdateTaskID()
	})
}

// Exec executes the query.
func (u *ContestTaskUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ContestTaskCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ContestTaskCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ContestTaskUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
