package datamodel

import "github.com/chvck/meal-planner/model/user"

// UserDataModel is the datamodel for data store operations on the User model
type UserDataModel interface {
	One(id int) (*user.User, error)
	AllWithLimit(limit int, offset int) ([]user.User, error)
	Create(u user.User, password []byte) (*int, error)
	ValidatePassword(username string, pw []byte) *user.User
}
