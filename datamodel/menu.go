package datamodel

import (
	"github.com/chvck/meal-planner/model/menu"
)

// MenuDataModel is the datamodel for data store operations on the Menu model
type MenuDataModel interface {
	One(id int, userID int) (*menu.Menu, error)
	AllWithLimit(limit int, offset int, userID int) ([]menu.Menu, error)
	ForPlanners(ids ...interface{}) (map[int][]menu.Menu, error)
	Create(menu.Menu, int) (*int, error)
	Update(menu.Menu, int, int) error
	Delete(int, userID int) error
}
