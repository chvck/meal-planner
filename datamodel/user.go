package datamodel

// UserDataModel is the datamodel for data store operations on the User model
type UserDataModel interface {
	One(id int) (*model.User, error)
	AllWithLimit(limit int, offset int) ([]model.User, error)
	Create(u model.User, password []byte) (*int, error)
	ValidatePassword(username string, pw []byte) *model.User
}
