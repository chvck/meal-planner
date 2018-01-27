package ingredient

import (
	"fmt"

	"github.com/chvck/meal-planner/db"
	"gopkg.in/guregu/null.v3"
)

// Ingredient is the model for the ingredient table
type Ingredient struct {
	ID       int         `db:"id" json:"id"`
	RecipeID int         `db:"recipe_id" json:"recipe_id"`
	Name     string      `db:"name"`
	Measure  null.String `db:"measure" json:"measure"`
	Quantity int         `db:"quantity" json:"quantity"`
}

// String is the string representation of an Ingredient
func (i Ingredient) String() string {
	return fmt.Sprintf("%v %v %v", i.Quantity, i.Measure.String, i.Name)
}

// CreateMany creates a list of Ingredients
func CreateMany(tx db.Transaction, ingredients []Ingredient, recipeID int) error {
	for _, i := range ingredients {
		row := tx.QueryOne(
			"INSERT INTO ingredient (recipe_id, name, measure, quantity) VALUES (?, ?, ?, ?) RETURNING id;",
			recipeID, i.Name, i.Measure, i.Quantity)

		var ingID int
		if err := row.Scan(&ingID); err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}

// DeleteAllByRecipe all of the Ingredients for a Recipe
func DeleteAllByRecipe(tx db.Transaction, recipeID int) error {
	_, err := tx.Exec(
		"DELETE FROM ingredient WHERE recipe_id = ?;",
		recipeID)

	return err
}
