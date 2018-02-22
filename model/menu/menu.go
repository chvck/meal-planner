package menu

import (
	"github.com/chvck/meal-planner/model/recipe"
	null "gopkg.in/guregu/null.v3"
)

// Menu is the model for the menu table
type Menu struct {
	ID          int             `db:"id" json:"id"`
	UserID      int             `db:"user_id" json:"user_id"`
	Name        string          `db:"name" json:"name"`
	Description null.String     `db:"description" json:"description"`
	Recipes     []recipe.Recipe `json:"recipes"`
}
