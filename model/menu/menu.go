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
	if err != nil {
		return nil, err
	}
	if recipes[m.ID] != nil {
		m.Recipes = recipes[m.ID]
	}
	return &m, nil
}

// AllWithLimit retrieves x menus starting from an offset
func AllWithLimit(dataStore model.IDataStoreAdapter, limit int, offset int, userID int) (*[]Menu, error) {
	idToMenu := make(map[int]*Menu)
	var menuIDs []interface{}
	rows, err := dataStore.Query(`SELECT m.id, m.name, m.description, m.user_id
		FROM menu m
		WHERE m.user_id = ?
		ORDER BY m.id
		LIMIT ? OFFSET ?;`,
		userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		m := Menu{}
		rows.Scan(&m.ID, &m.Name, &m.Description, &m.UserID)

		idToMenu[m.ID] = &m
		menuIDs = append(menuIDs, m.ID)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	// if there aren't any recipes then return empty slice
	if len(idToMenu) == 0 {
		menus := make([]Menu, 0)
		return &menus, nil
	}

	recipes, err := recipe.ForMenus(dataStore, menuIDs...)
	if err != nil {
		return nil, err
	}
	for mID, recs := range recipes {
		m := idToMenu[mID]

		m.Recipes = recs
	}

	menus := make([]Menu, 0, len(idToMenu))
	for _, m := range idToMenu {
		menus = append(menus, *m)
	}

	return &menus, nil
}
