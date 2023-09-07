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

	mysqlConfig := &mysql.Config{
		DBName:    "szpp-judge-test-db",
		User:      "root",
		Passwd:    "root",
		Addr:      "0.0.0.0:" + os.Getenv("TEST_MYSQL_PORT"),
		Net:       "tcp",
		ParseTime: true,
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
