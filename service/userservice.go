package service

import (
	"github.com/chvck/meal-planner/datamodel"
	"github.com/chvck/meal-planner/model"
)

// UserService is the service for interacting with Users
type UserService interface {
	GetByID(id int, userID int) (*model.User, error)
	All(limit int, offset int, userID int) ([]model.User, error)
	Create(u model.User, password []byte) (*model.User, error)
	ValidatePassword(username string, password []byte) *model.User
}

type userService struct {
	udm datamodel.UserDataModel
}

// NewUserService creates a new user service
func NewUserService(udm datamodel.UserDataModel) UserService {
	return &userService{udm: udm}
}

// GetByID retrieves a user by id
func (us userService) GetByID(id int, userID int) (*model.User, error) {
	return us.udm.One(id)
}

// All retrieves all users
func (us userService) All(limit int, offset int, userID int) ([]model.User, error) {
	return us.udm.AllWithLimit(limit, offset)
}

// Create creates a new user
func (us userService) Create(u model.User, password []byte) (*model.User, error) {
	uID, err := us.udm.Create(u, password)
	if err != nil {
		return nil, err
	}

	return us.udm.One(*uID)
}

// ValidatePassword validates a password for a user
func (us userService) ValidatePassword(username string, password []byte) *model.User {
	return us.udm.ValidatePassword(username, password)
}
