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

	ing1 := ingredientWithProps{Id: 1, Name: "Chicken breast", Measure: null.String{}, Quantity: 2}
	ing2 := ingredientWithProps{Id: 2, Name: "Paprika", Measure: null.StringFrom("tsp"), Quantity: 1}

	expected := recipe{
		Id:           1,
		Name:         "Chicken curry",
		Description:  null.StringFrom("A tasty chicken curry"),
		Instructions: "Cook it real good",
		CookTime:     null.IntFrom(10),
		PrepTime:     null.IntFrom(15),
		Yield:        null.IntFrom(2),
		Ingredients:  []ingredientWithProps{ing1, ing2},
	}

	openDb.Exec(`INSERT INTO "user" (id, "username", "email", "password", "salt", "algorithm", "iterations", "created_at", "updated_at")
    VALUES (1, 'user', '"user@email.com', 'password', 'salt', 'algo', 12, 0, 0)`)

	openDb.Exec(`INSERT INTO recipe (id, name, instructions, yield, prep_time, cook_time, description, user_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		expected.Id,
		expected.Name,
		expected.Instructions,
		expected.Yield,
		expected.PrepTime,
		expected.CookTime,
		expected.Description,
		1,
	)

	openDb.Exec(`INSERT INTO ingredientWithProps (id, name) VALUES (?, ?)`, ing1.Id, ing1.Name)
	openDb.Exec(`INSERT INTO ingredientWithProps (id, name) VALUES (?, ?)`, ing2.Id, ing2.Name)

	openDb.Exec(`INSERT INTO recipe_to_ingredient (recipe_id, ingredient_id, measure, quantity) VALUES (?, ?, ?, ?)`,
		expected.Id,
		ing1.Id,
		ing1.Measure,
		ing1.Quantity,
	)
	openDb.Exec(`INSERT INTO recipe_to_ingredient (recipe_id, ingredient_id, measure, quantity) VALUES (?, ?, ?, ?)`,
		expected.Id,
		ing2.Id,
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

func TestIntegrationAll(t *testing.T) {
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

	ing1 := ingredientWithProps{Id: 1, Name: "Chicken breast", Measure: null.String{}, Quantity: 2}
	ing2 := ingredientWithProps{Id: 2, Name: "Paprika", Measure: null.StringFrom("tsp"), Quantity: 1}

	r1 := recipe{
		Id:           1,
		Name:         "Chicken curry",
		Description:  null.StringFrom("A tasty chicken curry"),
		Instructions: "Cook it real good",
		CookTime:     null.IntFrom(10),
		PrepTime:     null.IntFrom(15),
		Yield:        null.IntFrom(2),
		Ingredients:  []ingredientWithProps{ing1, ing2},
	}

	r2 := recipe{
		Id:           2,
		Name:         "Beef curry",
		Description:  null.StringFrom("A tasty beef curry"),
		Instructions: "Cook it slow",
		CookTime:     null.IntFrom(60),
		PrepTime:     null.IntFrom(5),
		Yield:        null.IntFrom(4),
		Ingredients:  []ingredientWithProps{ing2},
	}

	openDb.Exec(`INSERT INTO "user" (id, "username", "email", "password", "salt", "algorithm", "iterations", "created_at", "updated_at")
    VALUES (1, 'user', '"user@email.com', 'password', 'salt', 'algo', 12, 0, 0)`)

	openDb.Exec(`INSERT INTO recipe (id, name, instructions, yield, prep_time, cook_time, description, user_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		r1.Id,
		r1.Name,
		r1.Instructions,
		r1.Yield,
		r1.PrepTime,
		r1.CookTime,
		r1.Description,
		1,
	)

	openDb.Exec(`INSERT INTO recipe (id, name, instructions, yield, prep_time, cook_time, description, user_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		r2.Id,
		r2.Name,
		r2.Instructions,
		r2.Yield,
		r2.PrepTime,
		r2.CookTime,
		r2.Description,
		1,
	)

	openDb.Exec(`INSERT INTO ingredientWithProps (id, name) VALUES (?, ?)`, ing1.Id, ing1.Name)
	openDb.Exec(`INSERT INTO ingredientWithProps (id, name) VALUES (?, ?)`, ing2.Id, ing2.Name)

	openDb.Exec(`INSERT INTO recipe_to_ingredient (recipe_id, ingredient_id, measure, quantity) VALUES (?, ?, ?, ?)`,
		r1.Id,
		ing1.Id,
		ing1.Measure,
		ing1.Quantity,
	)
	openDb.Exec(`INSERT INTO recipe_to_ingredient (recipe_id, ingredient_id, measure, quantity) VALUES (?, ?, ?, ?)`,
		r1.Id,
		ing2.Id,
		ing2.Measure,
		ing2.Quantity,
	)
	openDb.Exec(`INSERT INTO recipe_to_ingredient (recipe_id, ingredient_id, measure, quantity) VALUES (?, ?, ?, ?)`,
		r2.Id,
		ing2.Id,
		ing2.Measure,
		ing2.Quantity,
	)

	recipe, err := One(&adapter, 1)

	if err != nil {
		t.Fatal(err)
		return
	}

	assertRecipe(t, &r1, recipe)
	assertRecipe(t, &r2, recipe)
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
