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
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/submit"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcase"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcaseset"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTitle sets the "title" field.
func (tc *TaskCreate) SetTitle(s string) *TaskCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// SetStatement sets the "statement" field.
func (tc *TaskCreate) SetStatement(s string) *TaskCreate {
	tc.mutation.SetStatement(s)
	return tc
}

// SetDifficulty sets the "difficulty" field.
func (tc *TaskCreate) SetDifficulty(s string) *TaskCreate {
	tc.mutation.SetDifficulty(s)
	return tc
}

// SetExecTimeLimit sets the "exec_time_limit" field.
func (tc *TaskCreate) SetExecTimeLimit(u uint) *TaskCreate {
	tc.mutation.SetExecTimeLimit(u)
	return tc
}

// SetExecMemoryLimit sets the "exec_memory_limit" field.
func (tc *TaskCreate) SetExecMemoryLimit(u uint) *TaskCreate {
	tc.mutation.SetExecMemoryLimit(u)
	return tc
}

// SetJudgeType sets the "judge_type" field.
func (tc *TaskCreate) SetJudgeType(tt task.JudgeType) *TaskCreate {
	tc.mutation.SetJudgeType(tt)
	return tc
}

// SetCaseInsensitive sets the "case_insensitive" field.
func (tc *TaskCreate) SetCaseInsensitive(b bool) *TaskCreate {
	tc.mutation.SetCaseInsensitive(b)
	return tc
}

// SetNillableCaseInsensitive sets the "case_insensitive" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCaseInsensitive(b *bool) *TaskCreate {
	if b != nil {
		tc.SetCaseInsensitive(*b)
	}
	return tc
}

// SetNdigits sets the "ndigits" field.
func (tc *TaskCreate) SetNdigits(u uint) *TaskCreate {
	tc.mutation.SetNdigits(u)
	return tc
}

// SetNillableNdigits sets the "ndigits" field if the given value is not nil.
func (tc *TaskCreate) SetNillableNdigits(u *uint) *TaskCreate {
	if u != nil {
		tc.SetNdigits(*u)
	}
	return tc
}

// SetJudgeCodePath sets the "judge_code_path" field.
func (tc *TaskCreate) SetJudgeCodePath(s string) *TaskCreate {
	tc.mutation.SetJudgeCodePath(s)
	return tc
}

// SetNillableJudgeCodePath sets the "judge_code_path" field if the given value is not nil.
func (tc *TaskCreate) SetNillableJudgeCodePath(s *string) *TaskCreate {
	if s != nil {
		tc.SetJudgeCodePath(*s)
	}
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TaskCreate) SetCreatedAt(t time.Time) *TaskCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TaskCreate) SetUpdatedAt(t time.Time) *TaskCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableUpdatedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TaskCreate) SetID(i int) *TaskCreate {
	tc.mutation.SetID(i)
	return tc
}

// AddTestcaseSetIDs adds the "testcase_sets" edge to the TestcaseSet entity by IDs.
func (tc *TaskCreate) AddTestcaseSetIDs(ids ...int) *TaskCreate {
	tc.mutation.AddTestcaseSetIDs(ids...)
	return tc
}

// AddTestcaseSets adds the "testcase_sets" edges to the TestcaseSet entity.
func (tc *TaskCreate) AddTestcaseSets(t ...*TestcaseSet) *TaskCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTestcaseSetIDs(ids...)
}

// AddTestcaseIDs adds the "testcases" edge to the Testcase entity by IDs.
func (tc *TaskCreate) AddTestcaseIDs(ids ...int) *TaskCreate {
	tc.mutation.AddTestcaseIDs(ids...)
	return tc
}

// AddTestcases adds the "testcases" edges to the Testcase entity.
func (tc *TaskCreate) AddTestcases(t ...*Testcase) *TaskCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTestcaseIDs(ids...)
}

// AddSubmitIDs adds the "submits" edge to the Submit entity by IDs.
func (tc *TaskCreate) AddSubmitIDs(ids ...int) *TaskCreate {
	tc.mutation.AddSubmitIDs(ids...)
	return tc
}

// AddSubmits adds the "submits" edges to the Submit entity.
func (tc *TaskCreate) AddSubmits(s ...*Submit) *TaskCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tc.AddSubmitIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tc *TaskCreate) SetUserID(id int) *TaskCreate {
	tc.mutation.SetUserID(id)
	return tc
}

// SetUser sets the "user" edge to the User entity.
func (tc *TaskCreate) SetUser(u *User) *TaskCreate {
	return tc.SetUserID(u.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tc *TaskCreate) Mutation() *TaskMutation {
	return tc.mutation
}

// Save creates the Task in the database.
func (tc *TaskCreate) Save(ctx context.Context) (*Task, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TaskCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TaskCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaskCreate) check() error {
	if _, ok := tc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Task.title"`)}
	}
	if _, ok := tc.mutation.Statement(); !ok {
		return &ValidationError{Name: "statement", err: errors.New(`ent: missing required field "Task.statement"`)}
	}
	if _, ok := tc.mutation.Difficulty(); !ok {
		return &ValidationError{Name: "difficulty", err: errors.New(`ent: missing required field "Task.difficulty"`)}
	}
	if _, ok := tc.mutation.ExecTimeLimit(); !ok {
		return &ValidationError{Name: "exec_time_limit", err: errors.New(`ent: missing required field "Task.exec_time_limit"`)}
	}
	if _, ok := tc.mutation.ExecMemoryLimit(); !ok {
		return &ValidationError{Name: "exec_memory_limit", err: errors.New(`ent: missing required field "Task.exec_memory_limit"`)}
	}
	if _, ok := tc.mutation.JudgeType(); !ok {
		return &ValidationError{Name: "judge_type", err: errors.New(`ent: missing required field "Task.judge_type"`)}
	}
	if v, ok := tc.mutation.JudgeType(); ok {
		if err := task.JudgeTypeValidator(v); err != nil {
			return &ValidationError{Name: "judge_type", err: fmt.Errorf(`ent: validator failed for field "Task.judge_type": %w`, err)}
		}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Task.created_at"`)}
	}
	if _, ok := tc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Task.user"`)}
	}
	return nil
}

func (tc *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
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

func (tc *TaskCreate) createSpec() (*Task, *sqlgraph.CreateSpec) {
	var (
		_node = &Task{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(task.Table, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	)
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Title(); ok {
		_spec.SetField(task.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := tc.mutation.Statement(); ok {
		_spec.SetField(task.FieldStatement, field.TypeString, value)
		_node.Statement = value
	}
	if value, ok := tc.mutation.Difficulty(); ok {
		_spec.SetField(task.FieldDifficulty, field.TypeString, value)
		_node.Difficulty = value
	}
	if value, ok := tc.mutation.ExecTimeLimit(); ok {
		_spec.SetField(task.FieldExecTimeLimit, field.TypeUint, value)
		_node.ExecTimeLimit = value
	}
	if value, ok := tc.mutation.ExecMemoryLimit(); ok {
		_spec.SetField(task.FieldExecMemoryLimit, field.TypeUint, value)
		_node.ExecMemoryLimit = value
	}
	if value, ok := tc.mutation.JudgeType(); ok {
		_spec.SetField(task.FieldJudgeType, field.TypeEnum, value)
		_node.JudgeType = value
	}
	if value, ok := tc.mutation.CaseInsensitive(); ok {
		_spec.SetField(task.FieldCaseInsensitive, field.TypeBool, value)
		_node.CaseInsensitive = &value
	}
	if value, ok := tc.mutation.Ndigits(); ok {
		_spec.SetField(task.FieldNdigits, field.TypeUint, value)
		_node.Ndigits = &value
	}
	if value, ok := tc.mutation.JudgeCodePath(); ok {
		_spec.SetField(task.FieldJudgeCodePath, field.TypeString, value)
		_node.JudgeCodePath = &value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(task.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(task.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if nodes := tc.mutation.TestcaseSetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.TestcaseSetsTable,
			Columns: []string{task.TestcaseSetsColumn},
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
	if nodes := tc.mutation.TestcasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.TestcasesTable,
			Columns: []string{task.TestcasesColumn},
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
	if nodes := tc.mutation.SubmitsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.SubmitsTable,
			Columns: []string{task.SubmitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.UserTable,
			Columns: []string{task.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_tasks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Task.Create().
//		SetTitle(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
func (tc *TaskCreate) OnConflict(opts ...sql.ConflictOption) *TaskUpsertOne {
	tc.conflict = opts
	return &TaskUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TaskCreate) OnConflictColumns(columns ...string) *TaskUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TaskUpsertOne{
		create: tc,
	}
}

type (
	// TaskUpsertOne is the builder for "upsert"-ing
	//  one Task node.
	TaskUpsertOne struct {
		create *TaskCreate
	}

	// TaskUpsert is the "OnConflict" setter.
	TaskUpsert struct {
		*sql.UpdateSet
	}
)

// SetTitle sets the "title" field.
func (u *TaskUpsert) SetTitle(v string) *TaskUpsert {
	u.Set(task.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *TaskUpsert) UpdateTitle() *TaskUpsert {
	u.SetExcluded(task.FieldTitle)
	return u
}

// SetStatement sets the "statement" field.
func (u *TaskUpsert) SetStatement(v string) *TaskUpsert {
	u.Set(task.FieldStatement, v)
	return u
}

// UpdateStatement sets the "statement" field to the value that was provided on create.
func (u *TaskUpsert) UpdateStatement() *TaskUpsert {
	u.SetExcluded(task.FieldStatement)
	return u
}

// SetDifficulty sets the "difficulty" field.
func (u *TaskUpsert) SetDifficulty(v string) *TaskUpsert {
	u.Set(task.FieldDifficulty, v)
	return u
}

// UpdateDifficulty sets the "difficulty" field to the value that was provided on create.
func (u *TaskUpsert) UpdateDifficulty() *TaskUpsert {
	u.SetExcluded(task.FieldDifficulty)
	return u
}

// SetExecTimeLimit sets the "exec_time_limit" field.
func (u *TaskUpsert) SetExecTimeLimit(v uint) *TaskUpsert {
	u.Set(task.FieldExecTimeLimit, v)
	return u
}

// UpdateExecTimeLimit sets the "exec_time_limit" field to the value that was provided on create.
func (u *TaskUpsert) UpdateExecTimeLimit() *TaskUpsert {
	u.SetExcluded(task.FieldExecTimeLimit)
	return u
}

// AddExecTimeLimit adds v to the "exec_time_limit" field.
func (u *TaskUpsert) AddExecTimeLimit(v uint) *TaskUpsert {
	u.Add(task.FieldExecTimeLimit, v)
	return u
}

// SetExecMemoryLimit sets the "exec_memory_limit" field.
func (u *TaskUpsert) SetExecMemoryLimit(v uint) *TaskUpsert {
	u.Set(task.FieldExecMemoryLimit, v)
	return u
}

// UpdateExecMemoryLimit sets the "exec_memory_limit" field to the value that was provided on create.
func (u *TaskUpsert) UpdateExecMemoryLimit() *TaskUpsert {
	u.SetExcluded(task.FieldExecMemoryLimit)
	return u
}

// AddExecMemoryLimit adds v to the "exec_memory_limit" field.
func (u *TaskUpsert) AddExecMemoryLimit(v uint) *TaskUpsert {
	u.Add(task.FieldExecMemoryLimit, v)
	return u
}

// SetJudgeType sets the "judge_type" field.
func (u *TaskUpsert) SetJudgeType(v task.JudgeType) *TaskUpsert {
	u.Set(task.FieldJudgeType, v)
	return u
}

// UpdateJudgeType sets the "judge_type" field to the value that was provided on create.
func (u *TaskUpsert) UpdateJudgeType() *TaskUpsert {
	u.SetExcluded(task.FieldJudgeType)
	return u
}

// SetCaseInsensitive sets the "case_insensitive" field.
func (u *TaskUpsert) SetCaseInsensitive(v bool) *TaskUpsert {
	u.Set(task.FieldCaseInsensitive, v)
	return u
}

// UpdateCaseInsensitive sets the "case_insensitive" field to the value that was provided on create.
func (u *TaskUpsert) UpdateCaseInsensitive() *TaskUpsert {
	u.SetExcluded(task.FieldCaseInsensitive)
	return u
}

// ClearCaseInsensitive clears the value of the "case_insensitive" field.
func (u *TaskUpsert) ClearCaseInsensitive() *TaskUpsert {
	u.SetNull(task.FieldCaseInsensitive)
	return u
}

// SetNdigits sets the "ndigits" field.
func (u *TaskUpsert) SetNdigits(v uint) *TaskUpsert {
	u.Set(task.FieldNdigits, v)
	return u
}

// UpdateNdigits sets the "ndigits" field to the value that was provided on create.
func (u *TaskUpsert) UpdateNdigits() *TaskUpsert {
	u.SetExcluded(task.FieldNdigits)
	return u
}

// AddNdigits adds v to the "ndigits" field.
func (u *TaskUpsert) AddNdigits(v uint) *TaskUpsert {
	u.Add(task.FieldNdigits, v)
	return u
}

// ClearNdigits clears the value of the "ndigits" field.
func (u *TaskUpsert) ClearNdigits() *TaskUpsert {
	u.SetNull(task.FieldNdigits)
	return u
}

// SetJudgeCodePath sets the "judge_code_path" field.
func (u *TaskUpsert) SetJudgeCodePath(v string) *TaskUpsert {
	u.Set(task.FieldJudgeCodePath, v)
	return u
}

// UpdateJudgeCodePath sets the "judge_code_path" field to the value that was provided on create.
func (u *TaskUpsert) UpdateJudgeCodePath() *TaskUpsert {
	u.SetExcluded(task.FieldJudgeCodePath)
	return u
}

// ClearJudgeCodePath clears the value of the "judge_code_path" field.
func (u *TaskUpsert) ClearJudgeCodePath() *TaskUpsert {
	u.SetNull(task.FieldJudgeCodePath)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskUpsert) SetCreatedAt(v time.Time) *TaskUpsert {
	u.Set(task.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskUpsert) UpdateCreatedAt() *TaskUpsert {
	u.SetExcluded(task.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskUpsert) SetUpdatedAt(v time.Time) *TaskUpsert {
	u.Set(task.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskUpsert) UpdateUpdatedAt() *TaskUpsert {
	u.SetExcluded(task.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *TaskUpsert) ClearUpdatedAt() *TaskUpsert {
	u.SetNull(task.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(task.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TaskUpsertOne) UpdateNewValues() *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(task.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Task.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TaskUpsertOne) Ignore() *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskUpsertOne) DoNothing() *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskCreate.OnConflict
// documentation for more info.
func (u *TaskUpsertOne) Update(set func(*TaskUpsert)) *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *TaskUpsertOne) SetTitle(v string) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateTitle() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateTitle()
	})
}

// SetStatement sets the "statement" field.
func (u *TaskUpsertOne) SetStatement(v string) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetStatement(v)
	})
}

// UpdateStatement sets the "statement" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateStatement() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateStatement()
	})
}

// SetDifficulty sets the "difficulty" field.
func (u *TaskUpsertOne) SetDifficulty(v string) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetDifficulty(v)
	})
}

// UpdateDifficulty sets the "difficulty" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateDifficulty() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateDifficulty()
	})
}

// SetExecTimeLimit sets the "exec_time_limit" field.
func (u *TaskUpsertOne) SetExecTimeLimit(v uint) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetExecTimeLimit(v)
	})
}

// AddExecTimeLimit adds v to the "exec_time_limit" field.
func (u *TaskUpsertOne) AddExecTimeLimit(v uint) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.AddExecTimeLimit(v)
	})
}

// UpdateExecTimeLimit sets the "exec_time_limit" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateExecTimeLimit() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateExecTimeLimit()
	})
}

// SetExecMemoryLimit sets the "exec_memory_limit" field.
func (u *TaskUpsertOne) SetExecMemoryLimit(v uint) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetExecMemoryLimit(v)
	})
}

// AddExecMemoryLimit adds v to the "exec_memory_limit" field.
func (u *TaskUpsertOne) AddExecMemoryLimit(v uint) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.AddExecMemoryLimit(v)
	})
}

// UpdateExecMemoryLimit sets the "exec_memory_limit" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateExecMemoryLimit() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateExecMemoryLimit()
	})
}

// SetJudgeType sets the "judge_type" field.
func (u *TaskUpsertOne) SetJudgeType(v task.JudgeType) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetJudgeType(v)
	})
}

// UpdateJudgeType sets the "judge_type" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateJudgeType() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateJudgeType()
	})
}

// SetCaseInsensitive sets the "case_insensitive" field.
func (u *TaskUpsertOne) SetCaseInsensitive(v bool) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetCaseInsensitive(v)
	})
}

// UpdateCaseInsensitive sets the "case_insensitive" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateCaseInsensitive() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateCaseInsensitive()
	})
}

// ClearCaseInsensitive clears the value of the "case_insensitive" field.
func (u *TaskUpsertOne) ClearCaseInsensitive() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.ClearCaseInsensitive()
	})
}

// SetNdigits sets the "ndigits" field.
func (u *TaskUpsertOne) SetNdigits(v uint) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetNdigits(v)
	})
}

// AddNdigits adds v to the "ndigits" field.
func (u *TaskUpsertOne) AddNdigits(v uint) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.AddNdigits(v)
	})
}

// UpdateNdigits sets the "ndigits" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateNdigits() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateNdigits()
	})
}

// ClearNdigits clears the value of the "ndigits" field.
func (u *TaskUpsertOne) ClearNdigits() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.ClearNdigits()
	})
}

// SetJudgeCodePath sets the "judge_code_path" field.
func (u *TaskUpsertOne) SetJudgeCodePath(v string) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetJudgeCodePath(v)
	})
}

// UpdateJudgeCodePath sets the "judge_code_path" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateJudgeCodePath() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateJudgeCodePath()
	})
}

// ClearJudgeCodePath clears the value of the "judge_code_path" field.
func (u *TaskUpsertOne) ClearJudgeCodePath() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.ClearJudgeCodePath()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskUpsertOne) SetCreatedAt(v time.Time) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateCreatedAt() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskUpsertOne) SetUpdatedAt(v time.Time) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateUpdatedAt() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *TaskUpsertOne) ClearUpdatedAt() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.ClearUpdatedAt()
	})
}

// Exec executes the query.
func (u *TaskUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TaskUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TaskUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	builders []*TaskCreate
	conflict []sql.ConflictOption
}

// Save creates the Task entities in the database.
func (tcb *TaskCreateBulk) Save(ctx context.Context) ([]*Task, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Task, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskMutation)
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
func (tcb *TaskCreateBulk) SaveX(ctx context.Context) []*Task {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TaskCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TaskCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Task.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
func (tcb *TaskCreateBulk) OnConflict(opts ...sql.ConflictOption) *TaskUpsertBulk {
	tcb.conflict = opts
	return &TaskUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TaskCreateBulk) OnConflictColumns(columns ...string) *TaskUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TaskUpsertBulk{
		create: tcb,
	}
}

// TaskUpsertBulk is the builder for "upsert"-ing
// a bulk of Task nodes.
type TaskUpsertBulk struct {
	create *TaskCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(task.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TaskUpsertBulk) UpdateNewValues() *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(task.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TaskUpsertBulk) Ignore() *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskUpsertBulk) DoNothing() *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskCreateBulk.OnConflict
// documentation for more info.
func (u *TaskUpsertBulk) Update(set func(*TaskUpsert)) *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *TaskUpsertBulk) SetTitle(v string) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateTitle() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateTitle()
	})
}

// SetStatement sets the "statement" field.
func (u *TaskUpsertBulk) SetStatement(v string) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetStatement(v)
	})
}

// UpdateStatement sets the "statement" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateStatement() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateStatement()
	})
}

// SetDifficulty sets the "difficulty" field.
func (u *TaskUpsertBulk) SetDifficulty(v string) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetDifficulty(v)
	})
}

// UpdateDifficulty sets the "difficulty" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateDifficulty() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateDifficulty()
	})
}

// SetExecTimeLimit sets the "exec_time_limit" field.
func (u *TaskUpsertBulk) SetExecTimeLimit(v uint) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetExecTimeLimit(v)
	})
}

// AddExecTimeLimit adds v to the "exec_time_limit" field.
func (u *TaskUpsertBulk) AddExecTimeLimit(v uint) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.AddExecTimeLimit(v)
	})
}

// UpdateExecTimeLimit sets the "exec_time_limit" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateExecTimeLimit() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateExecTimeLimit()
	})
}

// SetExecMemoryLimit sets the "exec_memory_limit" field.
func (u *TaskUpsertBulk) SetExecMemoryLimit(v uint) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetExecMemoryLimit(v)
	})
}

// AddExecMemoryLimit adds v to the "exec_memory_limit" field.
func (u *TaskUpsertBulk) AddExecMemoryLimit(v uint) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.AddExecMemoryLimit(v)
	})
}

// UpdateExecMemoryLimit sets the "exec_memory_limit" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateExecMemoryLimit() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateExecMemoryLimit()
	})
}

// SetJudgeType sets the "judge_type" field.
func (u *TaskUpsertBulk) SetJudgeType(v task.JudgeType) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetJudgeType(v)
	})
}

// UpdateJudgeType sets the "judge_type" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateJudgeType() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateJudgeType()
	})
}

// SetCaseInsensitive sets the "case_insensitive" field.
func (u *TaskUpsertBulk) SetCaseInsensitive(v bool) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetCaseInsensitive(v)
	})
}

// UpdateCaseInsensitive sets the "case_insensitive" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateCaseInsensitive() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateCaseInsensitive()
	})
}

// ClearCaseInsensitive clears the value of the "case_insensitive" field.
func (u *TaskUpsertBulk) ClearCaseInsensitive() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.ClearCaseInsensitive()
	})
}

// SetNdigits sets the "ndigits" field.
func (u *TaskUpsertBulk) SetNdigits(v uint) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetNdigits(v)
	})
}

// AddNdigits adds v to the "ndigits" field.
func (u *TaskUpsertBulk) AddNdigits(v uint) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.AddNdigits(v)
	})
}

// UpdateNdigits sets the "ndigits" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateNdigits() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateNdigits()
	})
}

// ClearNdigits clears the value of the "ndigits" field.
func (u *TaskUpsertBulk) ClearNdigits() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.ClearNdigits()
	})
}

// SetJudgeCodePath sets the "judge_code_path" field.
func (u *TaskUpsertBulk) SetJudgeCodePath(v string) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetJudgeCodePath(v)
	})
}

// UpdateJudgeCodePath sets the "judge_code_path" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateJudgeCodePath() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateJudgeCodePath()
	})
}

// ClearJudgeCodePath clears the value of the "judge_code_path" field.
func (u *TaskUpsertBulk) ClearJudgeCodePath() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.ClearJudgeCodePath()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskUpsertBulk) SetCreatedAt(v time.Time) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateCreatedAt() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskUpsertBulk) SetUpdatedAt(v time.Time) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateUpdatedAt() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *TaskUpsertBulk) ClearUpdatedAt() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.ClearUpdatedAt()
	})
}

// Exec executes the query.
func (u *TaskUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TaskCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
