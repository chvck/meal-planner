package datamodel

import "github.com/chvck/meal-planner/model"

// MenuDataModel is the datamodel for data store operations on the Menu model
type MenuDataModel interface {
	One(id int, userID int) (*model.Menu, error)
	AllWithLimit(limit int, offset int, userID int) ([]model.Menu, error)
	ForPlanners(ids ...interface{}) (map[int][]model.Menu, error)
	Create(model.Menu, int) (*int, error)
	Update(model.Menu, int, int) error
	Delete(int, userID int) error
}
