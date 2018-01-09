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
	"gopkg.in/guregu/null.v3"
	"github.com/stretchr/testify/assert"
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

func TestIntegrationOne(t *testing.T) {
	openDb, teardown, err := setup()

	if err != nil {
		t.Error(err.Error())
		return
	}
	defer openDb.Close()
	defer teardown()

	adapter := db.SqlxAdapter{}

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, "sqlite3")); err != nil {
		t.Fatal(err)
		return
	}

	ing1 := ingredient{Name: "Chicken breast", Measure: "", Quantity: 2}
	ing2 := ingredient{Name: "Paprika", Measure: "tsp", Quantity: 1}

	expected := recipe{
		Name:         "Chicken curry",
		Description:  null.StringFrom("A tasty chicken curry"),
		Instructions: "Cook it real good",
		CookTime:     null.IntFrom(10),
		PrepTime:     null.IntFrom(15),
		Yield:        null.IntFrom(2),
		Ingredients:  []ingredient{ing1, ing2},
	}

	openDb.Exec(`INSERT INTO "user" (id, "username", "email", "password", "salt", "algorithm", "iterations", "created_at", "updated_at")
    VALUES (1, 'user', '"user@email.com', 'password', 'salt', 'algo', 12, 0, 0)`)

	openDb.Exec(`INSERT INTO recipe (id, name, instructions, yield, prep_time, cook_time, description, user_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		1,
		expected.Name,
		expected.Instructions,
		expected.Yield,
		expected.PrepTime,
		expected.CookTime,
		expected.Description,
		1,
	)

	openDb.Exec(`INSERT INTO ingredient (id, name) VALUES (?, ?)`, 1, ing1.Name)
	openDb.Exec(`INSERT INTO ingredient (id, name) VALUES (?, ?)`, 2, ing2.Name)

	openDb.Exec(`INSERT INTO recipe_to_ingredient (recipe_id, ingredient_id, measure, quantity) VALUES (?, ?, ?, ?)`,
		1,
		1,
		"",
		ing1.Quantity,
	)
	openDb.Exec(`INSERT INTO recipe_to_ingredient (recipe_id, ingredient_id, measure, quantity) VALUES (?, ?, ?, ?)`,
		1,
		2,
		ing2.Measure,
		ing2.Quantity,
	)

	recipe, err := One(&adapter, 1)

	if err != nil {
		t.Fatal(err)
		return
	}

	assertRecipe(t, &expected, recipe)

}

func assertRecipe(t *testing.T, expected *recipe, actual *recipe) {
	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.CookTime, actual.CookTime)
	assert.Equal(t, expected.PrepTime, actual.PrepTime)
	assert.Equal(t, expected.Yield, actual.Yield)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Instructions, actual.Instructions)
	assert.Equal(t, expected.Ingredients, actual.Ingredients)
}
