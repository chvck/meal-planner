// +build integration

package recipe

import (
	"testing"
	"github.com/chvck/meal-planner/db"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/sqlite3"
	"database/sql"
	_ "github.com/mattes/migrate/source/file"
	"github.com/jmoiron/sqlx"
)

func setup() (*sql.DB, func(), error) {
	openDb, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		return nil, nil, err
	}

	driver, err := sqlite3.WithInstance(openDb, &sqlite3.Config{})

	if err != nil {
		return nil, nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("file://../../migrations/", "sqlite3", driver)

	if err != nil {
		return nil, nil, err
	}

	m.Up()

	down := func() {
		m.Down()
	}

	return openDb, down, err
}

func TestIntegration(t *testing.T) {
	openDb, teardown, err := setup()

	if err != nil {
		t.Error(err.Error())
		return
	}
	defer openDb.Close()
	defer teardown()

	adapter := db.SqlxAdapter{}

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, "sqlite3")); err != nil {
		t.Error(err)
		return
	}
}
