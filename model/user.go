package model

import (
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
