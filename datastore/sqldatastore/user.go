package sqldatastore

import (
	"github.com/chvck/meal-planner/model"
	"strconv"
	"gopkg.in/guregu/null.v3"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type user struct {
	ID        int      `db:"id"`
	Username  string   `db:"username"`
	Email     string   `db:"email"`
	CreatedAt int      `db:"created_at"`
	UpdatedAt int      `db:"updated_at"`
	LastLogin null.Int `db:"last_login"`
}

func (u user) toModelUser() *model.User {
	return &model.User{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		LastLogin: int(u.LastLogin.Int64),
	}
}

func (ds *SQLDataStore) User(id int) (*model.User, error) {
	u := user{}
	err := ds.DB.Get(
		u,
		ds.DB.Rebind(`SELECT id, username, email, created_at, updated_at, last_login
		FROM "user"
		WHERE id = ?;`),
		strconv.Itoa(id),
	)
	if err != nil {
		return nil, err
	}

	return u.toModelUser(), nil
}

func (ds *SQLDataStore) Users(limit, offset int) ([]model.User, error) {
	var users []model.User
	rows, err := ds.DB.Queryx(
		ds.DB.Rebind(`SELECT id, username, email, created_at, updated_at, last_login
		FROM user
		ORDER BY id
		LIMIT ? OFFSET ?;`),
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		u := user{}
		err = rows.StructScan(&u)
		if err != nil {
			return make([]model.User, 0), nil
		}

		users = append(users, *u.toModelUser())
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (ds *SQLDataStore) UserCreate(u model.User, password []byte) (*model.User, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()

	var newU user
	err = ds.DB.Get(
		&newU,
		ds.DB.Rebind(
			`INSERT INTO "user" (username, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?) 
					RETURNING id, username, email, created_at, updated_at;`),
		u.Username, u.Email, string(hash), now, now,
	)
	if err != nil {
		return nil, err
	}

	return newU.toModelUser(), nil
}

func (ds *SQLDataStore) UserValidatePassword(username string, pw []byte) *model.User {
	var u struct {
		user
		Password string `db:"password"`
	}
	{
	}
	err := ds.DB.Get(
		&u,
		ds.DB.Rebind(`SELECT id, username, email, password, created_at, updated_at, last_login FROM "user" WHERE username = ?`), username,
	)
	if err != nil {
		return nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), pw)
	if err != nil {
		return nil
	}

	return u.toModelUser()
}
