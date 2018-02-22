package datamodel

import (
	"database/sql"

	"github.com/chvck/meal-planner/db"
	"github.com/chvck/meal-planner/model/planner"
)

// SQLPlanner is a Planner datamodel backing onto a sql database
type SQLPlanner struct {
	dataStore db.DataStoreAdapter
}

// All retrieves Planners between two dates
func (sqlp SQLPlanner) All(start int, end int, userID int) ([]planner.Planner, error) {
	rows, err := sqlp.dataStore.Query(
		`SELECT p.id, p.user_id, p.when, p.for
		FROM planner p
		WHERE p.user_id = ? AND p.when BETWEEN ? AND ?
		ORDER BY p.id;`,
		userID,
		start,
		end,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var planners []planner.Planner
	for rows.Next() {
		p := planner.Planner{}
		if err := rows.Scan(&p.ID, &p.UserID, &p.When, &p.For); err != nil {
			return nil, err
		}

		planners = append(planners, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return planners, nil
}

// One retrieves a single Planner
func (sqlp SQLPlanner) One(when int, mealtime string, userID int) (*planner.Planner, error) {
	query := `SELECT p.id, p.when, p.for, p.user_id FROM planner p WHERE "when" = ? AND "for" = ?;`

	row := sqlp.dataStore.QueryOne(query, when, mealtime)

	var p planner.Planner
	if err := row.Scan(&p.ID, &p.When, &p.For, &p.UserID); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &p, nil
}

// Create creates a new Planner
func (sqlp SQLPlanner) Create(when int, mealtime string, userID int) (*int, error) {
	query := `INSERT INTO "planner" (when, mealtime, user_id) VALUES (?, ?, ?) RETURNING id;`

	row := sqlp.dataStore.QueryOne(
		query,
		when, mealtime, userID)

	var menuID int
	if err := row.Scan(&menuID); err != nil {
		return nil, err
	}

	return &menuID, nil
}

// AddRecipe adds a recipe to a planner
func (sqlp SQLPlanner) AddRecipe(plannerID int, recipeID int) error {
	query := `INSERT INTO "planner_to_recipe" ("planner_id", "recipe_id") VALUES (?, ?) RETURNING id;`

	_, err := sqlp.dataStore.Exec(query, plannerID, recipeID)
	if err != nil {
		return err
	}

	return nil
}

// RemoveRecipe removes a recipe from the planner
func (sqlp SQLPlanner) RemoveRecipe(when int, mealtime string, recipeID int, userID int) error {
	query := `DELETE pr FROM "planner_to_recipe" pr JOIN "planner" p
	ON p.id = pm.planner_id WHERE p.when = ? AND p.for = ? AND p.user_id = ? AND pr.recipe_id = ?;`

	if _, err := sqlp.dataStore.Exec(query, when, mealtime, userID, recipeID); err != nil {
		return err
	}

	return nil
}

// AddMenu adds a menu to the planner, if the planner when/mealtime pair doesn't exist then creates it
// func (sqlp SQLPlanner) AddMenu(when int, mealtime string, menuID int, userID int) error {
// 	if !validateMealtime(mealtime) {
// 		return errors.New("mealtime must be one of breakfast, lunch, dinner or snack")
// 	}

// 	pID, err := getAndCreatePlanner(dataStore, when, mealtime, userID)
// 	if err != nil {
// 		return err
// 	}

// 	query := `INSERT INTO "planner_to_menu" ("planner_id", "menu_id") VALUES (?, ?) RETURNING id;`

// 	_, err = sqlp.dataStore.Exec(query, pID, menuID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // RemoveMenu removes a menu from the planner
// func RemoveMenu(dataStore db.DataStoreAdapter, when int, mealtime string, menuID int, userID int) error {
// 	query := `DELETE pm FROM "planner_to_menu" pm JOIN "planner" p
// 	ON p.id = pm.planner_id WHERE p.when = ? AND p.for = ? AND p.user_id = ? AND pm.menu_id = ?;`

// 	if _, err := dataStore.Exec(query, when, mealtime, userID, menuID); err != nil {
// 		return err
// 	}

// 	return nil
// }

func validateMealtime(mealtime string) bool {
	switch mealtime {
	case
		"breakfast",
		"lunch",
		"dinner",
		"snack":
		return true
	}

	return false
}
