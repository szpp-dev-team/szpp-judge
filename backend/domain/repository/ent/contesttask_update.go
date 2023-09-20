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
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/predicate"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
)

// ContestTaskUpdate is the builder for updating ContestTask entities.
type ContestTaskUpdate struct {
	config
	hooks    []Hook
	mutation *ContestTaskMutation
}

// Where appends a list predicates to the ContestTaskUpdate builder.
func (ctu *ContestTaskUpdate) Where(ps ...predicate.ContestTask) *ContestTaskUpdate {
	ctu.mutation.Where(ps...)
	return ctu
}

// SetScore sets the "score" field.
func (ctu *ContestTaskUpdate) SetScore(i int) *ContestTaskUpdate {
	ctu.mutation.ResetScore()
	ctu.mutation.SetScore(i)
	return ctu
}

// AddScore adds i to the "score" field.
func (ctu *ContestTaskUpdate) AddScore(i int) *ContestTaskUpdate {
	ctu.mutation.AddScore(i)
	return ctu
}

// SetContestID sets the "contest_id" field.
func (ctu *ContestTaskUpdate) SetContestID(i int) *ContestTaskUpdate {
	ctu.mutation.SetContestID(i)
	return ctu
}

// SetTaskID sets the "task_id" field.
func (ctu *ContestTaskUpdate) SetTaskID(i int) *ContestTaskUpdate {
	ctu.mutation.SetTaskID(i)
	return ctu
}

// SetContest sets the "contest" edge to the Contest entity.
func (ctu *ContestTaskUpdate) SetContest(c *Contest) *ContestTaskUpdate {
	return ctu.SetContestID(c.ID)
}

// SetTask sets the "task" edge to the Task entity.
func (ctu *ContestTaskUpdate) SetTask(t *Task) *ContestTaskUpdate {
	return ctu.SetTaskID(t.ID)
}

// Mutation returns the ContestTaskMutation object of the builder.
func (ctu *ContestTaskUpdate) Mutation() *ContestTaskMutation {
	return ctu.mutation
}

// ClearContest clears the "contest" edge to the Contest entity.
func (ctu *ContestTaskUpdate) ClearContest() *ContestTaskUpdate {
	ctu.mutation.ClearContest()
	return ctu
}

// ClearTask clears the "task" edge to the Task entity.
func (ctu *ContestTaskUpdate) ClearTask() *ContestTaskUpdate {
	ctu.mutation.ClearTask()
	return ctu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctu *ContestTaskUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ctu.sqlSave, ctu.mutation, ctu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ctu *ContestTaskUpdate) SaveX(ctx context.Context) int {
	affected, err := ctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctu *ContestTaskUpdate) Exec(ctx context.Context) error {
	_, err := ctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctu *ContestTaskUpdate) ExecX(ctx context.Context) {
	if err := ctu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctu *ContestTaskUpdate) check() error {
	if _, ok := ctu.mutation.ContestID(); ctu.mutation.ContestCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ContestTask.contest"`)
	}
	if _, ok := ctu.mutation.TaskID(); ctu.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ContestTask.task"`)
	}
	return nil
}

func (ctu *ContestTaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ctu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(contesttask.Table, contesttask.Columns, sqlgraph.NewFieldSpec(contesttask.FieldID, field.TypeInt))
	if ps := ctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctu.mutation.Score(); ok {
		_spec.SetField(contesttask.FieldScore, field.TypeInt, value)
	}
	if value, ok := ctu.mutation.AddedScore(); ok {
		_spec.AddField(contesttask.FieldScore, field.TypeInt, value)
	}
	if ctu.mutation.ContestCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.ContestIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ctu.mutation.TaskCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.TaskIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contesttask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ctu.mutation.done = true
	return n, nil
}

// ContestTaskUpdateOne is the builder for updating a single ContestTask entity.
type ContestTaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContestTaskMutation
}

// SetScore sets the "score" field.
func (ctuo *ContestTaskUpdateOne) SetScore(i int) *ContestTaskUpdateOne {
	ctuo.mutation.ResetScore()
	ctuo.mutation.SetScore(i)
	return ctuo
}

// AddScore adds i to the "score" field.
func (ctuo *ContestTaskUpdateOne) AddScore(i int) *ContestTaskUpdateOne {
	ctuo.mutation.AddScore(i)
	return ctuo
}

// SetContestID sets the "contest_id" field.
func (ctuo *ContestTaskUpdateOne) SetContestID(i int) *ContestTaskUpdateOne {
	ctuo.mutation.SetContestID(i)
	return ctuo
}

// SetTaskID sets the "task_id" field.
func (ctuo *ContestTaskUpdateOne) SetTaskID(i int) *ContestTaskUpdateOne {
	ctuo.mutation.SetTaskID(i)
	return ctuo
}

// SetContest sets the "contest" edge to the Contest entity.
func (ctuo *ContestTaskUpdateOne) SetContest(c *Contest) *ContestTaskUpdateOne {
	return ctuo.SetContestID(c.ID)
}

// SetTask sets the "task" edge to the Task entity.
func (ctuo *ContestTaskUpdateOne) SetTask(t *Task) *ContestTaskUpdateOne {
	return ctuo.SetTaskID(t.ID)
}

// Mutation returns the ContestTaskMutation object of the builder.
func (ctuo *ContestTaskUpdateOne) Mutation() *ContestTaskMutation {
	return ctuo.mutation
}

// ClearContest clears the "contest" edge to the Contest entity.
func (ctuo *ContestTaskUpdateOne) ClearContest() *ContestTaskUpdateOne {
	ctuo.mutation.ClearContest()
	return ctuo
}

// ClearTask clears the "task" edge to the Task entity.
func (ctuo *ContestTaskUpdateOne) ClearTask() *ContestTaskUpdateOne {
	ctuo.mutation.ClearTask()
	return ctuo
}

// Where appends a list predicates to the ContestTaskUpdate builder.
func (ctuo *ContestTaskUpdateOne) Where(ps ...predicate.ContestTask) *ContestTaskUpdateOne {
	ctuo.mutation.Where(ps...)
	return ctuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctuo *ContestTaskUpdateOne) Select(field string, fields ...string) *ContestTaskUpdateOne {
	ctuo.fields = append([]string{field}, fields...)
	return ctuo
}

// Save executes the query and returns the updated ContestTask entity.
func (ctuo *ContestTaskUpdateOne) Save(ctx context.Context) (*ContestTask, error) {
	return withHooks(ctx, ctuo.sqlSave, ctuo.mutation, ctuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ctuo *ContestTaskUpdateOne) SaveX(ctx context.Context) *ContestTask {
	node, err := ctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctuo *ContestTaskUpdateOne) Exec(ctx context.Context) error {
	_, err := ctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctuo *ContestTaskUpdateOne) ExecX(ctx context.Context) {
	if err := ctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctuo *ContestTaskUpdateOne) check() error {
	if _, ok := ctuo.mutation.ContestID(); ctuo.mutation.ContestCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ContestTask.contest"`)
	}
	if _, ok := ctuo.mutation.TaskID(); ctuo.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ContestTask.task"`)
	}
	return nil
}

func (ctuo *ContestTaskUpdateOne) sqlSave(ctx context.Context) (_node *ContestTask, err error) {
	if err := ctuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(contesttask.Table, contesttask.Columns, sqlgraph.NewFieldSpec(contesttask.FieldID, field.TypeInt))
	id, ok := ctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ContestTask.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contesttask.FieldID)
		for _, f := range fields {
			if !contesttask.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != contesttask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctuo.mutation.Score(); ok {
		_spec.SetField(contesttask.FieldScore, field.TypeInt, value)
	}
	if value, ok := ctuo.mutation.AddedScore(); ok {
		_spec.AddField(contesttask.FieldScore, field.TypeInt, value)
	}
	if ctuo.mutation.ContestCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.ContestIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ctuo.mutation.TaskCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.TaskIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ContestTask{config: ctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contesttask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ctuo.mutation.done = true
	return _node, nil
}
