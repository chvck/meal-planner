package service

import (
	"github.com/chvck/meal-planner/datamodel"
	"github.com/chvck/meal-planner/model"
)

type UserService struct {
	udm datamodel.UserDataModel
}

// GetByID retrieves a user by id
func (us UserService) GetByID(id int, userID int) (*model.User, error) {
	return us.udm.One(id)
}

// All retrieves all users
func (us UserService) All(limit int, offset int, userID int) ([]model.User, error) {
	return us.udm.AllWithLimit(limit, offset)
}

// Create creates a new user
func (us UserService) Create(u model.User, password []byte) (*model.User, error) {
	uID, err := us.udm.Create(u, password)
	if err != nil {
		return nil, err
	}

	return us.udm.One(*uID)
}

// ValidatePassword validates a password for a user
func (us UserService) ValidatePassword(username string, password []byte) *model.User {
	return us.udm.ValidatePassword(username, password)
}
