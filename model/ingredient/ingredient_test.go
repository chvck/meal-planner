package ingredient_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	null "gopkg.in/guregu/null.v3"

	"github.com/chvck/meal-planner/model/ingredient"
	"github.com/chvck/meal-planner/testhelper"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/sqlite3"
)

func setup(t *testing.T) (*sql.DB, func()) {
	openDb, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		t.Fatal(err)
	}

	driver, err := sqlite3.WithInstance(openDb, &sqlite3.Config{})

	if err != nil {
		t.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://../../migrations/", "sqlite3", driver)

	if err != nil {
		t.Fatal(err)
	}

	if err := m.Up(); err != nil {
		t.Fatal(err)
	}

	testhelper.HelperCreateUsers(t, openDb, "../testdata/users.json")

	down := func() {
		openDb.Close()
		m.Down()
	}

	return openDb, down
}

func TestString(t *testing.T) {
	i := ingredient.Ingredient{
		ID:       1,
		Measure:  null.String{sql.NullString{String: "tbsp"}},
		Name:     "Paprika",
		Quantity: 2,
		RecipeID: 2,
	}

	assert.Equal(t, "2 tbsp Paprika", fmt.Sprint(i))
}
