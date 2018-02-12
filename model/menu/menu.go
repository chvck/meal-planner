package menu

import (
	"database/sql"

	"github.com/chvck/meal-planner/model"
	"github.com/chvck/meal-planner/model/recipe"
	null "gopkg.in/guregu/null.v3"
)

// Menu is the model for the menu table
type Menu struct {
	ID          int             `db:"id" json:"id"`
	UserID      int             `db:"user_id" json:"user_id"`
	Name        string          `db:"name" json:"name"`
	Description null.String     `db:"description" json:"description"`
	Recipes     []recipe.Recipe `json:"recipes"`
}

// One retrieves a single Menu by id
func One(dataStore model.IDataStoreAdapter, id int, userID int) (*Menu, error) {
	row := dataStore.QueryOne(
		`SELECT m.id, m.name, m.description, m.user_id
		FROM menu m
		WHERE m.id = ? and m.user_id = ?;`,
		id,
		userID,
	)

	m := Menu{}
	if err := row.Scan(&m.ID, &m.Name, &m.Description, &m.UserID); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	recipes, err := recipe.ForMenus(dataStore, m.ID)
	if err == nil {
		return nil, err
	}
	if recipes[m.ID] != nil {
		m.Recipes = recipes[m.ID]
	}
	return &m, nil
}
