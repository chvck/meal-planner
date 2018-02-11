package ingredient

import (
	"fmt"
	"strings"

	"github.com/chvck/meal-planner/db"
	"github.com/chvck/meal-planner/model"
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

// ForRecipes returns the ingredients for a list of recipe ids. Ingredients are keyed by recipe ID
func ForRecipes(dataStore model.IDataStoreAdapter, ids ...interface{}) (map[int][]Ingredient, error) {
	m := make(map[int][]Ingredient)
	in := strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ",")

	query := fmt.Sprintf(
		`SELECT id, recipe_id, name, measure, quantity
		FROM ingredient
		WHERE recipe_id IN (%v)
		ORDER BY recipe_id;`,
		in,
	)

	if rows, err := dataStore.Query(query, ids...); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		for rows.Next() {
			var (
				rID     int
				ingID   int
				ingName string
				mName   null.String
				q       int
			)
			if err := rows.Scan(&ingID, &rID, &ingName, &mName, &q); err != nil {
				return nil, err
			}

			arr := m[rID]
			i := Ingredient{ID: ingID, RecipeID: rID, Name: ingName, Measure: mName, Quantity: q}
			arr = append(arr, i)
			m[rID] = arr
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}
	}

	return m, nil
}
