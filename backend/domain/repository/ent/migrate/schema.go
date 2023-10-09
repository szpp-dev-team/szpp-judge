// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ContestsColumns holds the columns for the "contests" table.
	ContestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "penalty_seconds", Type: field.TypeInt},
		{Name: "contest_type", Type: field.TypeString},
		{Name: "is_public", Type: field.TypeBool},
		{Name: "start_at", Type: field.TypeTime},
		{Name: "end_at", Type: field.TypeTime},
	}
	// ContestsTable holds the schema information for the "contests" table.
	ContestsTable = &schema.Table{
		Name:       "contests",
		Columns:    ContestsColumns,
		PrimaryKey: []*schema.Column{ContestsColumns[0]},
	}
	// ContestTasksColumns holds the columns for the "contest_tasks" table.
	ContestTasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "score", Type: field.TypeInt},
		{Name: "order", Type: field.TypeInt},
		{Name: "contest_id", Type: field.TypeInt},
		{Name: "task_id", Type: field.TypeInt},
	}
	// ContestTasksTable holds the schema information for the "contest_tasks" table.
	ContestTasksTable = &schema.Table{
		Name:       "contest_tasks",
		Columns:    ContestTasksColumns,
		PrimaryKey: []*schema.Column{ContestTasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "contest_tasks_contests_contest",
				Columns:    []*schema.Column{ContestTasksColumns[3]},
				RefColumns: []*schema.Column{ContestsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "contest_tasks_tasks_task",
				Columns:    []*schema.Column{ContestTasksColumns[4]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "contesttask_contest_id_task_id",
				Unique:  true,
				Columns: []*schema.Column{ContestTasksColumns[3], ContestTasksColumns[4]},
			},
		},
	}
	// ContestUsersColumns holds the columns for the "contest_users" table.
	ContestUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role", Type: field.TypeString},
		{Name: "contest_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// ContestUsersTable holds the schema information for the "contest_users" table.
	ContestUsersTable = &schema.Table{
		Name:       "contest_users",
		Columns:    ContestUsersColumns,
		PrimaryKey: []*schema.Column{ContestUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "contest_users_contests_contest",
				Columns:    []*schema.Column{ContestUsersColumns[2]},
				RefColumns: []*schema.Column{ContestsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "contest_users_users_user",
				Columns:    []*schema.Column{ContestUsersColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "contestuser_contest_id_user_id",
				Unique:  true,
				Columns: []*schema.Column{ContestUsersColumns[2], ContestUsersColumns[3]},
			},
		},
	}
	// LanguagesColumns holds the columns for the "languages" table.
	LanguagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString, Unique: true},
	}
	// LanguagesTable holds the schema information for the "languages" table.
	LanguagesTable = &schema.Table{
		Name:       "languages",
		Columns:    LanguagesColumns,
		PrimaryKey: []*schema.Column{LanguagesColumns[0]},
	}
	// RefreshTokensColumns holds the columns for the "refresh_tokens" table.
	RefreshTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token", Type: field.TypeString, Unique: true},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "is_dead", Type: field.TypeBool},
		{Name: "user_refresh_tokens", Type: field.TypeInt, Nullable: true},
	}
	// RefreshTokensTable holds the schema information for the "refresh_tokens" table.
	RefreshTokensTable = &schema.Table{
		Name:       "refresh_tokens",
		Columns:    RefreshTokensColumns,
		PrimaryKey: []*schema.Column{RefreshTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "refresh_tokens_users_refresh_tokens",
				Columns:    []*schema.Column{RefreshTokensColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SubmitsColumns holds the columns for the "submits" table.
	SubmitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeString, Nullable: true},
		{Name: "exec_time", Type: field.TypeInt, Nullable: true},
		{Name: "exec_memory", Type: field.TypeInt, Nullable: true},
		{Name: "score", Type: field.TypeInt, Nullable: true},
		{Name: "submitted_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "contest_submits", Type: field.TypeInt, Nullable: true},
		{Name: "language_submits", Type: field.TypeInt},
		{Name: "task_submits", Type: field.TypeInt},
		{Name: "user_submits", Type: field.TypeInt},
	}
	// SubmitsTable holds the schema information for the "submits" table.
	SubmitsTable = &schema.Table{
		Name:       "submits",
		Columns:    SubmitsColumns,
		PrimaryKey: []*schema.Column{SubmitsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "submits_contests_submits",
				Columns:    []*schema.Column{SubmitsColumns[8]},
				RefColumns: []*schema.Column{ContestsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "submits_languages_submits",
				Columns:    []*schema.Column{SubmitsColumns[9]},
				RefColumns: []*schema.Column{LanguagesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "submits_tasks_submits",
				Columns:    []*schema.Column{SubmitsColumns[10]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "submits_users_submits",
				Columns:    []*schema.Column{SubmitsColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "statement", Type: field.TypeString},
		{Name: "difficulty", Type: field.TypeString},
		{Name: "exec_time_limit", Type: field.TypeUint},
		{Name: "exec_memory_limit", Type: field.TypeUint},
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
				Columns:    []*schema.Column{TasksColumns[8]},
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
	// TestcaseResultsColumns holds the columns for the "testcase_results" table.
	TestcaseResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeString},
		{Name: "exec_time", Type: field.TypeInt},
		{Name: "exec_memory", Type: field.TypeInt},
		{Name: "submit_testcase_results", Type: field.TypeInt, Nullable: true},
		{Name: "testcase_result_testcase", Type: field.TypeInt, Nullable: true},
	}
	// TestcaseResultsTable holds the schema information for the "testcase_results" table.
	TestcaseResultsTable = &schema.Table{
		Name:       "testcase_results",
		Columns:    TestcaseResultsColumns,
		PrimaryKey: []*schema.Column{TestcaseResultsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "testcase_results_submits_testcase_results",
				Columns:    []*schema.Column{TestcaseResultsColumns[4]},
				RefColumns: []*schema.Column{SubmitsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "testcase_results_testcases_testcase",
				Columns:    []*schema.Column{TestcaseResultsColumns[5]},
				RefColumns: []*schema.Column{TestcasesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TestcaseSetsColumns holds the columns for the "testcase_sets" table.
	TestcaseSetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "score_ratio", Type: field.TypeInt},
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
		ContestsTable,
		ContestTasksTable,
		ContestUsersTable,
		LanguagesTable,
		RefreshTokensTable,
		SubmitsTable,
		TasksTable,
		TestcasesTable,
		TestcaseResultsTable,
		TestcaseSetsTable,
		UsersTable,
		TestcaseSetTestcasesTable,
	}
)

func init() {
	ContestTasksTable.ForeignKeys[0].RefTable = ContestsTable
	ContestTasksTable.ForeignKeys[1].RefTable = TasksTable
	ContestUsersTable.ForeignKeys[0].RefTable = ContestsTable
	ContestUsersTable.ForeignKeys[1].RefTable = UsersTable
	RefreshTokensTable.ForeignKeys[0].RefTable = UsersTable
	SubmitsTable.ForeignKeys[0].RefTable = ContestsTable
	SubmitsTable.ForeignKeys[1].RefTable = LanguagesTable
	SubmitsTable.ForeignKeys[2].RefTable = TasksTable
	SubmitsTable.ForeignKeys[3].RefTable = UsersTable
	TasksTable.ForeignKeys[0].RefTable = UsersTable
	TestcasesTable.ForeignKeys[0].RefTable = TasksTable
	TestcaseResultsTable.ForeignKeys[0].RefTable = SubmitsTable
	TestcaseResultsTable.ForeignKeys[1].RefTable = TestcasesTable
	TestcaseSetsTable.ForeignKeys[0].RefTable = TasksTable
	TestcaseSetTestcasesTable.ForeignKeys[0].RefTable = TestcaseSetsTable
	TestcaseSetTestcasesTable.ForeignKeys[1].RefTable = TestcasesTable
}
