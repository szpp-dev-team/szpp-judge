// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcaseresult"
)

// TestcaseResult is the model entity for the TestcaseResult schema.
type TestcaseResult struct {
	config
	// ID of the ent.
	ID                      int `json:"id,omitempty"`
	submit_testcase_results *int
	selectValues            sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TestcaseResult) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case testcaseresult.FieldID:
			values[i] = new(sql.NullInt64)
		case testcaseresult.ForeignKeys[0]: // submit_testcase_results
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TestcaseResult fields.
func (tr *TestcaseResult) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case testcaseresult.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tr.ID = int(value.Int64)
		case testcaseresult.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field submit_testcase_results", value)
			} else if value.Valid {
				tr.submit_testcase_results = new(int)
				*tr.submit_testcase_results = int(value.Int64)
			}
		default:
			tr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TestcaseResult.
// This includes values selected through modifiers, order, etc.
func (tr *TestcaseResult) Value(name string) (ent.Value, error) {
	return tr.selectValues.Get(name)
}

// Update returns a builder for updating this TestcaseResult.
// Note that you need to call TestcaseResult.Unwrap() before calling this method if this TestcaseResult
// was returned from a transaction, and the transaction was committed or rolled back.
func (tr *TestcaseResult) Update() *TestcaseResultUpdateOne {
	return NewTestcaseResultClient(tr.config).UpdateOne(tr)
}

// Unwrap unwraps the TestcaseResult entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tr *TestcaseResult) Unwrap() *TestcaseResult {
	_tx, ok := tr.config.driver.(*txDriver)
	if !ok {
		panic("ent: TestcaseResult is not a transactional entity")
	}
	tr.config.driver = _tx.drv
	return tr
}

// String implements the fmt.Stringer.
func (tr *TestcaseResult) String() string {
	var builder strings.Builder
	builder.WriteString("TestcaseResult(")
	builder.WriteString(fmt.Sprintf("id=%v", tr.ID))
	builder.WriteByte(')')
	return builder.String()
}

// TestcaseResults is a parsable slice of TestcaseResult.
type TestcaseResults []*TestcaseResult
