package model

import (
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
