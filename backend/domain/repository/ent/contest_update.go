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
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contest"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contesttask"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contestuser"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/predicate"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/submit"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
)

// ContestUpdate is the builder for updating Contest entities.
type ContestUpdate struct {
	config
	hooks    []Hook
	mutation *ContestMutation
}

// Where appends a list predicates to the ContestUpdate builder.
func (cu *ContestUpdate) Where(ps ...predicate.Contest) *ContestUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *ContestUpdate) SetName(s string) *ContestUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetSlug sets the "slug" field.
func (cu *ContestUpdate) SetSlug(s string) *ContestUpdate {
	cu.mutation.SetSlug(s)
	return cu
}

// SetDescription sets the "description" field.
func (cu *ContestUpdate) SetDescription(s string) *ContestUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetPenaltySeconds sets the "penalty_seconds" field.
func (cu *ContestUpdate) SetPenaltySeconds(i int) *ContestUpdate {
	cu.mutation.ResetPenaltySeconds()
	cu.mutation.SetPenaltySeconds(i)
	return cu
}

// AddPenaltySeconds adds i to the "penalty_seconds" field.
func (cu *ContestUpdate) AddPenaltySeconds(i int) *ContestUpdate {
	cu.mutation.AddPenaltySeconds(i)
	return cu
}

// SetContestType sets the "contest_type" field.
func (cu *ContestUpdate) SetContestType(s string) *ContestUpdate {
	cu.mutation.SetContestType(s)
	return cu
}

// SetIsPublic sets the "is_public" field.
func (cu *ContestUpdate) SetIsPublic(b bool) *ContestUpdate {
	cu.mutation.SetIsPublic(b)
	return cu
}

// SetStartAt sets the "start_at" field.
func (cu *ContestUpdate) SetStartAt(t time.Time) *ContestUpdate {
	cu.mutation.SetStartAt(t)
	return cu
}

// SetEndAt sets the "end_at" field.
func (cu *ContestUpdate) SetEndAt(t time.Time) *ContestUpdate {
	cu.mutation.SetEndAt(t)
	return cu
}

// AddSubmitIDs adds the "submits" edge to the Submit entity by IDs.
func (cu *ContestUpdate) AddSubmitIDs(ids ...int) *ContestUpdate {
	cu.mutation.AddSubmitIDs(ids...)
	return cu
}

// AddSubmits adds the "submits" edges to the Submit entity.
func (cu *ContestUpdate) AddSubmits(s ...*Submit) *ContestUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddSubmitIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (cu *ContestUpdate) AddUserIDs(ids ...int) *ContestUpdate {
	cu.mutation.AddUserIDs(ids...)
	return cu
}

// AddUsers adds the "users" edges to the User entity.
func (cu *ContestUpdate) AddUsers(u ...*User) *ContestUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.AddUserIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (cu *ContestUpdate) AddTaskIDs(ids ...int) *ContestUpdate {
	cu.mutation.AddTaskIDs(ids...)
	return cu
}

// AddTasks adds the "tasks" edges to the Task entity.
func (cu *ContestUpdate) AddTasks(t ...*Task) *ContestUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.AddTaskIDs(ids...)
}

// AddContestUserIDs adds the "contest_user" edge to the ContestUser entity by IDs.
func (cu *ContestUpdate) AddContestUserIDs(ids ...int) *ContestUpdate {
	cu.mutation.AddContestUserIDs(ids...)
	return cu
}

// AddContestUser adds the "contest_user" edges to the ContestUser entity.
func (cu *ContestUpdate) AddContestUser(c ...*ContestUser) *ContestUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddContestUserIDs(ids...)
}

// AddContestTaskIDs adds the "contest_task" edge to the ContestTask entity by IDs.
func (cu *ContestUpdate) AddContestTaskIDs(ids ...int) *ContestUpdate {
	cu.mutation.AddContestTaskIDs(ids...)
	return cu
}

// AddContestTask adds the "contest_task" edges to the ContestTask entity.
func (cu *ContestUpdate) AddContestTask(c ...*ContestTask) *ContestUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddContestTaskIDs(ids...)
}

// Mutation returns the ContestMutation object of the builder.
func (cu *ContestUpdate) Mutation() *ContestMutation {
	return cu.mutation
}

// ClearSubmits clears all "submits" edges to the Submit entity.
func (cu *ContestUpdate) ClearSubmits() *ContestUpdate {
	cu.mutation.ClearSubmits()
	return cu
}

// RemoveSubmitIDs removes the "submits" edge to Submit entities by IDs.
func (cu *ContestUpdate) RemoveSubmitIDs(ids ...int) *ContestUpdate {
	cu.mutation.RemoveSubmitIDs(ids...)
	return cu
}

// RemoveSubmits removes "submits" edges to Submit entities.
func (cu *ContestUpdate) RemoveSubmits(s ...*Submit) *ContestUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveSubmitIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (cu *ContestUpdate) ClearUsers() *ContestUpdate {
	cu.mutation.ClearUsers()
	return cu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (cu *ContestUpdate) RemoveUserIDs(ids ...int) *ContestUpdate {
	cu.mutation.RemoveUserIDs(ids...)
	return cu
}

// RemoveUsers removes "users" edges to User entities.
func (cu *ContestUpdate) RemoveUsers(u ...*User) *ContestUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.RemoveUserIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (cu *ContestUpdate) ClearTasks() *ContestUpdate {
	cu.mutation.ClearTasks()
	return cu
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (cu *ContestUpdate) RemoveTaskIDs(ids ...int) *ContestUpdate {
	cu.mutation.RemoveTaskIDs(ids...)
	return cu
}

// RemoveTasks removes "tasks" edges to Task entities.
func (cu *ContestUpdate) RemoveTasks(t ...*Task) *ContestUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.RemoveTaskIDs(ids...)
}

// ClearContestUser clears all "contest_user" edges to the ContestUser entity.
func (cu *ContestUpdate) ClearContestUser() *ContestUpdate {
	cu.mutation.ClearContestUser()
	return cu
}

// RemoveContestUserIDs removes the "contest_user" edge to ContestUser entities by IDs.
func (cu *ContestUpdate) RemoveContestUserIDs(ids ...int) *ContestUpdate {
	cu.mutation.RemoveContestUserIDs(ids...)
	return cu
}

// RemoveContestUser removes "contest_user" edges to ContestUser entities.
func (cu *ContestUpdate) RemoveContestUser(c ...*ContestUser) *ContestUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveContestUserIDs(ids...)
}

// ClearContestTask clears all "contest_task" edges to the ContestTask entity.
func (cu *ContestUpdate) ClearContestTask() *ContestUpdate {
	cu.mutation.ClearContestTask()
	return cu
}

// RemoveContestTaskIDs removes the "contest_task" edge to ContestTask entities by IDs.
func (cu *ContestUpdate) RemoveContestTaskIDs(ids ...int) *ContestUpdate {
	cu.mutation.RemoveContestTaskIDs(ids...)
	return cu
}

// RemoveContestTask removes "contest_task" edges to ContestTask entities.
func (cu *ContestUpdate) RemoveContestTask(c ...*ContestTask) *ContestUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveContestTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ContestUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ContestUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ContestUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ContestUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ContestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(contest.Table, contest.Columns, sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(contest.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Slug(); ok {
		_spec.SetField(contest.FieldSlug, field.TypeString, value)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(contest.FieldDescription, field.TypeString, value)
	}
	if value, ok := cu.mutation.PenaltySeconds(); ok {
		_spec.SetField(contest.FieldPenaltySeconds, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedPenaltySeconds(); ok {
		_spec.AddField(contest.FieldPenaltySeconds, field.TypeInt, value)
	}
	if value, ok := cu.mutation.ContestType(); ok {
		_spec.SetField(contest.FieldContestType, field.TypeString, value)
	}
	if value, ok := cu.mutation.IsPublic(); ok {
		_spec.SetField(contest.FieldIsPublic, field.TypeBool, value)
	}
	if value, ok := cu.mutation.StartAt(); ok {
		_spec.SetField(contest.FieldStartAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.EndAt(); ok {
		_spec.SetField(contest.FieldEndAt, field.TypeTime, value)
	}
	if cu.mutation.SubmitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.SubmitsTable,
			Columns: []string{contest.SubmitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedSubmitsIDs(); len(nodes) > 0 && !cu.mutation.SubmitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.SubmitsTable,
			Columns: []string{contest.SubmitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.SubmitsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.SubmitsTable,
			Columns: []string{contest.SubmitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   contest.UsersTable,
			Columns: contest.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !cu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   contest.UsersTable,
			Columns: contest.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   contest.UsersTable,
			Columns: contest.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   contest.TasksTable,
			Columns: contest.TasksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedTasksIDs(); len(nodes) > 0 && !cu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   contest.TasksTable,
			Columns: contest.TasksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   contest.TasksTable,
			Columns: contest.TasksPrimaryKey,
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
	if cu.mutation.ContestUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   contest.ContestUserTable,
			Columns: []string{contest.ContestUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contestuser.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedContestUserIDs(); len(nodes) > 0 && !cu.mutation.ContestUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   contest.ContestUserTable,
			Columns: []string{contest.ContestUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contestuser.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ContestUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   contest.ContestUserTable,
			Columns: []string{contest.ContestUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contestuser.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ContestTaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   contest.ContestTaskTable,
			Columns: []string{contest.ContestTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contesttask.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedContestTaskIDs(); len(nodes) > 0 && !cu.mutation.ContestTaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   contest.ContestTaskTable,
			Columns: []string{contest.ContestTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contesttask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ContestTaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   contest.ContestTaskTable,
			Columns: []string{contest.ContestTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contesttask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ContestUpdateOne is the builder for updating a single Contest entity.
type ContestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContestMutation
}

// SetName sets the "name" field.
func (cuo *ContestUpdateOne) SetName(s string) *ContestUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetSlug sets the "slug" field.
func (cuo *ContestUpdateOne) SetSlug(s string) *ContestUpdateOne {
	cuo.mutation.SetSlug(s)
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *ContestUpdateOne) SetDescription(s string) *ContestUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetPenaltySeconds sets the "penalty_seconds" field.
func (cuo *ContestUpdateOne) SetPenaltySeconds(i int) *ContestUpdateOne {
	cuo.mutation.ResetPenaltySeconds()
	cuo.mutation.SetPenaltySeconds(i)
	return cuo
}

// AddPenaltySeconds adds i to the "penalty_seconds" field.
func (cuo *ContestUpdateOne) AddPenaltySeconds(i int) *ContestUpdateOne {
	cuo.mutation.AddPenaltySeconds(i)
	return cuo
}

// SetContestType sets the "contest_type" field.
func (cuo *ContestUpdateOne) SetContestType(s string) *ContestUpdateOne {
	cuo.mutation.SetContestType(s)
	return cuo
}

// SetIsPublic sets the "is_public" field.
func (cuo *ContestUpdateOne) SetIsPublic(b bool) *ContestUpdateOne {
	cuo.mutation.SetIsPublic(b)
	return cuo
}

// SetStartAt sets the "start_at" field.
func (cuo *ContestUpdateOne) SetStartAt(t time.Time) *ContestUpdateOne {
	cuo.mutation.SetStartAt(t)
	return cuo
}

// SetEndAt sets the "end_at" field.
func (cuo *ContestUpdateOne) SetEndAt(t time.Time) *ContestUpdateOne {
	cuo.mutation.SetEndAt(t)
	return cuo
}

// AddSubmitIDs adds the "submits" edge to the Submit entity by IDs.
func (cuo *ContestUpdateOne) AddSubmitIDs(ids ...int) *ContestUpdateOne {
	cuo.mutation.AddSubmitIDs(ids...)
	return cuo
}

// AddSubmits adds the "submits" edges to the Submit entity.
func (cuo *ContestUpdateOne) AddSubmits(s ...*Submit) *ContestUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddSubmitIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (cuo *ContestUpdateOne) AddUserIDs(ids ...int) *ContestUpdateOne {
	cuo.mutation.AddUserIDs(ids...)
	return cuo
}

// AddUsers adds the "users" edges to the User entity.
func (cuo *ContestUpdateOne) AddUsers(u ...*User) *ContestUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.AddUserIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (cuo *ContestUpdateOne) AddTaskIDs(ids ...int) *ContestUpdateOne {
	cuo.mutation.AddTaskIDs(ids...)
	return cuo
}

// AddTasks adds the "tasks" edges to the Task entity.
func (cuo *ContestUpdateOne) AddTasks(t ...*Task) *ContestUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.AddTaskIDs(ids...)
}

// AddContestUserIDs adds the "contest_user" edge to the ContestUser entity by IDs.
func (cuo *ContestUpdateOne) AddContestUserIDs(ids ...int) *ContestUpdateOne {
	cuo.mutation.AddContestUserIDs(ids...)
	return cuo
}

// AddContestUser adds the "contest_user" edges to the ContestUser entity.
func (cuo *ContestUpdateOne) AddContestUser(c ...*ContestUser) *ContestUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddContestUserIDs(ids...)
}

// AddContestTaskIDs adds the "contest_task" edge to the ContestTask entity by IDs.
func (cuo *ContestUpdateOne) AddContestTaskIDs(ids ...int) *ContestUpdateOne {
	cuo.mutation.AddContestTaskIDs(ids...)
	return cuo
}

// AddContestTask adds the "contest_task" edges to the ContestTask entity.
func (cuo *ContestUpdateOne) AddContestTask(c ...*ContestTask) *ContestUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddContestTaskIDs(ids...)
}

// Mutation returns the ContestMutation object of the builder.
func (cuo *ContestUpdateOne) Mutation() *ContestMutation {
	return cuo.mutation
}

// ClearSubmits clears all "submits" edges to the Submit entity.
func (cuo *ContestUpdateOne) ClearSubmits() *ContestUpdateOne {
	cuo.mutation.ClearSubmits()
	return cuo
}

// RemoveSubmitIDs removes the "submits" edge to Submit entities by IDs.
func (cuo *ContestUpdateOne) RemoveSubmitIDs(ids ...int) *ContestUpdateOne {
	cuo.mutation.RemoveSubmitIDs(ids...)
	return cuo
}

// RemoveSubmits removes "submits" edges to Submit entities.
func (cuo *ContestUpdateOne) RemoveSubmits(s ...*Submit) *ContestUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveSubmitIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (cuo *ContestUpdateOne) ClearUsers() *ContestUpdateOne {
	cuo.mutation.ClearUsers()
	return cuo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (cuo *ContestUpdateOne) RemoveUserIDs(ids ...int) *ContestUpdateOne {
	cuo.mutation.RemoveUserIDs(ids...)
	return cuo
}

// RemoveUsers removes "users" edges to User entities.
func (cuo *ContestUpdateOne) RemoveUsers(u ...*User) *ContestUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.RemoveUserIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (cuo *ContestUpdateOne) ClearTasks() *ContestUpdateOne {
	cuo.mutation.ClearTasks()
	return cuo
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (cuo *ContestUpdateOne) RemoveTaskIDs(ids ...int) *ContestUpdateOne {
	cuo.mutation.RemoveTaskIDs(ids...)
	return cuo
}

// RemoveTasks removes "tasks" edges to Task entities.
func (cuo *ContestUpdateOne) RemoveTasks(t ...*Task) *ContestUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.RemoveTaskIDs(ids...)
}

// ClearContestUser clears all "contest_user" edges to the ContestUser entity.
func (cuo *ContestUpdateOne) ClearContestUser() *ContestUpdateOne {
	cuo.mutation.ClearContestUser()
	return cuo
}

// RemoveContestUserIDs removes the "contest_user" edge to ContestUser entities by IDs.
func (cuo *ContestUpdateOne) RemoveContestUserIDs(ids ...int) *ContestUpdateOne {
	cuo.mutation.RemoveContestUserIDs(ids...)
	return cuo
}

// RemoveContestUser removes "contest_user" edges to ContestUser entities.
func (cuo *ContestUpdateOne) RemoveContestUser(c ...*ContestUser) *ContestUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveContestUserIDs(ids...)
}

// ClearContestTask clears all "contest_task" edges to the ContestTask entity.
func (cuo *ContestUpdateOne) ClearContestTask() *ContestUpdateOne {
	cuo.mutation.ClearContestTask()
	return cuo
}

// RemoveContestTaskIDs removes the "contest_task" edge to ContestTask entities by IDs.
func (cuo *ContestUpdateOne) RemoveContestTaskIDs(ids ...int) *ContestUpdateOne {
	cuo.mutation.RemoveContestTaskIDs(ids...)
	return cuo
}

// RemoveContestTask removes "contest_task" edges to ContestTask entities.
func (cuo *ContestUpdateOne) RemoveContestTask(c ...*ContestTask) *ContestUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveContestTaskIDs(ids...)
}

// Where appends a list predicates to the ContestUpdate builder.
func (cuo *ContestUpdateOne) Where(ps ...predicate.Contest) *ContestUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ContestUpdateOne) Select(field string, fields ...string) *ContestUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Contest entity.
func (cuo *ContestUpdateOne) Save(ctx context.Context) (*Contest, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ContestUpdateOne) SaveX(ctx context.Context) *Contest {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ContestUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ContestUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ContestUpdateOne) sqlSave(ctx context.Context) (_node *Contest, err error) {
	_spec := sqlgraph.NewUpdateSpec(contest.Table, contest.Columns, sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Contest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contest.FieldID)
		for _, f := range fields {
			if !contest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != contest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(contest.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Slug(); ok {
		_spec.SetField(contest.FieldSlug, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(contest.FieldDescription, field.TypeString, value)
	}
	if value, ok := cuo.mutation.PenaltySeconds(); ok {
		_spec.SetField(contest.FieldPenaltySeconds, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedPenaltySeconds(); ok {
		_spec.AddField(contest.FieldPenaltySeconds, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.ContestType(); ok {
		_spec.SetField(contest.FieldContestType, field.TypeString, value)
	}
	if value, ok := cuo.mutation.IsPublic(); ok {
		_spec.SetField(contest.FieldIsPublic, field.TypeBool, value)
	}
	if value, ok := cuo.mutation.StartAt(); ok {
		_spec.SetField(contest.FieldStartAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.EndAt(); ok {
		_spec.SetField(contest.FieldEndAt, field.TypeTime, value)
	}
	if cuo.mutation.SubmitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.SubmitsTable,
			Columns: []string{contest.SubmitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedSubmitsIDs(); len(nodes) > 0 && !cuo.mutation.SubmitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.SubmitsTable,
			Columns: []string{contest.SubmitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.SubmitsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.SubmitsTable,
			Columns: []string{contest.SubmitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   contest.UsersTable,
			Columns: contest.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !cuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   contest.UsersTable,
			Columns: contest.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   contest.UsersTable,
			Columns: contest.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   contest.TasksTable,
			Columns: contest.TasksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !cuo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   contest.TasksTable,
			Columns: contest.TasksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   contest.TasksTable,
			Columns: contest.TasksPrimaryKey,
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
	if cuo.mutation.ContestUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   contest.ContestUserTable,
			Columns: []string{contest.ContestUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contestuser.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedContestUserIDs(); len(nodes) > 0 && !cuo.mutation.ContestUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   contest.ContestUserTable,
			Columns: []string{contest.ContestUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contestuser.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ContestUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   contest.ContestUserTable,
			Columns: []string{contest.ContestUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contestuser.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ContestTaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   contest.ContestTaskTable,
			Columns: []string{contest.ContestTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contesttask.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedContestTaskIDs(); len(nodes) > 0 && !cuo.mutation.ContestTaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   contest.ContestTaskTable,
			Columns: []string{contest.ContestTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contesttask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ContestTaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   contest.ContestTaskTable,
			Columns: []string{contest.ContestTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contesttask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Contest{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
