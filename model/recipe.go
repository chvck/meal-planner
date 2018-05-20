package model

import (
	"errors"
)

// Recipe is the model for the recipe table
type Recipe struct {
	ID           int          `json:"id,omitempty"`
	UserID       int          `json:"userId,omitempty"`
	Name         string       `json:"name,omitempty"`
	Instructions string       `json:"instructions,omitempty"`
	Yield        int          `json:"yield,omitempty"`
	PrepTime     int          `json:"prepTime,omitempty"`
	CookTime     int          `json:"cookTime,omitempty"`
	Description  string       `json:"description,omitempty"`
	Ingredients  []Ingredient `json:"ingredients,omitempty"`
}

// Validate checks that the recipe is valid
func (r Recipe) Validate() []error {
	var errs []error
	if r.Name == "" {
		errs = append(errs, errors.New("name cannot be empty"))
	}
	if r.Instructions == "" {
		errs = append(errs, errors.New("instructions cannot be empty"))
	}

	return errs
}
