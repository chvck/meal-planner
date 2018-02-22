package recipe

import (
	"fmt"

	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v3"
)

// Recipe is the model for the recipe table
type Recipe struct {
	ID           int          `db:"id" json:"id"`
	UserID       int          `db:"user_id" json:"user_id"`
	Name         string       `db:"name" json:"name"`
	Instructions string       `db:"instructions" json:"instructions"`
	Yield        null.Int     `db:"yield" json:"yield"`
	PrepTime     null.Int     `db:"prep_time" json:"prep_time"`
	CookTime     null.Int     `db:"cook_time" json:"cook_time"`
	Description  null.String  `db:"description" json:"description"`
	Ingredients  []Ingredient `json:"ingredients"`
}

// Ingredient is the model for the ingredient table, it is only ever used as a part of Recipe
// and by using Recipe as a proxy
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
