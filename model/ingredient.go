package model

import (
	"fmt"

	"github.com/shopspring/decimal"
	null "gopkg.in/guregu/null.v3"
)

// Ingredient is the model for the ingredient table
type Ingredient struct {
	ID       int             `db:"id" json:"id"`
	RecipeID int             `db:"recipe_id" json:"recipe_id"`
	Name     string          `db:"name"`
	Measure  null.String     `db:"measure" json:"measure"`
	Quantity decimal.Decimal `db:"quantity" json:"quantity"`
}

// String is the string representation of an ingredient.Ingredient
func (i Ingredient) String() string {
	return fmt.Sprintf("%v %v %v", i.Quantity, i.Measure.String, i.Name)
}
