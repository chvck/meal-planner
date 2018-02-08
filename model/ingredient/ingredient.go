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

// CreateMany creates many ingredients for a given recipe id using a transaction
func CreateMany(tx db.Transaction, ingredients []Ingredient, recipeID int) error {
	query := "INSERT INTO ingredient (name, measure, quantity, recipe_id) VALUES (?, ?, ?, ?);"
	for _, ing := range ingredients {
		if _, err := tx.Exec(query, ing.Name, ing.Measure, ing.Quantity, recipeID); err != nil {
			return err
		}
	}

	return nil
}

// DeleteAllByRecipe deletes all ingredients for a given recipe id using a transaction
func DeleteAllByRecipe(tx db.Transaction, recipeID int) error {
	query := "DELETE FROM ingredient WHERE recipe_id = ?;"
	if _, err := tx.Exec(query, recipeID); err != nil {
		return err
	}

	return nil
}
