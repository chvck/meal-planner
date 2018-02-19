package recipe

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/chvck/meal-planner/model"
	"github.com/chvck/meal-planner/model/ingredient"
	"gopkg.in/guregu/null.v3"
)

// Recipe is the model for the recipe table
type Recipe struct {
	ID           int                     `db:"id" json:"id"`
	UserID       int                     `db:"user_id" json:"user_id"`
	Name         string                  `db:"name" json:"name"`
	Instructions string                  `db:"instructions" json:"instructions"`
	Yield        null.Int                `db:"yield" json:"yield"`
	PrepTime     null.Int                `db:"prep_time" json:"prep_time"`
	CookTime     null.Int                `db:"cook_time" json:"cook_time"`
	Description  null.String             `db:"description" json:"description"`
	Ingredients  []ingredient.Ingredient `json:"ingredients"`
}

type recipeWithMenuID struct {
	Recipe
	MenuID int `db:"menu_id" json:"menu_id"`
}

type recipeWithPlannerID struct {
	Recipe
	PlannerID int `db:"planner_id" json:"planner_id"`
}

// FindByIngredientNames executes a search for recipes by ingredient name
func FindByIngredientNames(dataStore model.IDataStoreAdapter, names ...interface{}) (*[]Recipe, error) {
	if len(names) == 0 {
		var recipes []Recipe
		return &recipes, nil
	}

	m := make(map[int]*Recipe)
	var ids []interface{}
	where := "i.name = ?"
	for i := 0; i < len(names[1:]); i++ {
		where = fmt.Sprintf("%v OR i.name = ?", where)
	}
	query := fmt.Sprintf(
		`SELECT DISTINCT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time, r.user_id
		FROM ingredient i
		JOIN recipe r ON r.id = i.recipe_id
		WHERE %v;`,
		where,
	)

	rows, err := dataStore.Query(query, names...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		r := Recipe{}
		rows.Scan(&r.ID, &r.Name, &r.Instructions, &r.Description, &r.Yield, &r.PrepTime, &r.CookTime, &r.UserID)

		m[r.ID] = &r
		ids = append(ids, r.ID)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(m) == 0 {
		var recipes []Recipe
		return &recipes, nil
	}

	recipes := make([]Recipe, 0, len(m))
	ingredients, err := ingredient.ForRecipes(dataStore, ids...)

	if err != nil {
		return nil, err
	}
	for rID, i := range ingredients {
		r := m[rID]

		r.Ingredients = i
		recipes = append(recipes, *r)
	}

	return &recipes, nil
}

// One retrieves a single Recipe by id
func One(dataStore model.IDataStoreAdapter, id int, userID int) (*Recipe, error) {
	row := dataStore.QueryOne(
		`SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time, r.user_id
		FROM recipe r
		WHERE r.id = ? and r.user_id = ?;`,
		id,
		userID,
	)

	r := Recipe{}
	if err := row.Scan(&r.ID, &r.Name, &r.Instructions, &r.Description, &r.Yield, &r.PrepTime, &r.CookTime, &r.UserID); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	ingredients, err := ingredient.ForRecipes(dataStore, r.ID)
	if err != nil {
		return nil, err
	}
	if ingredients[r.ID] != nil {
		r.Ingredients = ingredients[r.ID]
	}
	return &r, nil
}

// AllWithLimit retrieves x recipes starting from an offset
func AllWithLimit(dataStore model.IDataStoreAdapter, limit int, offset int, userID int) (*[]Recipe, error) {
	var recipeIDs []interface{}
	var recipes []Recipe
	rows, err := dataStore.Query(`SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time, r.user_id
		FROM recipe r
		WHERE r.user_id = ?
		ORDER BY r.id
		LIMIT ? OFFSET ?;`,
		userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		r := Recipe{}
		if err := rows.Scan(&r.ID, &r.Name, &r.Instructions, &r.Description, &r.Yield, &r.PrepTime, &r.CookTime, &r.UserID); err != nil {
			return nil, err
		}

		recipes = append(recipes, r)
		recipeIDs = append(recipeIDs, r.ID)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	// if there aren't any recipes then return empty slice
	if len(recipes) == 0 {
		return &recipes, nil
	}

	ingredientsByRecipe, err := ingredient.ForRecipes(dataStore, recipeIDs...)
	if err != nil {
		return nil, err
	}

	for i, recipe := range recipes {
		ingredients, ok := ingredientsByRecipe[recipe.ID]
		if ok {
			recipes[i].Ingredients = ingredients
		}
	}

	return &recipes, nil
}

// ForMenus returns the recipes for a list of menu IDs. Recipes are keyed by menu ID
func ForMenus(dataStore model.IDataStoreAdapter, ids ...interface{}) (map[int][]Recipe, error) {
	in := strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ",")
	var recipeIDs []interface{}
	var recipes []recipeWithMenuID

	rows, err := dataStore.Query(
		fmt.Sprintf(`SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time, r.user_id, mr.menu_id
				FROM recipe r
				JOIN menu_to_recipe mr ON mr.recipe_id = r.id
				WHERE mr.menu_id IN (%v)
				ORDER BY mr.menu_id, r.id`,
			in), ids...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		r := recipeWithMenuID{}
		if err := rows.Scan(&r.ID, &r.Name, &r.Instructions, &r.Description, &r.Yield, &r.PrepTime, &r.CookTime, &r.UserID, &r.MenuID); err != nil {
			return nil, err
		}

		recipeIDs = append(recipeIDs, r.ID)
		recipes = append(recipes, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(recipes) == 0 {
		return make(map[int][]Recipe), nil
	}

	recipeIDToIngredients, err := ingredient.ForRecipes(dataStore, recipeIDs...)
	if err != nil {
		return nil, err
	}
	menuIDToRecipe := make(map[int][]Recipe)
	for _, rec := range recipes {
		rec.Ingredients = recipeIDToIngredients[rec.ID]

		_, ok := menuIDToRecipe[rec.MenuID]
		if !ok {
			menuIDToRecipe[rec.MenuID] = make([]Recipe, 0)
		}

		menuIDToRecipe[rec.MenuID] = append(menuIDToRecipe[rec.MenuID], rec.Recipe)
	}

	return menuIDToRecipe, nil
}

// ForPlanners returns the recipes for a list of planner IDs. Recipes are keyed by planner ID
func ForPlanners(dataStore model.IDataStoreAdapter, ids ...interface{}) (map[int][]Recipe, error) {
	in := strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ",")
	var recipeIDs []interface{}
	var recipes []recipeWithPlannerID

	rows, err := dataStore.Query(
		fmt.Sprintf(`SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time, r.user_id, pr.planner_id
				FROM recipe r
				JOIN planner_to_recipe pr ON pr.recipe_id = r.id
				WHERE pr.planner_id IN (%v)
				ORDER BY pr.planner_id, r.id`,
			in), ids...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		r := recipeWithPlannerID{}
		if err := rows.Scan(&r.ID, &r.Name, &r.Instructions, &r.Description, &r.Yield, &r.PrepTime, &r.CookTime, &r.UserID, &r.PlannerID); err != nil {
			return nil, err
		}

		recipeIDs = append(recipeIDs, r.ID)
		recipes = append(recipes, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(recipes) == 0 {
		return make(map[int][]Recipe), nil
	}

	recipeIDToIngredients, err := ingredient.ForRecipes(dataStore, recipeIDs...)
	if err != nil {
		return nil, err
	}
	plannerIDToRecipe := make(map[int][]Recipe)
	for _, rec := range recipes {
		rec.Ingredients = recipeIDToIngredients[rec.ID]

		_, ok := plannerIDToRecipe[rec.PlannerID]
		if !ok {
			plannerIDToRecipe[rec.PlannerID] = make([]Recipe, 0)
		}

		plannerIDToRecipe[rec.PlannerID] = append(plannerIDToRecipe[rec.PlannerID], rec.Recipe)
	}

	return plannerIDToRecipe, nil
}

// Create creates the specific Recipe
func Create(dataStore model.IDataStoreAdapter, r Recipe, userID int) (*int, error) {
	if err := validate(r); err != nil {
		return nil, err
	}

	tx, err := dataStore.NewTransaction()
	if err != nil {
		return nil, err
	}

	query := "INSERT INTO recipe (name, instructions, yield, prep_time, cook_time, description, user_id) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING id;"

	row := tx.QueryOne(
		query,
		r.Name, r.Instructions, r.Yield, r.PrepTime, r.CookTime, r.Description, userID)

	var recipeID int
	if err = row.Scan(&recipeID); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := ingredient.CreateMany(tx, r.Ingredients, recipeID); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &recipeID, err
}

// Update updates the specific Recipe
func Update(dataStore model.IDataStoreAdapter, r Recipe, id int, userID int) error {
	if err := validate(r); err != nil {
		return err
	}

	tx, err := dataStore.NewTransaction()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(
		"UPDATE recipe SET name = ?, instructions = ?, yield = ?, prep_time = ?, cook_time = ?, description = ? WHERE id = ? and user_id = ?;",
		r.Name, r.Instructions, r.Yield, r.PrepTime, r.CookTime, r.Description, id, userID); err != nil {
		tx.Rollback()
		return err
	}

	// This isn't exactly efficient but ok for now
	if err := ingredient.DeleteAllByRecipe(tx, id); err != nil {
		tx.Rollback()
		return err
	}

	if err := ingredient.CreateMany(tx, r.Ingredients, id); err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func validate(r Recipe) error {
	if r.Name == "" {
		return errors.New("name cannot be empty")
	}
	if r.Instructions == "" {
		return errors.New("instructions cannot be empty")
	}

	return nil
}
