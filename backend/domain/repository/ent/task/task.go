// Code generated by ent, DO NOT EDIT.

package task

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldStatement holds the string denoting the statement field in the database.
	FieldStatement = "statement"
	// FieldDifficulty holds the string denoting the difficulty field in the database.
	FieldDifficulty = "difficulty"
	// FieldExecTimeLimit holds the string denoting the exec_time_limit field in the database.
	FieldExecTimeLimit = "exec_time_limit"
	// FieldExecMemoryLimit holds the string denoting the exec_memory_limit field in the database.
	FieldExecMemoryLimit = "exec_memory_limit"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTestcaseSets holds the string denoting the testcase_sets edge name in mutations.
	EdgeTestcaseSets = "testcase_sets"
	// EdgeTestcases holds the string denoting the testcases edge name in mutations.
	EdgeTestcases = "testcases"
	// EdgeSubmits holds the string denoting the submits edge name in mutations.
	EdgeSubmits = "submits"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeContests holds the string denoting the contests edge name in mutations.
	EdgeContests = "contests"
	// EdgeContestTask holds the string denoting the contest_task edge name in mutations.
	EdgeContestTask = "contest_task"
	// Table holds the table name of the task in the database.
	Table = "tasks"
	// TestcaseSetsTable is the table that holds the testcase_sets relation/edge.
	TestcaseSetsTable = "testcase_sets"
	// TestcaseSetsInverseTable is the table name for the TestcaseSet entity.
	// It exists in this package in order to avoid circular dependency with the "testcaseset" package.
	TestcaseSetsInverseTable = "testcase_sets"
	// TestcaseSetsColumn is the table column denoting the testcase_sets relation/edge.
	TestcaseSetsColumn = "task_testcase_sets"
	// TestcasesTable is the table that holds the testcases relation/edge.
	TestcasesTable = "testcases"
	// TestcasesInverseTable is the table name for the Testcase entity.
	// It exists in this package in order to avoid circular dependency with the "testcase" package.
	TestcasesInverseTable = "testcases"
	// TestcasesColumn is the table column denoting the testcases relation/edge.
	TestcasesColumn = "task_testcases"
	// SubmitsTable is the table that holds the submits relation/edge.
	SubmitsTable = "submits"
	// SubmitsInverseTable is the table name for the Submit entity.
	// It exists in this package in order to avoid circular dependency with the "submit" package.
	SubmitsInverseTable = "submits"
	// SubmitsColumn is the table column denoting the submits relation/edge.
	SubmitsColumn = "task_submits"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "tasks"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_tasks"
	// ContestsTable is the table that holds the contests relation/edge. The primary key declared below.
	ContestsTable = "contest_tasks"
	// ContestsInverseTable is the table name for the Contest entity.
	// It exists in this package in order to avoid circular dependency with the "contest" package.
	ContestsInverseTable = "contests"
	// ContestTaskTable is the table that holds the contest_task relation/edge.
	ContestTaskTable = "contest_tasks"
	// ContestTaskInverseTable is the table name for the ContestTask entity.
	// It exists in this package in order to avoid circular dependency with the "contesttask" package.
	ContestTaskInverseTable = "contest_tasks"
	// ContestTaskColumn is the table column denoting the contest_task relation/edge.
	ContestTaskColumn = "task_id"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldStatement,
	FieldDifficulty,
	FieldExecTimeLimit,
	FieldExecMemoryLimit,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tasks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_tasks",
}

var (
	// ContestsPrimaryKey and ContestsColumn2 are the table columns denoting the
	// primary key for the contests relation (M2M).
	ContestsPrimaryKey = []string{"contest_id", "task_id"}
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

var (
	// StatementValidator is a validator for the "statement" field. It is called by the builders before save.
	StatementValidator func(string) error
)

// OrderOption defines the ordering options for the Task queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByStatement orders the results by the statement field.
func ByStatement(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatement, opts...).ToFunc()
}

// ByDifficulty orders the results by the difficulty field.
func ByDifficulty(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDifficulty, opts...).ToFunc()
}

// ByExecTimeLimit orders the results by the exec_time_limit field.
func ByExecTimeLimit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExecTimeLimit, opts...).ToFunc()
}

// ByExecMemoryLimit orders the results by the exec_memory_limit field.
func ByExecMemoryLimit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExecMemoryLimit, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByTestcaseSetsCount orders the results by testcase_sets count.
func ByTestcaseSetsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTestcaseSetsStep(), opts...)
	}
}

// ByTestcaseSets orders the results by testcase_sets terms.
func ByTestcaseSets(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTestcaseSetsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTestcasesCount orders the results by testcases count.
func ByTestcasesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTestcasesStep(), opts...)
	}
}

// ByTestcases orders the results by testcases terms.
func ByTestcases(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTestcasesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySubmitsCount orders the results by submits count.
func BySubmitsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubmitsStep(), opts...)
	}
}

// BySubmits orders the results by submits terms.
func BySubmits(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubmitsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByContestsCount orders the results by contests count.
func ByContestsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newContestsStep(), opts...)
	}
}

// ByContests orders the results by contests terms.
func ByContests(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContestsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByContestTaskCount orders the results by contest_task count.
func ByContestTaskCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newContestTaskStep(), opts...)
	}
}

// ByContestTask orders the results by contest_task terms.
func ByContestTask(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContestTaskStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTestcaseSetsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TestcaseSetsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TestcaseSetsTable, TestcaseSetsColumn),
	)
}
func newTestcasesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TestcasesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TestcasesTable, TestcasesColumn),
	)
}
func newSubmitsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubmitsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SubmitsTable, SubmitsColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newContestsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContestsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ContestsTable, ContestsPrimaryKey...),
	)
}
func newContestTaskStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContestTaskInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, ContestTaskTable, ContestTaskColumn),
	)
}
