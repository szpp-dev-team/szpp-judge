// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "statement", Type: field.TypeString},
		{Name: "difficulty", Type: field.TypeString},
		{Name: "exec_time_limit", Type: field.TypeUint},
		{Name: "exec_memory_limit", Type: field.TypeUint},
		{Name: "judge_type", Type: field.TypeEnum, Enums: []string{"normal", "eps", "interactive", "custom"}},
		{Name: "case_insensitive", Type: field.TypeBool, Nullable: true},
		{Name: "ndigits", Type: field.TypeUint, Nullable: true},
		{Name: "judge_code_path", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_tasks", Type: field.TypeInt},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tasks_users_tasks",
				Columns:    []*schema.Column{TasksColumns[12]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TestcasesColumns holds the columns for the "testcases" table.
	TestcasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "task_testcases", Type: field.TypeInt},
	}
	// TestcasesTable holds the schema information for the "testcases" table.
	TestcasesTable = &schema.Table{
		Name:       "testcases",
		Columns:    TestcasesColumns,
		PrimaryKey: []*schema.Column{TestcasesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "testcases_tasks_testcases",
				Columns:    []*schema.Column{TestcasesColumns[5]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "testcase_name_task_testcases",
				Unique:  true,
				Columns: []*schema.Column{TestcasesColumns[1], TestcasesColumns[5]},
			},
		},
	}
	// TestcaseSetsColumns holds the columns for the "testcase_sets" table.
	TestcaseSetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "score", Type: field.TypeInt},
		{Name: "is_sample", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "task_testcase_sets", Type: field.TypeInt},
	}
	// TestcaseSetsTable holds the schema information for the "testcase_sets" table.
	TestcaseSetsTable = &schema.Table{
		Name:       "testcase_sets",
		Columns:    TestcaseSetsColumns,
		PrimaryKey: []*schema.Column{TestcaseSetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "testcase_sets_tasks_testcase_sets",
				Columns:    []*schema.Column{TestcaseSetsColumns[6]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "testcaseset_name_task_testcase_sets",
				Unique:  true,
				Columns: []*schema.Column{TestcaseSetsColumns[1], TestcaseSetsColumns[6]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "role", Type: field.TypeString},
		{Name: "hashed_password", Type: field.TypeBytes},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// TestcaseSetTestcasesColumns holds the columns for the "testcase_set_testcases" table.
	TestcaseSetTestcasesColumns = []*schema.Column{
		{Name: "testcase_set_id", Type: field.TypeInt},
		{Name: "testcase_id", Type: field.TypeInt},
	}
	// TestcaseSetTestcasesTable holds the schema information for the "testcase_set_testcases" table.
	TestcaseSetTestcasesTable = &schema.Table{
		Name:       "testcase_set_testcases",
		Columns:    TestcaseSetTestcasesColumns,
		PrimaryKey: []*schema.Column{TestcaseSetTestcasesColumns[0], TestcaseSetTestcasesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "testcase_set_testcases_testcase_set_id",
				Columns:    []*schema.Column{TestcaseSetTestcasesColumns[0]},
				RefColumns: []*schema.Column{TestcaseSetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "testcase_set_testcases_testcase_id",
				Columns:    []*schema.Column{TestcaseSetTestcasesColumns[1]},
				RefColumns: []*schema.Column{TestcasesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TasksTable,
		TestcasesTable,
		TestcaseSetsTable,
		UsersTable,
		TestcaseSetTestcasesTable,
	}
)

func init() {
	TasksTable.ForeignKeys[0].RefTable = UsersTable
	TestcasesTable.ForeignKeys[0].RefTable = TasksTable
	TestcaseSetsTable.ForeignKeys[0].RefTable = TasksTable
	TestcaseSetTestcasesTable.ForeignKeys[0].RefTable = TestcaseSetsTable
	TestcaseSetTestcasesTable.ForeignKeys[1].RefTable = TestcasesTable
}
