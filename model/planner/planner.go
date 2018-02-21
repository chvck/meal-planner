package planner

import (
	"database/sql"
	"errors"

	"github.com/chvck/meal-planner/db"
	"github.com/chvck/meal-planner/model/menu"
	"github.com/chvck/meal-planner/model/recipe"
)

// Planner is the model for the planner table
type Planner struct {
	ID      int             `db:"id" json:"id"`
	UserID  int             `db:"user_id" json:"user_id"`
	When    int             `db:"when" json:"when"`
	For     string          `db:"for" json:"for"`
	Menus   []menu.Menu     `json:"menus"`
	Recipes []recipe.Recipe `json:"recipes"`
}

// All retrieves Planners between two dates
func All(dataStore db.DataStoreAdapter, start int, end int, userID int) (*[]Planner, error) {
	rows, err := dataStore.Query(
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

	var planners []Planner
	var plannerIDs []interface{}
	for rows.Next() {
		p := Planner{}
		if err := rows.Scan(&p.ID, &p.UserID, &p.When, &p.For); err != nil {
			return nil, err
		}

		planners = append(planners, p)
		plannerIDs = append(plannerIDs, p.ID)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	// if there aren't any planners then return empty slice
	if len(planners) == 0 {
		return &planners, nil
	}

	recipesByPlannerID, err := recipe.ForPlanners(dataStore, plannerIDs...)
	if err != nil {
		return nil, err
	}

	menusByPlannerID, err := menu.ForPlanners(dataStore, plannerIDs...)
	if err != nil {
		return nil, err
	}

	for i, p := range planners {
		recipes, ok := recipesByPlannerID[p.ID]
		if ok {
			planners[i].Recipes = recipes
		}
		menus, ok := menusByPlannerID[p.ID]
		if ok {
			planners[i].Menus = menus
		}
	}

	return &planners, nil
}

// AddMenu adds a menu to the planner, if the planner when/mealtime pair doesn't exist then creates it
func AddMenu(dataStore db.DataStoreAdapter, when int, mealtime string, menuID int, userID int) error {
	if !validateMealtime(mealtime) {
		return errors.New("mealtime must be one of breakfast, lunch, dinner or snack")
	}

	pID, err := getAndCreatePlanner(dataStore, when, mealtime, userID)
	if err != nil {
		return err
	}

	query := `INSERT INTO "planner_to_menu" ("planner_id", "menu_id") VALUES (?, ?) RETURNING id;`

	_, err = dataStore.Exec(query, pID, menuID)
	if err != nil {
		return err
	}

	return nil
}

// AddRecipe adds a recipe to the planner, if the planner when/mealtime pair doesn't exist then creates it
func AddRecipe(dataStore db.DataStoreAdapter, when int, mealtime string, recipeID int, userID int) error {
	if !validateMealtime(mealtime) {
		return errors.New("mealtime must be one of breakfast, lunch, dinner or snack")
	}

	pID, err := getAndCreatePlanner(dataStore, when, mealtime, userID)
	if err != nil {
		return err
	}

	query := `INSERT INTO "planner_to_recipe" ("planner_id", "recipe_id") VALUES (?, ?) RETURNING id;`

	_, err = dataStore.Exec(query, pID, recipeID)
	if err != nil {
		return err
	}

	return nil
}

// RemoveMenu removes a menu from the planner
func RemoveMenu(dataStore db.DataStoreAdapter, when int, mealtime string, menuID int, userID int) error {
	query := `DELETE pm FROM "planner_to_menu" pm JOIN "planner" p 
	ON p.id = pm.planner_id WHERE p.when = ? AND p.for = ? AND p.user_id = ? AND pm.menu_id = ?;`

	if _, err := dataStore.Exec(query, when, mealtime, userID, menuID); err != nil {
		return err
	}

	return nil
}

// RemoveRecipe removes a recipe from the planner
func RemoveRecipe(dataStore db.DataStoreAdapter, when int, mealtime string, recipeID int, userID int) error {
	query := `DELETE pr FROM "planner_to_recipe" pr JOIN "planner" p 
	ON p.id = pm.planner_id WHERE p.when = ? AND p.for = ? AND p.user_id = ? AND pr.recipe_id = ?;`

	if _, err := dataStore.Exec(query, when, mealtime, userID, recipeID); err != nil {
		return err
	}

	return nil
}

func getAndCreatePlanner(dataStore db.DataStoreAdapter, when int, mealtime string, userID int) (*int, error) {
	p, err := plannerExists(dataStore, when, mealtime)
	if err != nil {
		return nil, err
	}

	var pID int
	if p == nil {
		createPQuery := `INSERT INTO "planner" ("user_id", "when", "for") VALUES (?, ?, ?) RETURNING id;`
		row := dataStore.QueryOne(createPQuery, userID, when, mealtime)

		if err := row.Scan(&pID); err != nil {
			return nil, err
		}
	} else {
		pID = p.ID
	}

	return &pID, nil

}

func plannerExists(dataStore db.DataStoreAdapter, when int, mealtime string) (*Planner, error) {
	query := `SELECT p.id, p.when, p.for, p.user_id FROM planner p WHERE "when" = ? AND "for" = ?;`

	row := dataStore.QueryOne(query, when, mealtime)

	var p Planner
	if err := row.Scan(&p.ID, &p.When, &p.For, &p.UserID); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &p, nil
}

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
