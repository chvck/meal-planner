package ingredient

import (
	"gopkg.in/guregu/null.v3"
	"github.com/chvck/meal-planner/model"
	"fmt"
	"github.com/chvck/meal-planner/db"
)

type Ingredient struct {
	Id       int    `db:"id"`
	RecipeId int    `db:"recipe_id"`
	Name     string `db:"name"`
	Measure  null.String
	Quantity int
}

func (i Ingredient) String() string {
	return fmt.Sprintf("%v %v %v", i.Quantity, i.Measure, i.Name)
}

// All retrieves all ingredients
func All(dataStore model.IDataStoreAdapter) (*[]Ingredient, error) {
	return AllWithLimit(dataStore, "NULL", 0)
}

// AllWithLimit retrieves x ingredients starting from an offset
// limit is expected to a positive int or string NULL (for no limit)
func AllWithLimit(dataStore model.IDataStoreAdapter, limit interface{}, offset int) (*[]Ingredient, error) {
	var ingredients []Ingredient
	if rows, err := dataStore.Query(fmt.Sprintf(
		`SELECT id, recipe_id, name, measure, quantity
		FROM ingredient
		ORDER BY name
		LIMIT %v OFFSET %v;`,
		limit,
		offset,
	)); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		for rows.Next() {
			i := Ingredient{}
			rows.Scan(&i.Id, &i.RecipeId, &i.Name, &i.Measure, &i.Quantity)

			ingredients = append(ingredients, i)
		}
	}

	return &ingredients, nil
}

// Create creates a list of Ingredients
func CreateMany(tx db.Transaction, ingredients []Ingredient, recipeId int) error {
	for _, i := range ingredients {
		if _, err := tx.Exec(
			"INSERT INTO ingredient (recipe_id, name, measure, quantity) VALUES (?, ?, ?, ?);",
			recipeId, i.Name, i.Measure, i.Quantity); err != nil {
			return err
		}
	}

	return nil
}

// Delete all of the Ingredients for a Recipe
func DeleteAllByRecipe(tx db.Transaction, recipeId int) error {
	_, err := tx.Exec(
		"DELETE FROM ingredient WHERE recipe_id = ?;",
		recipeId)

	return err
}
