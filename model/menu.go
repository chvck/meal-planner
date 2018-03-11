package model

import (
	"errors"

	null "gopkg.in/guregu/null.v3"
)

// Menu is the model for the menu table
type Menu struct {
	ID          int         `db:"id" json:"id"`
	UserID      int         `db:"user_id" json:"user_id"`
	Name        string      `db:"name" json:"name"`
	Description null.String `db:"description" json:"description"`
	Recipes     []Recipe    `json:"recipes"`
}

// Validate checks that the menu is valid
func (m Menu) Validate() []error {
	errs := []error{}
	if m.Name == "" {
		errs = append(errs, errors.New("name cannot be empty"))
	}

	return errs
}
