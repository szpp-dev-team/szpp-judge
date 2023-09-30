package utils

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/migrate"
)

func NewTestClient(t *testing.T) *ent.Client {
	t.Helper()

	println("!!!!!!!!!!!!!!!!!!!!!!!!!")
	println(os.Getenv("MYSQL_DATABASE"))

	mysqlConfig := &mysql.Config{
		DBName:               os.Getenv("MYSQL_DATABASE"),
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_ROOT_PASSWORD"),
		Addr:                 os.Getenv("addr") + os.Getenv("TEST_MYSQL_PORT"),
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", mysqlConfig.FormatDSN())
	require.NoError(t, err)
	client := ent.NewClient(ent.Driver(entsql.OpenDB("mysql", db)))
	for i := time.Duration(2); ; i = i * 2 {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err := client.Schema.Create(ctx, migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err != nil {
			// sometimes "mysql: querying mysql version driver: bad connection" error occurs, so retry it
			t.Log(err)
			time.Sleep(i * time.Second)
		} else {
			break
		}
	}
	return client
}

func TruncateDB(t *testing.T, client *ent.Client) {
	t.Helper()
	ctx := context.Background()

	rows, err := client.QueryContext(ctx, "SHOW TABLES")
	require.NoError(t, err)
	defer rows.Close()
	for rows.Next() {
		var table string
		require.NoError(t, rows.Scan(&table))
		_, err := client.ExecContext(ctx, "SET foreign_key_checks = 0")
		require.NoError(t, err)
		defer client.ExecContext(ctx, "SET foreign_key_checks = 1")
		_, err = client.ExecContext(ctx, "TRUNCATE TABLE "+table)
		require.NoError(t, err)
	}
}
