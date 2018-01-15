package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/chvck/meal-planner/model"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v3"
)

type User struct {
	Id        int      `db:"id" json:"id"`
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
	if err := row.Scan(&u.Id, &u.Username, &u.Email, &u.CreatedAt, &u.UpdatedAt, &u.LastLogin); err != nil {
		return nil, err
	}

	return &u, nil
}

// All retrieves all users
func All(dataStore model.IDataStoreAdapter) (*[]User, error) {
	return AllWithLimit(dataStore, "NULL", 0)
}

// AllWithLimit retrieves x users starting from an offset
// limit is expected to a positive int or string NULL (for no limit)
func AllWithLimit(dataStore model.IDataStoreAdapter, limit interface{}, offset int) (*[]User, error) {
	var users []User
	if rows, err := dataStore.Query(fmt.Sprintf(
		`SELECT id, username, email, created_at, updated_at, last_login
		FROM user
		ORDER BY id
		LIMIT %v OFFSET %v;`,
		limit,
		offset,
	)); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		for rows.Next() {
			u := User{}
			rows.Scan(&u.Id, &u.Username, &u.Email, &u.CreatedAt, &u.UpdatedAt, &u.LastLogin)

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

	var userId int
	if err = row.Scan(&userId); err != nil {
		return err
	}

	return nil
}

func ValidatePassword(dataStore model.IDataStoreAdapter, username string, pw []byte) *User {
	row := dataStore.QueryOne(
		`SELECT id, username, email, password, created_at, updated_at, last_login FROM "user" WHERE username = ?`, username,
	)

	var actualPw string
	var u User
	if err := row.Scan(&u.Id, &u.Username, &u.Email, &actualPw, &u.CreatedAt, &u.UpdatedAt, &u.LastLogin); err != nil {
		return nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(actualPw), pw); err != nil {
		return nil
	} else {
		return &u
	}
}
