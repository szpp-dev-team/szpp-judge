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
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/language"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/predicate"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/submit"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcaseresult"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
)

// SubmitUpdate is the builder for updating Submit entities.
type SubmitUpdate struct {
	config
	hooks    []Hook
	mutation *SubmitMutation
}

// Where appends a list predicates to the SubmitUpdate builder.
func (su *SubmitUpdate) Where(ps ...predicate.Submit) *SubmitUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetStatus sets the "status" field.
func (su *SubmitUpdate) SetStatus(s string) *SubmitUpdate {
	su.mutation.SetStatus(s)
	return su
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableStatus(s *string) *SubmitUpdate {
	if s != nil {
		su.SetStatus(*s)
	}
	return su
}

// ClearStatus clears the value of the "status" field.
func (su *SubmitUpdate) ClearStatus() *SubmitUpdate {
	su.mutation.ClearStatus()
	return su
}

// SetExecTime sets the "exec_time" field.
func (su *SubmitUpdate) SetExecTime(i int) *SubmitUpdate {
	su.mutation.ResetExecTime()
	su.mutation.SetExecTime(i)
	return su
}

// SetNillableExecTime sets the "exec_time" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableExecTime(i *int) *SubmitUpdate {
	if i != nil {
		su.SetExecTime(*i)
	}
	return su
}

// AddExecTime adds i to the "exec_time" field.
func (su *SubmitUpdate) AddExecTime(i int) *SubmitUpdate {
	su.mutation.AddExecTime(i)
	return su
}

// ClearExecTime clears the value of the "exec_time" field.
func (su *SubmitUpdate) ClearExecTime() *SubmitUpdate {
	su.mutation.ClearExecTime()
	return su
}

// SetExecMemory sets the "exec_memory" field.
func (su *SubmitUpdate) SetExecMemory(i int) *SubmitUpdate {
	su.mutation.ResetExecMemory()
	su.mutation.SetExecMemory(i)
	return su
}

// SetNillableExecMemory sets the "exec_memory" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableExecMemory(i *int) *SubmitUpdate {
	if i != nil {
		su.SetExecMemory(*i)
	}
	return su
}

// AddExecMemory adds i to the "exec_memory" field.
func (su *SubmitUpdate) AddExecMemory(i int) *SubmitUpdate {
	su.mutation.AddExecMemory(i)
	return su
}

// ClearExecMemory clears the value of the "exec_memory" field.
func (su *SubmitUpdate) ClearExecMemory() *SubmitUpdate {
	su.mutation.ClearExecMemory()
	return su
}

// SetScore sets the "score" field.
func (su *SubmitUpdate) SetScore(i int) *SubmitUpdate {
	su.mutation.ResetScore()
	su.mutation.SetScore(i)
	return su
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableScore(i *int) *SubmitUpdate {
	if i != nil {
		su.SetScore(*i)
	}
	return su
}

// AddScore adds i to the "score" field.
func (su *SubmitUpdate) AddScore(i int) *SubmitUpdate {
	su.mutation.AddScore(i)
	return su
}

// ClearScore clears the value of the "score" field.
func (su *SubmitUpdate) ClearScore() *SubmitUpdate {
	su.mutation.ClearScore()
	return su
}

// SetSubmittedAt sets the "submitted_at" field.
func (su *SubmitUpdate) SetSubmittedAt(t time.Time) *SubmitUpdate {
	su.mutation.SetSubmittedAt(t)
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *SubmitUpdate) SetCreatedAt(t time.Time) *SubmitUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SubmitUpdate) SetUpdatedAt(t time.Time) *SubmitUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableUpdatedAt(t *time.Time) *SubmitUpdate {
	if t != nil {
		su.SetUpdatedAt(*t)
	}
	return su
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (su *SubmitUpdate) ClearUpdatedAt() *SubmitUpdate {
	su.mutation.ClearUpdatedAt()
	return su
}

// SetUserID sets the "user" edge to the User entity by ID.
func (su *SubmitUpdate) SetUserID(id int) *SubmitUpdate {
	su.mutation.SetUserID(id)
	return su
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (su *SubmitUpdate) SetNillableUserID(id *int) *SubmitUpdate {
	if id != nil {
		su = su.SetUserID(*id)
	}
	return su
}

// SetUser sets the "user" edge to the User entity.
func (su *SubmitUpdate) SetUser(u *User) *SubmitUpdate {
	return su.SetUserID(u.ID)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (su *SubmitUpdate) SetTaskID(id int) *SubmitUpdate {
	su.mutation.SetTaskID(id)
	return su
}

// SetNillableTaskID sets the "task" edge to the Task entity by ID if the given value is not nil.
func (su *SubmitUpdate) SetNillableTaskID(id *int) *SubmitUpdate {
	if id != nil {
		su = su.SetTaskID(*id)
	}
	return su
}

// SetTask sets the "task" edge to the Task entity.
func (su *SubmitUpdate) SetTask(t *Task) *SubmitUpdate {
	return su.SetTaskID(t.ID)
}

// SetLanguageID sets the "language" edge to the Language entity by ID.
func (su *SubmitUpdate) SetLanguageID(id int) *SubmitUpdate {
	su.mutation.SetLanguageID(id)
	return su
}

// SetNillableLanguageID sets the "language" edge to the Language entity by ID if the given value is not nil.
func (su *SubmitUpdate) SetNillableLanguageID(id *int) *SubmitUpdate {
	if id != nil {
		su = su.SetLanguageID(*id)
	}
	return su
}

// SetLanguage sets the "language" edge to the Language entity.
func (su *SubmitUpdate) SetLanguage(l *Language) *SubmitUpdate {
	return su.SetLanguageID(l.ID)
}

// AddTestcaseResultIDs adds the "testcase_results" edge to the TestcaseResult entity by IDs.
func (su *SubmitUpdate) AddTestcaseResultIDs(ids ...int) *SubmitUpdate {
	su.mutation.AddTestcaseResultIDs(ids...)
	return su
}

// AddTestcaseResults adds the "testcase_results" edges to the TestcaseResult entity.
func (su *SubmitUpdate) AddTestcaseResults(t ...*TestcaseResult) *SubmitUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.AddTestcaseResultIDs(ids...)
}

// Mutation returns the SubmitMutation object of the builder.
func (su *SubmitUpdate) Mutation() *SubmitMutation {
	return su.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (su *SubmitUpdate) ClearUser() *SubmitUpdate {
	su.mutation.ClearUser()
	return su
}

// ClearTask clears the "task" edge to the Task entity.
func (su *SubmitUpdate) ClearTask() *SubmitUpdate {
	su.mutation.ClearTask()
	return su
}

// ClearLanguage clears the "language" edge to the Language entity.
func (su *SubmitUpdate) ClearLanguage() *SubmitUpdate {
	su.mutation.ClearLanguage()
	return su
}

// ClearTestcaseResults clears all "testcase_results" edges to the TestcaseResult entity.
func (su *SubmitUpdate) ClearTestcaseResults() *SubmitUpdate {
	su.mutation.ClearTestcaseResults()
	return su
}

// RemoveTestcaseResultIDs removes the "testcase_results" edge to TestcaseResult entities by IDs.
func (su *SubmitUpdate) RemoveTestcaseResultIDs(ids ...int) *SubmitUpdate {
	su.mutation.RemoveTestcaseResultIDs(ids...)
	return su
}

// RemoveTestcaseResults removes "testcase_results" edges to TestcaseResult entities.
func (su *SubmitUpdate) RemoveTestcaseResults(t ...*TestcaseResult) *SubmitUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.RemoveTestcaseResultIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SubmitUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SubmitUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SubmitUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SubmitUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SubmitUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(submit.Table, submit.Columns, sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Status(); ok {
		_spec.SetField(submit.FieldStatus, field.TypeString, value)
	}
	if su.mutation.StatusCleared() {
		_spec.ClearField(submit.FieldStatus, field.TypeString)
	}
	if value, ok := su.mutation.ExecTime(); ok {
		_spec.SetField(submit.FieldExecTime, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedExecTime(); ok {
		_spec.AddField(submit.FieldExecTime, field.TypeInt, value)
	}
	if su.mutation.ExecTimeCleared() {
		_spec.ClearField(submit.FieldExecTime, field.TypeInt)
	}
	if value, ok := su.mutation.ExecMemory(); ok {
		_spec.SetField(submit.FieldExecMemory, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedExecMemory(); ok {
		_spec.AddField(submit.FieldExecMemory, field.TypeInt, value)
	}
	if su.mutation.ExecMemoryCleared() {
		_spec.ClearField(submit.FieldExecMemory, field.TypeInt)
	}
	if value, ok := su.mutation.Score(); ok {
		_spec.SetField(submit.FieldScore, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedScore(); ok {
		_spec.AddField(submit.FieldScore, field.TypeInt, value)
	}
	if su.mutation.ScoreCleared() {
		_spec.ClearField(submit.FieldScore, field.TypeInt)
	}
	if value, ok := su.mutation.SubmittedAt(); ok {
		_spec.SetField(submit.FieldSubmittedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.SetField(submit.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(submit.FieldUpdatedAt, field.TypeTime, value)
	}
	if su.mutation.UpdatedAtCleared() {
		_spec.ClearField(submit.FieldUpdatedAt, field.TypeTime)
	}
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submit.UserTable,
			Columns: []string{submit.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submit.UserTable,
			Columns: []string{submit.UserColumn},
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
	if su.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submit.TaskTable,
			Columns: []string{submit.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submit.TaskTable,
			Columns: []string{submit.TaskColumn},
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
	if su.mutation.LanguageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submit.LanguageTable,
			Columns: []string{submit.LanguageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.LanguageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submit.LanguageTable,
			Columns: []string{submit.LanguageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.TestcaseResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   submit.TestcaseResultsTable,
			Columns: []string{submit.TestcaseResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcaseresult.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedTestcaseResultsIDs(); len(nodes) > 0 && !su.mutation.TestcaseResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   submit.TestcaseResultsTable,
			Columns: []string{submit.TestcaseResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcaseresult.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.TestcaseResultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   submit.TestcaseResultsTable,
			Columns: []string{submit.TestcaseResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcaseresult.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{submit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SubmitUpdateOne is the builder for updating a single Submit entity.
type SubmitUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubmitMutation
}

// SetStatus sets the "status" field.
func (suo *SubmitUpdateOne) SetStatus(s string) *SubmitUpdateOne {
	suo.mutation.SetStatus(s)
	return suo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableStatus(s *string) *SubmitUpdateOne {
	if s != nil {
		suo.SetStatus(*s)
	}
	return suo
}

// ClearStatus clears the value of the "status" field.
func (suo *SubmitUpdateOne) ClearStatus() *SubmitUpdateOne {
	suo.mutation.ClearStatus()
	return suo
}

// SetExecTime sets the "exec_time" field.
func (suo *SubmitUpdateOne) SetExecTime(i int) *SubmitUpdateOne {
	suo.mutation.ResetExecTime()
	suo.mutation.SetExecTime(i)
	return suo
}

// SetNillableExecTime sets the "exec_time" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableExecTime(i *int) *SubmitUpdateOne {
	if i != nil {
		suo.SetExecTime(*i)
	}
	return suo
}

// AddExecTime adds i to the "exec_time" field.
func (suo *SubmitUpdateOne) AddExecTime(i int) *SubmitUpdateOne {
	suo.mutation.AddExecTime(i)
	return suo
}

// ClearExecTime clears the value of the "exec_time" field.
func (suo *SubmitUpdateOne) ClearExecTime() *SubmitUpdateOne {
	suo.mutation.ClearExecTime()
	return suo
}

// SetExecMemory sets the "exec_memory" field.
func (suo *SubmitUpdateOne) SetExecMemory(i int) *SubmitUpdateOne {
	suo.mutation.ResetExecMemory()
	suo.mutation.SetExecMemory(i)
	return suo
}

// SetNillableExecMemory sets the "exec_memory" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableExecMemory(i *int) *SubmitUpdateOne {
	if i != nil {
		suo.SetExecMemory(*i)
	}
	return suo
}

// AddExecMemory adds i to the "exec_memory" field.
func (suo *SubmitUpdateOne) AddExecMemory(i int) *SubmitUpdateOne {
	suo.mutation.AddExecMemory(i)
	return suo
}

// ClearExecMemory clears the value of the "exec_memory" field.
func (suo *SubmitUpdateOne) ClearExecMemory() *SubmitUpdateOne {
	suo.mutation.ClearExecMemory()
	return suo
}

// SetScore sets the "score" field.
func (suo *SubmitUpdateOne) SetScore(i int) *SubmitUpdateOne {
	suo.mutation.ResetScore()
	suo.mutation.SetScore(i)
	return suo
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableScore(i *int) *SubmitUpdateOne {
	if i != nil {
		suo.SetScore(*i)
	}
	return suo
}

// AddScore adds i to the "score" field.
func (suo *SubmitUpdateOne) AddScore(i int) *SubmitUpdateOne {
	suo.mutation.AddScore(i)
	return suo
}

// ClearScore clears the value of the "score" field.
func (suo *SubmitUpdateOne) ClearScore() *SubmitUpdateOne {
	suo.mutation.ClearScore()
	return suo
}

// SetSubmittedAt sets the "submitted_at" field.
func (suo *SubmitUpdateOne) SetSubmittedAt(t time.Time) *SubmitUpdateOne {
	suo.mutation.SetSubmittedAt(t)
	return suo
}

// SetCreatedAt sets the "created_at" field.
func (suo *SubmitUpdateOne) SetCreatedAt(t time.Time) *SubmitUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SubmitUpdateOne) SetUpdatedAt(t time.Time) *SubmitUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableUpdatedAt(t *time.Time) *SubmitUpdateOne {
	if t != nil {
		suo.SetUpdatedAt(*t)
	}
	return suo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (suo *SubmitUpdateOne) ClearUpdatedAt() *SubmitUpdateOne {
	suo.mutation.ClearUpdatedAt()
	return suo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (suo *SubmitUpdateOne) SetUserID(id int) *SubmitUpdateOne {
	suo.mutation.SetUserID(id)
	return suo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableUserID(id *int) *SubmitUpdateOne {
	if id != nil {
		suo = suo.SetUserID(*id)
	}
	return suo
}

// SetUser sets the "user" edge to the User entity.
func (suo *SubmitUpdateOne) SetUser(u *User) *SubmitUpdateOne {
	return suo.SetUserID(u.ID)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (suo *SubmitUpdateOne) SetTaskID(id int) *SubmitUpdateOne {
	suo.mutation.SetTaskID(id)
	return suo
}

// SetNillableTaskID sets the "task" edge to the Task entity by ID if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableTaskID(id *int) *SubmitUpdateOne {
	if id != nil {
		suo = suo.SetTaskID(*id)
	}
	return suo
}

// SetTask sets the "task" edge to the Task entity.
func (suo *SubmitUpdateOne) SetTask(t *Task) *SubmitUpdateOne {
	return suo.SetTaskID(t.ID)
}

// SetLanguageID sets the "language" edge to the Language entity by ID.
func (suo *SubmitUpdateOne) SetLanguageID(id int) *SubmitUpdateOne {
	suo.mutation.SetLanguageID(id)
	return suo
}

// SetNillableLanguageID sets the "language" edge to the Language entity by ID if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableLanguageID(id *int) *SubmitUpdateOne {
	if id != nil {
		suo = suo.SetLanguageID(*id)
	}
	return suo
}

// SetLanguage sets the "language" edge to the Language entity.
func (suo *SubmitUpdateOne) SetLanguage(l *Language) *SubmitUpdateOne {
	return suo.SetLanguageID(l.ID)
}

// AddTestcaseResultIDs adds the "testcase_results" edge to the TestcaseResult entity by IDs.
func (suo *SubmitUpdateOne) AddTestcaseResultIDs(ids ...int) *SubmitUpdateOne {
	suo.mutation.AddTestcaseResultIDs(ids...)
	return suo
}

// AddTestcaseResults adds the "testcase_results" edges to the TestcaseResult entity.
func (suo *SubmitUpdateOne) AddTestcaseResults(t ...*TestcaseResult) *SubmitUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.AddTestcaseResultIDs(ids...)
}

// Mutation returns the SubmitMutation object of the builder.
func (suo *SubmitUpdateOne) Mutation() *SubmitMutation {
	return suo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (suo *SubmitUpdateOne) ClearUser() *SubmitUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// ClearTask clears the "task" edge to the Task entity.
func (suo *SubmitUpdateOne) ClearTask() *SubmitUpdateOne {
	suo.mutation.ClearTask()
	return suo
}

// ClearLanguage clears the "language" edge to the Language entity.
func (suo *SubmitUpdateOne) ClearLanguage() *SubmitUpdateOne {
	suo.mutation.ClearLanguage()
	return suo
}

// ClearTestcaseResults clears all "testcase_results" edges to the TestcaseResult entity.
func (suo *SubmitUpdateOne) ClearTestcaseResults() *SubmitUpdateOne {
	suo.mutation.ClearTestcaseResults()
	return suo
}

// RemoveTestcaseResultIDs removes the "testcase_results" edge to TestcaseResult entities by IDs.
func (suo *SubmitUpdateOne) RemoveTestcaseResultIDs(ids ...int) *SubmitUpdateOne {
	suo.mutation.RemoveTestcaseResultIDs(ids...)
	return suo
}

// RemoveTestcaseResults removes "testcase_results" edges to TestcaseResult entities.
func (suo *SubmitUpdateOne) RemoveTestcaseResults(t ...*TestcaseResult) *SubmitUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.RemoveTestcaseResultIDs(ids...)
}

// Where appends a list predicates to the SubmitUpdate builder.
func (suo *SubmitUpdateOne) Where(ps ...predicate.Submit) *SubmitUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SubmitUpdateOne) Select(field string, fields ...string) *SubmitUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Submit entity.
func (suo *SubmitUpdateOne) Save(ctx context.Context) (*Submit, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SubmitUpdateOne) SaveX(ctx context.Context) *Submit {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SubmitUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SubmitUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SubmitUpdateOne) sqlSave(ctx context.Context) (_node *Submit, err error) {
	_spec := sqlgraph.NewUpdateSpec(submit.Table, submit.Columns, sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Submit.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, submit.FieldID)
		for _, f := range fields {
			if !submit.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != submit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Status(); ok {
		_spec.SetField(submit.FieldStatus, field.TypeString, value)
	}
	if suo.mutation.StatusCleared() {
		_spec.ClearField(submit.FieldStatus, field.TypeString)
	}
	if value, ok := suo.mutation.ExecTime(); ok {
		_spec.SetField(submit.FieldExecTime, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedExecTime(); ok {
		_spec.AddField(submit.FieldExecTime, field.TypeInt, value)
	}
	if suo.mutation.ExecTimeCleared() {
		_spec.ClearField(submit.FieldExecTime, field.TypeInt)
	}
	if value, ok := suo.mutation.ExecMemory(); ok {
		_spec.SetField(submit.FieldExecMemory, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedExecMemory(); ok {
		_spec.AddField(submit.FieldExecMemory, field.TypeInt, value)
	}
	if suo.mutation.ExecMemoryCleared() {
		_spec.ClearField(submit.FieldExecMemory, field.TypeInt)
	}
	if value, ok := suo.mutation.Score(); ok {
		_spec.SetField(submit.FieldScore, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedScore(); ok {
		_spec.AddField(submit.FieldScore, field.TypeInt, value)
	}
	if suo.mutation.ScoreCleared() {
		_spec.ClearField(submit.FieldScore, field.TypeInt)
	}
	if value, ok := suo.mutation.SubmittedAt(); ok {
		_spec.SetField(submit.FieldSubmittedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.SetField(submit.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(submit.FieldUpdatedAt, field.TypeTime, value)
	}
	if suo.mutation.UpdatedAtCleared() {
		_spec.ClearField(submit.FieldUpdatedAt, field.TypeTime)
	}
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submit.UserTable,
			Columns: []string{submit.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submit.UserTable,
			Columns: []string{submit.UserColumn},
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
	if suo.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submit.TaskTable,
			Columns: []string{submit.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submit.TaskTable,
			Columns: []string{submit.TaskColumn},
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
	if suo.mutation.LanguageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submit.LanguageTable,
			Columns: []string{submit.LanguageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.LanguageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submit.LanguageTable,
			Columns: []string{submit.LanguageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.TestcaseResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   submit.TestcaseResultsTable,
			Columns: []string{submit.TestcaseResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcaseresult.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedTestcaseResultsIDs(); len(nodes) > 0 && !suo.mutation.TestcaseResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   submit.TestcaseResultsTable,
			Columns: []string{submit.TestcaseResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcaseresult.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.TestcaseResultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   submit.TestcaseResultsTable,
			Columns: []string{submit.TestcaseResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcaseresult.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Submit{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{submit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}