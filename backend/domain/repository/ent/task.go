// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
)

// Task is the model entity for the Task schema.
type Task struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Statement holds the value of the "statement" field.
	Statement string `json:"statement,omitempty"`
	// Difficulty holds the value of the "difficulty" field.
	Difficulty string `json:"difficulty,omitempty"`
	// ExecTimeLimit holds the value of the "exec_time_limit" field.
	ExecTimeLimit uint `json:"exec_time_limit,omitempty"`
	// ExecMemoryLimit holds the value of the "exec_memory_limit" field.
	ExecMemoryLimit uint `json:"exec_memory_limit,omitempty"`
	// JudgeType holds the value of the "judge_type" field.
	JudgeType task.JudgeType `json:"judge_type,omitempty"`
	// CaseInsensitive holds the value of the "case_insensitive" field.
	CaseInsensitive *bool `json:"case_insensitive,omitempty"`
	// Ndigits holds the value of the "ndigits" field.
	Ndigits *uint `json:"ndigits,omitempty"`
	// JudgeCodePath holds the value of the "judge_code_path" field.
	JudgeCodePath *string `json:"judge_code_path,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskQuery when eager-loading is set.
	Edges        TaskEdges `json:"edges"`
	user_tasks   *int
	selectValues sql.SelectValues
}

// TaskEdges holds the relations/edges for other nodes in the graph.
type TaskEdges struct {
	// TestcaseSets holds the value of the testcase_sets edge.
	TestcaseSets []*TestcaseSet `json:"testcase_sets,omitempty"`
	// Testcases holds the value of the testcases edge.
	Testcases []*Testcase `json:"testcases,omitempty"`
	// Submits holds the value of the submits edge.
	Submits []*Submit `json:"submits,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// ContestTask holds the value of the contest_task edge.
	ContestTask []*ContestTask `json:"contest_task,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// TestcaseSetsOrErr returns the TestcaseSets value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) TestcaseSetsOrErr() ([]*TestcaseSet, error) {
	if e.loadedTypes[0] {
		return e.TestcaseSets, nil
	}
	return nil, &NotLoadedError{edge: "testcase_sets"}
}

// TestcasesOrErr returns the Testcases value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) TestcasesOrErr() ([]*Testcase, error) {
	if e.loadedTypes[1] {
		return e.Testcases, nil
	}
	return nil, &NotLoadedError{edge: "testcases"}
}

// SubmitsOrErr returns the Submits value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) SubmitsOrErr() ([]*Submit, error) {
	if e.loadedTypes[2] {
		return e.Submits, nil
	}
	return nil, &NotLoadedError{edge: "submits"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[3] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ContestTaskOrErr returns the ContestTask value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) ContestTaskOrErr() ([]*ContestTask, error) {
	if e.loadedTypes[4] {
		return e.ContestTask, nil
	}
	return nil, &NotLoadedError{edge: "contest_task"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Task) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case task.FieldCaseInsensitive:
			values[i] = new(sql.NullBool)
		case task.FieldID, task.FieldExecTimeLimit, task.FieldExecMemoryLimit, task.FieldNdigits:
			values[i] = new(sql.NullInt64)
		case task.FieldTitle, task.FieldStatement, task.FieldDifficulty, task.FieldJudgeType, task.FieldJudgeCodePath:
			values[i] = new(sql.NullString)
		case task.FieldCreatedAt, task.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case task.ForeignKeys[0]: // user_tasks
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Task fields.
func (t *Task) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case task.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case task.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				t.Title = value.String
			}
		case task.FieldStatement:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field statement", values[i])
			} else if value.Valid {
				t.Statement = value.String
			}
		case task.FieldDifficulty:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field difficulty", values[i])
			} else if value.Valid {
				t.Difficulty = value.String
			}
		case task.FieldExecTimeLimit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field exec_time_limit", values[i])
			} else if value.Valid {
				t.ExecTimeLimit = uint(value.Int64)
			}
		case task.FieldExecMemoryLimit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field exec_memory_limit", values[i])
			} else if value.Valid {
				t.ExecMemoryLimit = uint(value.Int64)
			}
		case task.FieldJudgeType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field judge_type", values[i])
			} else if value.Valid {
				t.JudgeType = task.JudgeType(value.String)
			}
		case task.FieldCaseInsensitive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field case_insensitive", values[i])
			} else if value.Valid {
				t.CaseInsensitive = new(bool)
				*t.CaseInsensitive = value.Bool
			}
		case task.FieldNdigits:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ndigits", values[i])
			} else if value.Valid {
				t.Ndigits = new(uint)
				*t.Ndigits = uint(value.Int64)
			}
		case task.FieldJudgeCodePath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field judge_code_path", values[i])
			} else if value.Valid {
				t.JudgeCodePath = new(string)
				*t.JudgeCodePath = value.String
			}
		case task.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case task.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = new(time.Time)
				*t.UpdatedAt = value.Time
			}
		case task.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_tasks", value)
			} else if value.Valid {
				t.user_tasks = new(int)
				*t.user_tasks = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Task.
// This includes values selected through modifiers, order, etc.
func (t *Task) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryTestcaseSets queries the "testcase_sets" edge of the Task entity.
func (t *Task) QueryTestcaseSets() *TestcaseSetQuery {
	return NewTaskClient(t.config).QueryTestcaseSets(t)
}

// QueryTestcases queries the "testcases" edge of the Task entity.
func (t *Task) QueryTestcases() *TestcaseQuery {
	return NewTaskClient(t.config).QueryTestcases(t)
}

// QuerySubmits queries the "submits" edge of the Task entity.
func (t *Task) QuerySubmits() *SubmitQuery {
	return NewTaskClient(t.config).QuerySubmits(t)
}

// QueryUser queries the "user" edge of the Task entity.
func (t *Task) QueryUser() *UserQuery {
	return NewTaskClient(t.config).QueryUser(t)
}

// QueryContestTask queries the "contest_task" edge of the Task entity.
func (t *Task) QueryContestTask() *ContestTaskQuery {
	return NewTaskClient(t.config).QueryContestTask(t)
}

// Update returns a builder for updating this Task.
// Note that you need to call Task.Unwrap() before calling this method if this Task
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Task) Update() *TaskUpdateOne {
	return NewTaskClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Task entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Task) Unwrap() *Task {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Task is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Task) String() string {
	var builder strings.Builder
	builder.WriteString("Task(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("title=")
	builder.WriteString(t.Title)
	builder.WriteString(", ")
	builder.WriteString("statement=")
	builder.WriteString(t.Statement)
	builder.WriteString(", ")
	builder.WriteString("difficulty=")
	builder.WriteString(t.Difficulty)
	builder.WriteString(", ")
	builder.WriteString("exec_time_limit=")
	builder.WriteString(fmt.Sprintf("%v", t.ExecTimeLimit))
	builder.WriteString(", ")
	builder.WriteString("exec_memory_limit=")
	builder.WriteString(fmt.Sprintf("%v", t.ExecMemoryLimit))
	builder.WriteString(", ")
	builder.WriteString("judge_type=")
	builder.WriteString(fmt.Sprintf("%v", t.JudgeType))
	builder.WriteString(", ")
	if v := t.CaseInsensitive; v != nil {
		builder.WriteString("case_insensitive=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := t.Ndigits; v != nil {
		builder.WriteString("ndigits=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := t.JudgeCodePath; v != nil {
		builder.WriteString("judge_code_path=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := t.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Tasks is a parsable slice of Task.
type Tasks []*Task
