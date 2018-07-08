package model

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// Ingredient is the model for the ingredient table
type Ingredient struct {
	ID       string          `json:"id,omitempty"`
	RecipeID string          `json:"recipeId,omitempty"`
	Name     string          `json:"name,omitempty"`
	Measure  string          `json:"measure,omitempty"`
	Quantity decimal.Decimal `json:"quantity,omitempty"`
}

// String is the string representation of an ingredient.Ingredient
func (i Ingredient) String() string {
	return fmt.Sprintf("%v %v %v", i.Quantity, i.Measure, i.Name)
}
