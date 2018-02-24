package datamodel

import (
	"errors"
	"time"

	"github.com/chvck/meal-planner/db"
	"golang.org/x/crypto/bcrypt"
)

// SQLUser is a User datamodel backing onto a sql database
type SQLUser struct {
	dataStore db.DataStoreAdapter
}

const (
	saltSize   = 64
	iterations = 1e4
)

// One retrieves a single User by id
func (sqlu SQLUser) One(id int) (*model.User, error) {
	row := sqlu.dataStore.QueryOne(
		`SELECT id, username, email, created_at, updated_at, last_login
		FROM user
		WHERE id = ?;`,
		id,
	)

	u := model.User{}
	if err := row.Scan(&u.ID, &u.Username, &u.Email, &u.CreatedAt, &u.UpdatedAt, &u.LastLogin); err != nil {
		return nil, err
	}

	return &u, nil
}

// AllWithLimit retrieves x users starting from an offset
func (sqlu SQLUser) AllWithLimit(limit int, offset int) ([]model.User, error) {
	var users []model.User
	rows, err := sqlu.dataStore.Query(
		`SELECT id, username, email, created_at, updated_at, last_login
		FROM user
		ORDER BY id
		LIMIT ? OFFSET ?;`,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		u := model.User{}
		rows.Scan(&u.ID, &u.Username, &u.Email, &u.CreatedAt, &u.UpdatedAt, &u.LastLogin)

		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Create persists the specific User
func (sqlu SQLUser) Create(u model.User, password []byte) (*int, error) {
	if string(password) == "" {
		return nil, errors.New("password cannot be empty")
	}

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()

	row := sqlu.dataStore.QueryOne(
		`INSERT INTO "user" (username, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?) RETURNING id;`,
		u.Username, u.Email, string(hash), now, now,
	)

	var userID int
	row.Scan(&userID)
	return &userID, nil
}

// ValidatePassword verifies a password for a user
func (sqlu SQLUser) ValidatePassword(username string, pw []byte) *model.User {
	row := sqlu.dataStore.QueryOne(
		`SELECT id, username, email, password, created_at, updated_at, last_login FROM "user" WHERE username = ?`, username,
	)

	var actualPw string
	var u model.User
	if err := row.Scan(&u.ID, &u.Username, &u.Email, &actualPw, &u.CreatedAt, &u.UpdatedAt, &u.LastLogin); err != nil {
		return nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(actualPw), pw); err != nil {
		return nil
	}

	return &u
}
