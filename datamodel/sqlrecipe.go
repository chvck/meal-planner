package datamodel

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/chvck/meal-planner/db"
	"github.com/chvck/meal-planner/model"
	"github.com/shopspring/decimal"
	null "gopkg.in/guregu/null.v3"
)

// SQLRecipe is a Recipe datamodel backing onto a sql database.
// It also deals with ingredients as a part of recipes as the
// two are intrinsically linked together
type SQLRecipe struct {
	dataStore db.DataStoreAdapter
}

type recipeWithMenuID struct {
	model.Recipe
	MenuID int `db:"menu_id" json:"menu_id"`
}

type recipeWithPlannerID struct {
	model.Recipe
	PlannerID int `db:"planner_id" json:"planner_id"`
}

// NewSQLRecipe creates a new SQLRecipe datastore
func NewSQLRecipe(dataStore db.DataStoreAdapter) *SQLRecipe {
	return &SQLRecipe{dataStore: dataStore}
}

// FindByIngredientNames executes a search for recipes by ingredient name
func (sqlr SQLRecipe) FindByIngredientNames(names ...interface{}) ([]model.Recipe, error) {
	if len(names) == 0 {
		var recipes []model.Recipe
		return recipes, nil
	}

	m := make(map[int]*model.Recipe)
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

	rows, err := sqlr.dataStore.Query(query, names...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		r := model.Recipe{}
		rows.Scan(&r.ID, &r.Name, &r.Instructions, &r.Description, &r.Yield, &r.PrepTime, &r.CookTime, &r.UserID)

		m[r.ID] = &r
		ids = append(ids, r.ID)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(m) == 0 {
		var recipes []model.Recipe
		return recipes, nil
	}

	recipes := make([]model.Recipe, 0, len(m))
	ingredients, err := ingredientsForRecipes(sqlr.dataStore, ids...)

	if err != nil {
		return nil, err
	}
	for rID, i := range ingredients {
		r := m[rID]

		r.Ingredients = i
		recipes = append(recipes, *r)
	}

	return recipes, nil
}

// One retrieves a single model.Recipe by id
func (sqlr SQLRecipe) One(id int, userID int) (*model.Recipe, error) {
	row := sqlr.dataStore.QueryOne(
		`SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time, r.user_id
		FROM recipe r
		WHERE r.id = ? and r.user_id = ?;`,
		id,
		userID,
	)

	r := model.Recipe{}
	if err := row.Scan(&r.ID, &r.Name, &r.Instructions, &r.Description, &r.Yield,
		&r.PrepTime, &r.CookTime, &r.UserID); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	ingredients, err := ingredientsForRecipes(sqlr.dataStore, r.ID)
	if err != nil {
		return nil, err
	}
	if ingredients[r.ID] != nil {
		r.Ingredients = ingredients[r.ID]
	}

	return &r, nil
}

// AllWithLimit retrieves x recipes starting from an offset
func (sqlr SQLRecipe) AllWithLimit(limit int, offset int, userID int) ([]model.Recipe, error) {
	recipes := []model.Recipe{}
	var recipeIDs []interface{}
	rows, err := sqlr.dataStore.Query(`SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time, r.user_id
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
		r := model.Recipe{}
		if err := rows.Scan(&r.ID, &r.Name, &r.Instructions, &r.Description, &r.Yield, &r.PrepTime, &r.CookTime, &r.UserID); err != nil {
			return nil, err
		}

		recipeIDs = append(recipeIDs, r.ID)
		recipes = append(recipes, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(recipes) == 0 {
		return recipes, nil
	}

	ingredientsByRecipe, err := ingredientsForRecipes(sqlr.dataStore, recipeIDs...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	for i, recipe := range recipes {
		ingredients, ok := ingredientsByRecipe[recipe.ID]
		if ok {
			recipes[i].Ingredients = ingredients
		}
	}

	return recipes, nil
}

// ForMenus returns the recipes for a list of menu IDs. Recipes are keyed by menu ID
func (sqlr SQLRecipe) ForMenus(ids ...interface{}) (map[int][]model.Recipe, error) {
	in := strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ",")
	var recipeIDs []interface{}
	var recipes []recipeWithMenuID

	rows, err := sqlr.dataStore.Query(
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
		return make(map[int][]model.Recipe), nil
	}

	recipeIDToIngredients, err := ingredientsForRecipes(sqlr.dataStore, recipeIDs...)
	if err != nil {
		return nil, err
	}
	menuIDToRecipe := make(map[int][]model.Recipe)
	for _, rec := range recipes {
		rec.Ingredients = recipeIDToIngredients[rec.ID]

		_, ok := menuIDToRecipe[rec.MenuID]
		if !ok {
			menuIDToRecipe[rec.MenuID] = make([]model.Recipe, 0)
		}

		menuIDToRecipe[rec.MenuID] = append(menuIDToRecipe[rec.MenuID], rec.Recipe)
	}

	return menuIDToRecipe, nil
}

// ForPlanners returns the recipes for a list of planner IDs. Recipes are keyed by planner ID
func (sqlr SQLRecipe) ForPlanners(ids ...interface{}) (map[int][]model.Recipe, error) {
	in := strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ",")
	var recipeIDs []interface{}
	var recipes []recipeWithPlannerID

	rows, err := sqlr.dataStore.Query(
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
		return make(map[int][]model.Recipe), nil
	}

	recipeIDToIngredients, err := ingredientsForRecipes(sqlr.dataStore, recipeIDs...)
	if err != nil {
		return nil, err
	}
	plannerIDToRecipe := make(map[int][]model.Recipe)
	for _, rec := range recipes {
		rec.Ingredients = recipeIDToIngredients[rec.ID]

		_, ok := plannerIDToRecipe[rec.PlannerID]
		if !ok {
			plannerIDToRecipe[rec.PlannerID] = make([]model.Recipe, 0)
		}

		plannerIDToRecipe[rec.PlannerID] = append(plannerIDToRecipe[rec.PlannerID], rec.Recipe)
	}

	return plannerIDToRecipe, nil
}

// Create creates the specific recipe
func (sqlr SQLRecipe) Create(r model.Recipe, userID int) (*int, error) {
	tx, err := sqlr.dataStore.NewTransaction()
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

	if err := createManyIngredients(tx, r.Ingredients, recipeID); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &recipeID, err
}

// Update updates the specific recipe
func (sqlr SQLRecipe) Update(r model.Recipe, id int, userID int) error {
	row := sqlr.dataStore.QueryOne("SELECT user_id from recipe where id = ?;", id)

	var rUserID int
	if err := row.Scan(&rUserID); err != nil {
		return err
	}

	if rUserID != userID {
		return errors.New("cannot update recipe, unauthorized")
	}

	tx, err := sqlr.dataStore.NewTransaction()
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
	if err := deleteAllIngredientsByRecipe(tx, id); err != nil {
		tx.Rollback()
		return err
	}

	if err := createManyIngredients(tx, r.Ingredients, id); err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the specific recipe
func (sqlr SQLRecipe) Delete(id int, userID int) error {
	rowsAccepted, err := sqlr.dataStore.Exec(
		`DELETE FROM "recipe" r
		WHERE r.id = ? and r.user_id = ?`, id, userID)
	if err != nil {
		return err
	}

	if rowsAccepted == 0 {
		return errors.New("No recipe to delete")
	}

	return nil
}

func createManyIngredients(tx db.Transaction, ingredients []model.Ingredient, recipeID int) error {
	query := `INSERT INTO "ingredient" (name, measure, quantity, recipe_id) VALUES (?, ?, ?, ?);`
	for _, ing := range ingredients {
		if _, err := tx.Exec(query, ing.Name, ing.Measure, ing.Quantity, recipeID); err != nil {
			return err
		}
	}

	return nil
}

func deleteAllIngredientsByRecipe(tx db.Transaction, recipeID int) error {
	query := "DELETE FROM ingredient WHERE recipe_id = ?;"
	if _, err := tx.Exec(query, recipeID); err != nil {
		return err
	}

	return nil
}

func ingredientsForRecipes(dataStore db.DataStoreAdapter, ids ...interface{}) (map[int][]model.Ingredient, error) {
	m := make(map[int][]model.Ingredient)
	in := strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ",")

	query := fmt.Sprintf(
		`SELECT id, recipe_id, name, measure, quantity
		FROM ingredient
		WHERE recipe_id IN (%v)
		ORDER BY recipe_id, id;`,
		in,
	)

	rows, err := dataStore.Query(query, ids...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			rID     int
			ingID   int
			ingName string
			mName   null.String
			q       decimal.Decimal
		)
		if err := rows.Scan(&ingID, &rID, &ingName, &mName, &q); err != nil {
			return nil, err
		}

		arr := m[rID]
		i := model.Ingredient{ID: ingID, RecipeID: rID, Name: ingName, Measure: mName, Quantity: q}
		arr = append(arr, i)
		m[rID] = arr
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return m, nil
}
