package model

import (
	"errors"

	"gopkg.in/guregu/null.v3"
)

// User is the model for the user table
type User struct {
	ID        int      `db:"id" json:"id"`
	Username  string   `db:"username" json:"username"`
	Email     string   `db:"email" json:"email"`
	CreatedAt int      `db:"created_at" json:"createdAt"`
	UpdatedAt int      `db:"updated_at" json:"updatedAt"`
	LastLogin null.Int `db:"last_login" json:"lastLogin"`
}

// Levels are for user access levels
const (
	LevelUser  = 1.0
	LevelAdmin = 2.0
)

// Validate checks that the user is valid
func (u User) Validate() []error {
	errs := []error{}
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
		return errors.New("password cannot be empty")
	}

	return nil
}
