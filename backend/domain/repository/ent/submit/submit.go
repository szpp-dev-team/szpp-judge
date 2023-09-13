// Code generated by ent, DO NOT EDIT.

package submit

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the submit type in the database.
	Label = "submit"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldExecTime holds the string denoting the exec_time field in the database.
	FieldExecTime = "exec_time"
	// FieldExecMemory holds the string denoting the exec_memory field in the database.
	FieldExecMemory = "exec_memory"
	// FieldSubmittedAt holds the string denoting the submitted_at field in the database.
	FieldSubmittedAt = "submitted_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeTask holds the string denoting the task edge name in mutations.
	EdgeTask = "task"
	// EdgeLanguage holds the string denoting the language edge name in mutations.
	EdgeLanguage = "language"
	// EdgeTestcaseResults holds the string denoting the testcase_results edge name in mutations.
	EdgeTestcaseResults = "testcase_results"
	// Table holds the table name of the submit in the database.
	Table = "submits"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "user_submits"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// TaskTable is the table that holds the task relation/edge.
	TaskTable = "submits"
	// TaskInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskInverseTable = "tasks"
	// TaskColumn is the table column denoting the task relation/edge.
	TaskColumn = "task_submits"
	// LanguageTable is the table that holds the language relation/edge.
	LanguageTable = "submits"
	// LanguageInverseTable is the table name for the Language entity.
	// It exists in this package in order to avoid circular dependency with the "language" package.
	LanguageInverseTable = "languages"
	// LanguageColumn is the table column denoting the language relation/edge.
	LanguageColumn = "language_submits"
	// TestcaseResultsTable is the table that holds the testcase_results relation/edge.
	TestcaseResultsTable = "testcase_results"
	// TestcaseResultsInverseTable is the table name for the TestcaseResult entity.
	// It exists in this package in order to avoid circular dependency with the "testcaseresult" package.
	TestcaseResultsInverseTable = "testcase_results"
	// TestcaseResultsColumn is the table column denoting the testcase_results relation/edge.
	TestcaseResultsColumn = "submit_testcase_results"
)

// Columns holds all SQL columns for submit fields.
var Columns = []string{
	FieldID,
	FieldStatus,
	FieldExecTime,
	FieldExecMemory,
	FieldSubmittedAt,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "submits"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"language_submits",
	"task_submits",
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "submit_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Submit queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByExecTime orders the results by the exec_time field.
func ByExecTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExecTime, opts...).ToFunc()
}

// ByExecMemory orders the results by the exec_memory field.
func ByExecMemory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExecMemory, opts...).ToFunc()
}

// BySubmittedAt orders the results by the submitted_at field.
func BySubmittedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubmittedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUsersCount orders the results by users count.
func ByUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersStep(), opts...)
	}
}

// ByUsers orders the results by users terms.
func ByUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTaskField orders the results by task field.
func ByTaskField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTaskStep(), sql.OrderByField(field, opts...))
	}
}

// ByLanguageField orders the results by language field.
func ByLanguageField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLanguageStep(), sql.OrderByField(field, opts...))
	}
}

// ByTestcaseResultsCount orders the results by testcase_results count.
func ByTestcaseResultsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTestcaseResultsStep(), opts...)
	}
}

// ByTestcaseResults orders the results by testcase_results terms.
func ByTestcaseResults(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTestcaseResultsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
	)
}
func newTaskStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TaskInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TaskTable, TaskColumn),
	)
}
func newLanguageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LanguageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, LanguageTable, LanguageColumn),
	)
}
func newTestcaseResultsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TestcaseResultsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TestcaseResultsTable, TestcaseResultsColumn),
	)
}
