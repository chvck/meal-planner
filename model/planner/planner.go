package planner

import (
	"github.com/chvck/meal-planner/model/menu"
	"github.com/chvck/meal-planner/model/recipe"
)

// Planner is the model for the planner table
type Planner struct {
	ID      int             `db:"id" json:"id"`
	UserID  int             `db:"user_id" json:"user_id"`
	When    int             `db:"when" json:"when"`
	For     string          `db:"for" json:"for"`
	Menus   []menu.Menu     `json:"menus"`
	Recipes []recipe.Recipe `json:"recipes"`
}
