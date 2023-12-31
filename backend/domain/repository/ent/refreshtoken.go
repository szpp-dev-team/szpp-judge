// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/refreshtoken"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
)

// RefreshToken is the model entity for the RefreshToken schema.
type RefreshToken struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// ExpiresAt holds the value of the "expires_at" field.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// IsDead holds the value of the "is_dead" field.
	IsDead bool `json:"is_dead,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RefreshTokenQuery when eager-loading is set.
	Edges               RefreshTokenEdges `json:"edges"`
	user_refresh_tokens *int
	selectValues        sql.SelectValues
}

// RefreshTokenEdges holds the relations/edges for other nodes in the graph.
type RefreshTokenEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RefreshTokenEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RefreshToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case refreshtoken.FieldIsDead:
			values[i] = new(sql.NullBool)
		case refreshtoken.FieldID:
			values[i] = new(sql.NullInt64)
		case refreshtoken.FieldToken:
			values[i] = new(sql.NullString)
		case refreshtoken.FieldExpiresAt:
			values[i] = new(sql.NullTime)
		case refreshtoken.ForeignKeys[0]: // user_refresh_tokens
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RefreshToken fields.
func (rt *RefreshToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case refreshtoken.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rt.ID = int(value.Int64)
		case refreshtoken.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				rt.Token = value.String
			}
		case refreshtoken.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				rt.ExpiresAt = value.Time
			}
		case refreshtoken.FieldIsDead:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_dead", values[i])
			} else if value.Valid {
				rt.IsDead = value.Bool
			}
		case refreshtoken.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_refresh_tokens", value)
			} else if value.Valid {
				rt.user_refresh_tokens = new(int)
				*rt.user_refresh_tokens = int(value.Int64)
			}
		default:
			rt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RefreshToken.
// This includes values selected through modifiers, order, etc.
func (rt *RefreshToken) Value(name string) (ent.Value, error) {
	return rt.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the RefreshToken entity.
func (rt *RefreshToken) QueryUser() *UserQuery {
	return NewRefreshTokenClient(rt.config).QueryUser(rt)
}

// Update returns a builder for updating this RefreshToken.
// Note that you need to call RefreshToken.Unwrap() before calling this method if this RefreshToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (rt *RefreshToken) Update() *RefreshTokenUpdateOne {
	return NewRefreshTokenClient(rt.config).UpdateOne(rt)
}

// Unwrap unwraps the RefreshToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rt *RefreshToken) Unwrap() *RefreshToken {
	_tx, ok := rt.config.driver.(*txDriver)
	if !ok {
		panic("ent: RefreshToken is not a transactional entity")
	}
	rt.config.driver = _tx.drv
	return rt
}

// String implements the fmt.Stringer.
func (rt *RefreshToken) String() string {
	var builder strings.Builder
	builder.WriteString("RefreshToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rt.ID))
	builder.WriteString("token=")
	builder.WriteString(rt.Token)
	builder.WriteString(", ")
	builder.WriteString("expires_at=")
	builder.WriteString(rt.ExpiresAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("is_dead=")
	builder.WriteString(fmt.Sprintf("%v", rt.IsDead))
	builder.WriteByte(')')
	return builder.String()
}

// RefreshTokens is a parsable slice of RefreshToken.
type RefreshTokens []*RefreshToken
