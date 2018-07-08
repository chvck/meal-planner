package cbdatastore

import (
	"fmt"
	"time"

	"github.com/chvck/gocb"
	"github.com/chvck/meal-planner/model"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	model.User
	Type     string `json:"type,omitempty"`
	Password string `json:"password,omitempty"`
}

func (ds *CBDataStore) User(id string) (*model.User, error) {
	u := model.User{}
	_, err := ds.bucket.Get(id, &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (ds *CBDataStore) Users(limit, offset int) ([]model.User, error) {
	var users []model.User
	query := gocb.NewN1qlQuery(`SELECT id, username, email, created_at, updated_at, last_login
		FROM meals
		where type = "user"
		ORDER BY id
		LIMIT $1 OFFSET $2;`)
	results, err := ds.bucket.ExecuteN1qlQuery(query, [2]int{limit, offset})
	if err != nil {
		return nil, err
	}

	u := model.User{}
	for results.Next(&u) {
		users = append(users, u)
	}

	if err = results.Close(); err != nil {
		return nil, err
	}

	return users, nil
}

func (ds *CBDataStore) UserCreate(u model.User, password []byte) (*model.User, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("user::%s", u.Username)
	newU := user{}
	newU.Username = u.Username
	newU.ID = key
	newU.Email = u.Email

	now := time.Now().Unix()
	newU.CreatedAt = int(now)
	newU.UpdatedAt = int(now)
	newU.Password = string(hash)
	newU.Type = "user"

	_, err = ds.bucket.Insert(key, newU, 0)
	if err != nil {
		return nil, err
	}

	return &newU.User, nil
}

func (ds *CBDataStore) UserValidatePassword(username string, pw []byte) *model.User {
	var u user
	_, err := ds.bucket.Get(fmt.Sprintf("user::%s", username), &u)
	if err != nil {
		return nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), pw)
	if err != nil {
		return nil
	}

	return &u.User
}
