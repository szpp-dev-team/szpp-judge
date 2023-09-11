// Code generated by ent, DO NOT EDIT.

package task

import (
	"fmt"

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
	// FieldJudgeType holds the string denoting the judge_type field in the database.
	FieldJudgeType = "judge_type"
	// FieldCaseInsensitive holds the string denoting the case_insensitive field in the database.
	FieldCaseInsensitive = "case_insensitive"
	// FieldNdigits holds the string denoting the ndigits field in the database.
	FieldNdigits = "ndigits"
	// FieldJudgeCodePath holds the string denoting the judge_code_path field in the database.
	FieldJudgeCodePath = "judge_code_path"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTestcaseSets holds the string denoting the testcase_sets edge name in mutations.
	EdgeTestcaseSets = "testcase_sets"
	// EdgeTestcases holds the string denoting the testcases edge name in mutations.
	EdgeTestcases = "testcases"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
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
	// UserTable is the table that holds the user relation/edge.
	UserTable = "tasks"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_tasks"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldStatement,
	FieldDifficulty,
	FieldExecTimeLimit,
	FieldExecMemoryLimit,
	FieldJudgeType,
	FieldCaseInsensitive,
	FieldNdigits,
	FieldJudgeCodePath,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tasks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_tasks",
}

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

// JudgeType defines the type for the "judge_type" enum field.
type JudgeType string

// JudgeType values.
const (
	JudgeTypeNormal      JudgeType = "normal"
	JudgeTypeEps         JudgeType = "eps"
	JudgeTypeInteractive JudgeType = "interactive"
	JudgeTypeCustom      JudgeType = "custom"
)

func (jt JudgeType) String() string {
	return string(jt)
}

// JudgeTypeValidator is a validator for the "judge_type" field enum values. It is called by the builders before save.
func JudgeTypeValidator(jt JudgeType) error {
	switch jt {
	case JudgeTypeNormal, JudgeTypeEps, JudgeTypeInteractive, JudgeTypeCustom:
		return nil
	default:
		return fmt.Errorf("task: invalid enum value for judge_type field: %q", jt)
	}
}

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

// ByJudgeType orders the results by the judge_type field.
func ByJudgeType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJudgeType, opts...).ToFunc()
}

// ByCaseInsensitive orders the results by the case_insensitive field.
func ByCaseInsensitive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCaseInsensitive, opts...).ToFunc()
}

// ByNdigits orders the results by the ndigits field.
func ByNdigits(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNdigits, opts...).ToFunc()
}

// ByJudgeCodePath orders the results by the judge_code_path field.
func ByJudgeCodePath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJudgeCodePath, opts...).ToFunc()
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

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
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
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}