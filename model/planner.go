package model

import "errors"

// Planner is the model for the planner table
type Planner struct {
	ID          string       `json:"id,omitempty"`
	UserID      string       `json:"userID,omitempty"`
	When        int          `json:"when,omitempty"`
	For         string       `json:"for,omitempty"`
	RecipeNames []RecipeName `json:"recipes,omitempty"`
}

type RecipeName struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Validate checks that the planner is valid
func (p Planner) Validate() []error {
	var errs []error
	if p.When == 0 {
		errs = append(errs, errors.New("when cannot be empty"))
	}
	if p.For == "" {
		errs = append(errs, errors.New("for cannot be empty"))
	}

	return errs
}
