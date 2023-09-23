// Code generated by ent, DO NOT EDIT.

package contestuser

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the contestuser type in the database.
	Label = "contest_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldContestID holds the string denoting the contest_id field in the database.
	FieldContestID = "contest_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeContest holds the string denoting the contest edge name in mutations.
	EdgeContest = "contest"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the contestuser in the database.
	Table = "contest_users"
	// ContestTable is the table that holds the contest relation/edge.
	ContestTable = "contest_users"
	// ContestInverseTable is the table name for the Contest entity.
	// It exists in this package in order to avoid circular dependency with the "contest" package.
	ContestInverseTable = "contests"
	// ContestColumn is the table column denoting the contest relation/edge.
	ContestColumn = "contest_id"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "contest_users"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for contestuser fields.
var Columns = []string{
	FieldID,
	FieldRole,
	FieldContestID,
	FieldUserID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the ContestUser queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByContestID orders the results by the contest_id field.
func ByContestID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContestID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByContestField orders the results by contest field.
func ByContestField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContestStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newContestStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContestInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ContestTable, ContestColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
	)
}
