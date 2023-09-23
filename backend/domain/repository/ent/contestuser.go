// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contest"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contestuser"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
)

// ContestUser is the model entity for the ContestUser schema.
type ContestUser struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Role holds the value of the "role" field.
	Role string `json:"role,omitempty"`
	// ContestID holds the value of the "contest_id" field.
	ContestID int `json:"contest_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContestUserQuery when eager-loading is set.
	Edges        ContestUserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ContestUserEdges holds the relations/edges for other nodes in the graph.
type ContestUserEdges struct {
	// Contest holds the value of the contest edge.
	Contest *Contest `json:"contest,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ContestOrErr returns the Contest value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContestUserEdges) ContestOrErr() (*Contest, error) {
	if e.loadedTypes[0] {
		if e.Contest == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: contest.Label}
		}
		return e.Contest, nil
	}
	return nil, &NotLoadedError{edge: "contest"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContestUserEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ContestUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case contestuser.FieldID, contestuser.FieldContestID, contestuser.FieldUserID:
			values[i] = new(sql.NullInt64)
		case contestuser.FieldRole:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ContestUser fields.
func (cu *ContestUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contestuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cu.ID = int(value.Int64)
		case contestuser.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				cu.Role = value.String
			}
		case contestuser.FieldContestID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field contest_id", values[i])
			} else if value.Valid {
				cu.ContestID = int(value.Int64)
			}
		case contestuser.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				cu.UserID = int(value.Int64)
			}
		default:
			cu.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ContestUser.
// This includes values selected through modifiers, order, etc.
func (cu *ContestUser) Value(name string) (ent.Value, error) {
	return cu.selectValues.Get(name)
}

// QueryContest queries the "contest" edge of the ContestUser entity.
func (cu *ContestUser) QueryContest() *ContestQuery {
	return NewContestUserClient(cu.config).QueryContest(cu)
}

// QueryUser queries the "user" edge of the ContestUser entity.
func (cu *ContestUser) QueryUser() *UserQuery {
	return NewContestUserClient(cu.config).QueryUser(cu)
}

// Update returns a builder for updating this ContestUser.
// Note that you need to call ContestUser.Unwrap() before calling this method if this ContestUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (cu *ContestUser) Update() *ContestUserUpdateOne {
	return NewContestUserClient(cu.config).UpdateOne(cu)
}

// Unwrap unwraps the ContestUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cu *ContestUser) Unwrap() *ContestUser {
	_tx, ok := cu.config.driver.(*txDriver)
	if !ok {
		panic("ent: ContestUser is not a transactional entity")
	}
	cu.config.driver = _tx.drv
	return cu
}

// String implements the fmt.Stringer.
func (cu *ContestUser) String() string {
	var builder strings.Builder
	builder.WriteString("ContestUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cu.ID))
	builder.WriteString("role=")
	builder.WriteString(cu.Role)
	builder.WriteString(", ")
	builder.WriteString("contest_id=")
	builder.WriteString(fmt.Sprintf("%v", cu.ContestID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", cu.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// ContestUsers is a parsable slice of ContestUser.
type ContestUsers []*ContestUser
