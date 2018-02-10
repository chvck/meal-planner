package user

import (
	"errors"
	"time"

	"github.com/chvck/meal-planner/model"
	"golang.org/x/crypto/bcrypt"
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

const (
	saltSize   = 64
	iterations = 1e4
)

// One retrieves a single User by id
func One(dataStore model.IDataStoreAdapter, id int) (*User, error) {
	row := dataStore.QueryOne(
		`SELECT id, username, email, created_at, updated_at, last_login
		FROM user
		WHERE id = ?;`,
		id,
	)

	u := User{}
	if err := row.Scan(&u.ID, &u.Username, &u.Email, &u.CreatedAt, &u.UpdatedAt, &u.LastLogin); err != nil {
		return nil, err
	}

	return &u, nil
}

// AllWithLimit retrieves x users starting from an offset
func AllWithLimit(dataStore model.IDataStoreAdapter, limit int, offset int) (*[]User, error) {
	var users []User
	if rows, err := dataStore.Query(
		`SELECT id, username, email, created_at, updated_at, last_login
		FROM user
		ORDER BY id
		LIMIT ? OFFSET ?;`,
		limit,
		offset,
	); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		for rows.Next() {
			u := User{}
			rows.Scan(&u.ID, &u.Username, &u.Email, &u.CreatedAt, &u.UpdatedAt, &u.LastLogin)

			users = append(users, u)
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}
	}

	return &users, nil
}

// Create persists the specific User
func Create(dataStore model.IDataStoreAdapter, u User, password []byte) error {
	if string(password) == "" {
		return errors.New("password cannot be empty")
	}

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	now := time.Now().Unix()

	row := dataStore.QueryOne(
		`INSERT INTO "user" (username, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?) RETURNING id;`,
		u.Username, u.Email, string(hash), now, now,
	)

	var userID int
	return row.Scan(&userID)
}

// ValidatePassword verifies a password for a user
func ValidatePassword(dataStore model.IDataStoreAdapter, username string, pw []byte) *User {
	row := dataStore.QueryOne(
		`SELECT id, username, email, password, created_at, updated_at, last_login FROM "user" WHERE username = ?`, username,
	)

	var actualPw string
	var u User
	if err := row.Scan(&u.ID, &u.Username, &u.Email, &actualPw, &u.CreatedAt, &u.UpdatedAt, &u.LastLogin); err != nil {
		return nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(actualPw), pw); err != nil {
		return nil
	} else {
		return &u
	}
}
