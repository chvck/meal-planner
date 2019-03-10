package model

import (
	"github.com/pkg/errors"
)

// Validate checks that the planner is valid
func (p Planner) Validate() []error {
	var errs []error
	if p.Date == 0 {
		errs = append(errs, errors.New("date cannot be empty"))
	}
	if _, ok := Planner_Mealtime_name[int32(p.Mealtime)]; !ok {
		errs = append(errs, errors.New("mealtime cannot be empty"))
	}

	return errs
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

// Levels are for user access levels
const (
	LevelUser  = 1.0
	LevelAdmin = 2.0
)

// Validate checks that the user is valid
func (u User) Validate() []error {
	var errs []error
	if u.Username == "" {
		errs = append(errs, errors.New("username cannot be empty"))
	}
	if u.Email == "" {
		errs = append(errs, errors.New("email cannot be empty"))
	}

	return errs
}

// ValidatePassword checks that the password is valid
func ValidatePassword(pwd string) error {
	if len(pwd) < 8 {
		return errors.New("password must be longer than 8 characters")
	}

	return nil
}

