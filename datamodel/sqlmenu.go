package datamodel

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/chvck/meal-planner/db"
	"github.com/chvck/meal-planner/model/menu"
)

// SQLMenu is a Menu datamodel backing onto a sql database
type SQLMenu struct {
	dataStore db.DataStoreAdapter
}

type menuWithPlannerID struct {
	menu.Menu
	PlannerID int
}

// One retrieves a single Menu by id
func (sqlm SQLMenu) One(id int, userID int) (*menu.Menu, error) {
	row := sqlm.dataStore.QueryOne(
		`SELECT m.id, m.name, m.description, m.user_id
		FROM menu m
		WHERE m.id = ? and m.user_id = ?;`,
		id,
		userID,
	)

	m := menu.Menu{}
	if err := row.Scan(&m.ID, &m.Name, &m.Description, &m.UserID); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &m, nil
}

// AllWithLimit retrieves x menus starting from an offset
func (sqlm SQLMenu) AllWithLimit(limit int, offset int, userID int) ([]menu.Menu, error) {
	var menus []menu.Menu
	rows, err := sqlm.dataStore.Query(`SELECT m.id, m.name, m.description, m.user_id
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
		m := menu.Menu{}
		if err := rows.Scan(&m.ID, &m.Name, &m.Description, &m.UserID); err != nil {
			return nil, err
		}

		menus = append(menus, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return menus, nil
}

// ForPlanners returns the menus for a list of planner IDs. Recipes are keyed by planner ID
func (sqlm SQLMenu) ForPlanners(ids ...interface{}) (map[int][]menu.Menu, error) {
	in := strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ",")
	var menus []menuWithPlannerID

	rows, err := sqlm.dataStore.Query(
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

		menus = append(menus, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	plannerIDToMenu := make(map[int][]menu.Menu)
	for _, m := range menus {

		_, ok := plannerIDToMenu[m.PlannerID]
		if !ok {
			plannerIDToMenu[m.PlannerID] = make([]menu.Menu, 0)
		}

		plannerIDToMenu[m.PlannerID] = append(plannerIDToMenu[m.PlannerID], m.Menu)
	}

	return plannerIDToMenu, nil
}

// Create creates the specific menu
func (sqlm SQLMenu) Create(m menu.Menu, userID int) (*int, error) {
	// if err := validateRecipe(r); err != nil {
	// 	return nil, err
	// }

	query := `INSERT INTO "menu" (name, description, user_id) VALUES (?, ?, ?) RETURNING id;`

	row := sqlm.dataStore.QueryOne(
		query,
		m.Name, m.Description, userID)

	var menuID int
	if err := row.Scan(&menuID); err != nil {
		return nil, err
	}

	return &menuID, nil
}

// Update updates the specific menu
func (sqlm SQLMenu) Update(m menu.Menu, id int, userID int) error {
	// if err := validateRecipe(r); err != nil {
	// 	return err
	// }

	if _, err := sqlm.dataStore.Exec(
		`UPDATE "menu" SET name = ?, description = ? WHERE id = ? and user_id = ?;`,
		m.Name, m.Description, id, userID); err != nil {
		return err
	}

	return nil
}

// Delete deletes the specific menu
func (sqlm SQLMenu) Delete(id int, userID int) error {
	rowsAccepted, err := sqlm.dataStore.Exec(
		`DELETE FROM "menu" m
		LEFT JOIN menu_to_recipe mr ON mr.menu_id = m.id
		LEFT JOIN planner_to_menu pm ON pm.menu_id = m.id
		WHERE m.id = ? and m.user_id = ?`, id, userID)
	if err != nil {
		return err
	}

	if rowsAccepted == 0 {
		return errors.New("No menu to delete")
	}

	return nil
}