package recipe

import (
	"database/sql"
	"fmt"
	"github.com/chvck/meal-planner/model"
	"strings"
)

// Recipe is the struct representing a recipe
type Recipe struct {
	Id           int            `db:"id"`
	Name         string         `db:"name"`
	Instructions string         `db:"instructions"`
	Yield        sql.NullInt64  `db:"yield"`
	PrepTime     sql.NullInt64  `db:"prep_time"`
	CookTime     sql.NullInt64  `db:"cook_time"`
	Description  sql.NullString `db:"description"`
}

// Find is a struct used for finding recipes
type Finder struct {
	where string
	vals  []interface{}
}

var table = "recipe"

// One retrieves a single Recipe by id
func One(dataStore model.IDataStoreAdapter, id int) (Recipe, error) {
	r := Recipe{}
	err := dataStore.QueryOne(&r, fmt.Sprintf("SELECT * FROM %v WHERE id = ?", table), id)

	return r, err
}

// Where sets up the Finder with a where clause containing the specified values
func (rf *Finder) Where(where string, operator string, val string) *Finder {
	where = strings.TrimLeft(where, "WHERE ")
	rf.where = fmt.Sprintf("WHERE %v %v ?", where, operator)
	rf.vals = append(rf.vals, val)

	return rf
}

// AndWhere extends an existing where clause on a Finder with an extra clause that will be connected with an and selection
func (rf *Finder) AndWhere(where string, operator string, val string) *Finder {
	rf.where = fmt.Sprintf("%v AND %v %v ?", rf.where, where, operator)
	rf.vals = append(rf.vals, val)

	return rf
}

// OrWhere extends an existing where clause on a Finder with an extra clause that will be connected with an or selection
func (rf *Finder) OrWhere(where string, operator string, val string) *Finder {
	rf.where = fmt.Sprintf("%v OR %v %v ?", rf.where, where, operator)
	rf.vals = append(rf.vals, val)

	return rf
}

// Find executes a search for recipes using the where string built within the Finder
func (rf *Finder) Find(dataStore model.IDataStoreAdapter) ([]Recipe, error) {
	var r []Recipe
	err := dataStore.Query(&r, fmt.Sprintf("SELECT * FROM %v %v", table, rf.where), rf.vals...)

	return r, err
}

// All retrieves all Recipes
func All(dataStore model.IDataStoreAdapter) ([]Recipe, error) {
	var r []Recipe
	err := dataStore.Query(&r, fmt.Sprintf("SELECT * FROM %v", table))

	return r, err
}

// AllWithLimit retrieves x Recipes starting from an offset
func AllWithLimit(dataStore model.IDataStoreAdapter, limit int, offset int) ([]Recipe, error) {
	var r []Recipe
	err := dataStore.Query(&r, fmt.Sprintf("SELECT * FROM %v LIMIT %v OFFSET %v", table, limit, offset))

	return r, err
}

// Save persists the specific Recipe
// TODO: Consider adding a Save function to the IDataStoreAdapter which uses reflection to accept an interface and then iterate over fields for updates/saves
func Save(dataStore model.IDataStoreAdapter, r Recipe) error {
	if r.Id == 0 {
		_, err := dataStore.Exec(fmt.Sprintf(
			"INSERT INTO %v (name, instructions, yield, prep_time, cook_time, description) VALUES (?, ?, ?, ?, ?, ?);",
			table,
		), r.Name, r.Instructions, r.Yield, r.PrepTime, r.CookTime, r.Description)

		return err
	} else {
		_, err := dataStore.Exec(fmt.Sprintf(
			"UPDATE %v SET name = ?, instructions = ?, yield = ?, prep_time = ?, cook_time = ?, description = ? WHERE id = ?;",
			table,
		), r.Name, r.Instructions, r.Yield, r.PrepTime, r.CookTime, r.Description, r.Id)

		return err
	}

}
