// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contest"
)

// Contest is the model entity for the Contest schema.
type Contest struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// PenaltySeconds holds the value of the "penalty_seconds" field.
	PenaltySeconds int `json:"penalty_seconds,omitempty"`
	// ContestType holds the value of the "contest_type" field.
	ContestType string `json:"contest_type,omitempty"`
	// IsPublic holds the value of the "is_public" field.
	IsPublic bool `json:"is_public,omitempty"`
	// StartAt holds the value of the "start_at" field.
	StartAt time.Time `json:"start_at,omitempty"`
	// EndAt holds the value of the "end_at" field.
	EndAt time.Time `json:"end_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContestQuery when eager-loading is set.
	Edges        ContestEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ContestEdges holds the relations/edges for other nodes in the graph.
type ContestEdges struct {
	// Submits holds the value of the submits edge.
	Submits []*Submit `json:"submits,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Tasks holds the value of the tasks edge.
	Tasks []*Task `json:"tasks,omitempty"`
	// ContestUser holds the value of the contest_user edge.
	ContestUser []*ContestUser `json:"contest_user,omitempty"`
	// ContestTask holds the value of the contest_task edge.
	ContestTask []*ContestTask `json:"contest_task,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// SubmitsOrErr returns the Submits value or an error if the edge
// was not loaded in eager-loading.
func (e ContestEdges) SubmitsOrErr() ([]*Submit, error) {
	if e.loadedTypes[0] {
		return e.Submits, nil
	}
	return nil, &NotLoadedError{edge: "submits"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e ContestEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// TasksOrErr returns the Tasks value or an error if the edge
// was not loaded in eager-loading.
func (e ContestEdges) TasksOrErr() ([]*Task, error) {
	if e.loadedTypes[2] {
		return e.Tasks, nil
	}
	return nil, &NotLoadedError{edge: "tasks"}
}

// ContestUserOrErr returns the ContestUser value or an error if the edge
// was not loaded in eager-loading.
func (e ContestEdges) ContestUserOrErr() ([]*ContestUser, error) {
	if e.loadedTypes[3] {
		return e.ContestUser, nil
	}
	return nil, &NotLoadedError{edge: "contest_user"}
}

// ContestTaskOrErr returns the ContestTask value or an error if the edge
// was not loaded in eager-loading.
func (e ContestEdges) ContestTaskOrErr() ([]*ContestTask, error) {
	if e.loadedTypes[4] {
		return e.ContestTask, nil
	}
	return nil, &NotLoadedError{edge: "contest_task"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Contest) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case contest.FieldIsPublic:
			values[i] = new(sql.NullBool)
		case contest.FieldID, contest.FieldPenaltySeconds:
			values[i] = new(sql.NullInt64)
		case contest.FieldName, contest.FieldSlug, contest.FieldDescription, contest.FieldContestType:
			values[i] = new(sql.NullString)
		case contest.FieldStartAt, contest.FieldEndAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Contest fields.
func (c *Contest) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contest.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case contest.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case contest.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				c.Slug = value.String
			}
		case contest.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case contest.FieldPenaltySeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field penalty_seconds", values[i])
			} else if value.Valid {
				c.PenaltySeconds = int(value.Int64)
			}
		case contest.FieldContestType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contest_type", values[i])
			} else if value.Valid {
				c.ContestType = value.String
			}
		case contest.FieldIsPublic:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_public", values[i])
			} else if value.Valid {
				c.IsPublic = value.Bool
			}
		case contest.FieldStartAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_at", values[i])
			} else if value.Valid {
				c.StartAt = value.Time
			}
		case contest.FieldEndAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_at", values[i])
			} else if value.Valid {
				c.EndAt = value.Time
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Contest.
// This includes values selected through modifiers, order, etc.
func (c *Contest) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QuerySubmits queries the "submits" edge of the Contest entity.
func (c *Contest) QuerySubmits() *SubmitQuery {
	return NewContestClient(c.config).QuerySubmits(c)
}

// QueryUsers queries the "users" edge of the Contest entity.
func (c *Contest) QueryUsers() *UserQuery {
	return NewContestClient(c.config).QueryUsers(c)
}

// QueryTasks queries the "tasks" edge of the Contest entity.
func (c *Contest) QueryTasks() *TaskQuery {
	return NewContestClient(c.config).QueryTasks(c)
}

// QueryContestUser queries the "contest_user" edge of the Contest entity.
func (c *Contest) QueryContestUser() *ContestUserQuery {
	return NewContestClient(c.config).QueryContestUser(c)
}

// QueryContestTask queries the "contest_task" edge of the Contest entity.
func (c *Contest) QueryContestTask() *ContestTaskQuery {
	return NewContestClient(c.config).QueryContestTask(c)
}

// Update returns a builder for updating this Contest.
// Note that you need to call Contest.Unwrap() before calling this method if this Contest
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Contest) Update() *ContestUpdateOne {
	return NewContestClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Contest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Contest) Unwrap() *Contest {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Contest is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Contest) String() string {
	var builder strings.Builder
	builder.WriteString("Contest(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(c.Slug)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	builder.WriteString("penalty_seconds=")
	builder.WriteString(fmt.Sprintf("%v", c.PenaltySeconds))
	builder.WriteString(", ")
	builder.WriteString("contest_type=")
	builder.WriteString(c.ContestType)
	builder.WriteString(", ")
	builder.WriteString("is_public=")
	builder.WriteString(fmt.Sprintf("%v", c.IsPublic))
	builder.WriteString(", ")
	builder.WriteString("start_at=")
	builder.WriteString(c.StartAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_at=")
	builder.WriteString(c.EndAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Contests is a parsable slice of Contest.
type Contests []*Contest