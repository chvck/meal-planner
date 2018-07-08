package model

import (
	"errors"
)

// User is the model for the user table
type User struct {
	ID        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	CreatedAt int    `json:"createdAt,omitempty"`
	UpdatedAt int    `json:"updatedAt,omitempty"`
	LastLogin int    `json:"lastLogin,omitempty"`
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
