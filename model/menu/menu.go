package menu

import (
	"database/sql"
	"fmt"
	"strings"

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

type menuWithPlannerID struct {
	Menu
	PlannerID int
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
	var menuIDs []interface{}
	var menus []Menu
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
		if err := rows.Scan(&m.ID, &m.Name, &m.Description, &m.UserID); err != nil {
			return nil, err
		}

		menuIDs = append(menuIDs, m.ID)
		menus = append(menus, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	// if there aren't any menus then return empty slice
	if len(menus) == 0 {
		return &menus, nil
	}

	recipesByMenuID, err := recipe.ForMenus(dataStore, menuIDs...)
	if err != nil {
		return nil, err
	}

	for i, m := range menus {
		recipes, ok := recipesByMenuID[m.ID]
		if ok {
			menus[i].Recipes = recipes
		}
	}

	return &menus, nil
}

// ForPlanners returns the menus for a list of planner IDs. Recipes are keyed by planner ID
func ForPlanners(dataStore model.IDataStoreAdapter, ids ...interface{}) (map[int][]Menu, error) {
	in := strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ",")
	var menuIDs []interface{}
	var menus []menuWithPlannerID

	rows, err := dataStore.Query(
		fmt.Sprintf(`SELECT m.id, m.name, m.description, m.user_id, pm.planner_id
				FROM menu m
				JOIN planner_to_menu pm ON pm.menu_id = m.id
				WHERE pm.planner_id IN (%v)
				ORDER BY pm.planner_id, m.id`,
			in), ids...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		m := menuWithPlannerID{}
		if err := rows.Scan(&m.ID, &m.Name, &m.Description, &m.UserID, &m.PlannerID); err != nil {
			return nil, err
		}

		menuIDs = append(menuIDs, m.ID)
		menus = append(menus, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(menus) == 0 {
		return make(map[int][]Menu), nil
	}

	recipesByMenuID, err := recipe.ForMenus(dataStore, menuIDs...)
	if err != nil {
		return nil, err
	}

	plannerIDToMenu := make(map[int][]Menu)
	for _, m := range menus {
		m.Recipes = recipesByMenuID[m.ID]

		_, ok := plannerIDToMenu[m.PlannerID]
		if !ok {
			plannerIDToMenu[m.PlannerID] = make([]Menu, 0)
		}

		plannerIDToMenu[m.PlannerID] = append(plannerIDToMenu[m.PlannerID], m.Menu)
	}

	return plannerIDToMenu, nil
}
