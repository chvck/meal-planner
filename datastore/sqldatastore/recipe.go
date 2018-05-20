package sqldatastore

import (
	"gopkg.in/guregu/null.v3"
	"github.com/shopspring/decimal"
	"github.com/chvck/meal-planner/model"
	"fmt"
	"database/sql"
	"strings"
	"errors"
	"github.com/jmoiron/sqlx"
)

//type recipeWithMenuID struct {
//	recipe
//	MenuID int `db:"menu_id"`
//}
//
//type recipeWithPlannerID struct {
//	recipe
//	PlannerID int `db:"planner_id"`
//}

type recipe struct {
	ID           int         `db:"id"`
	UserID       int         `db:"user_id"`
	Name         string      `db:"name"`
	Instructions string      `db:"instructions"`
	Yield        null.Int    `db:"yield"`
	PrepTime     null.Int    `db:"prep_time"`
	CookTime     null.Int    `db:"cook_time"`
	Description  null.String `db:"description"`
	Ingredients  []ingredient
}

type ingredient struct {
	ID       int             `db:"id"`
	RecipeID int             `db:"recipe_id"`
	Name     string          `db:"name"`
	Measure  null.String     `db:"measure"`
	Quantity decimal.Decimal `db:"quantity"`
}

func (i *ingredient) toModelIngredient() *model.Ingredient {
	modelIng := model.Ingredient{}
	modelIng.RecipeID = i.RecipeID
	modelIng.Name = i.Name
	modelIng.Measure = i.Measure.String
	modelIng.Quantity = i.Quantity

	return &modelIng
}

func ingredientFromModelIngredient(modelIng model.Ingredient) *ingredient {
	i := ingredient{}
	i.ID = modelIng.ID
	i.RecipeID = modelIng.RecipeID
	i.Name = modelIng.Name
	i.Measure = null.StringFrom(modelIng.Measure)
	i.Quantity = modelIng.Quantity

	return &i
}

func (r recipe) toModelRecipe() *model.Recipe {
	modelRecipe := model.Recipe{}
	modelRecipe.ID = r.ID
	modelRecipe.UserID = r.UserID
	modelRecipe.Name = r.Name
	modelRecipe.Instructions = r.Instructions
	modelRecipe.Yield = int(r.Yield.Int64)
	modelRecipe.PrepTime = int(r.PrepTime.Int64)
	modelRecipe.CookTime = int(r.CookTime.Int64)
	modelRecipe.Description = r.Description.String

	return &modelRecipe
}

func recipeFromModelRecipe(modelRecipe model.Recipe) *recipe {
	r := recipe{}
	r.ID = modelRecipe.ID
	r.UserID = modelRecipe.UserID
	r.Name = modelRecipe.Name
	r.ID = modelRecipe.ID
	r.Instructions = modelRecipe.Instructions
	r.Yield = null.IntFrom(int64(modelRecipe.Yield))
	r.PrepTime = null.IntFrom(int64(modelRecipe.PrepTime))
	r.CookTime = null.IntFrom(int64(modelRecipe.CookTime))
	r.Description = null.StringFrom(modelRecipe.Description)

	r.Ingredients = make([]ingredient, len(modelRecipe.Ingredients))

	for i, ing := range modelRecipe.Ingredients {
		r.Ingredients[i] = *ingredientFromModelIngredient(ing)
	}

	return &r
}

// RecipesFromIngredientNames executes a search for recipes by ingredient name
func (ds SQLDataStore) RecipesFromIngredientNames(names []string) ([]model.Recipe, error) {
	if len(names) == 0 {
		return make([]model.Recipe, 0), nil
	}

	m := make(map[int]model.Recipe)
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

	rows, err := ds.DB.Queryx(ds.DB.Rebind(query), names)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		r := recipe{}
		rows.StructScan(&r)

		m[r.ID] = *r.toModelRecipe()
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
	ingredients, err := ds.ingredientsForRecipes(ids...)

	if err != nil {
		return nil, err
	}
	for rID, i := range ingredients {
		r := m[rID]
		r.Ingredients = i
		recipes = append(recipes, r)
	}

	return recipes, nil
}

// One retrieves a single model.Recipe by id
func (ds SQLDataStore) Recipe(id, userID int) (*model.Recipe, error) {
	r := recipe{}
	err := ds.DB.Get(&r,
		ds.DB.Rebind(`SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time, r.user_id
		FROM recipe r
		WHERE r.id = ? and r.user_id = ?;`),
		id,
		userID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	ingredients, err := ds.ingredientsForRecipes(r.ID)
	if err != nil {
		return nil, err
	}

	modelR := r.toModelRecipe()
	if ingredients[r.ID] != nil {
		modelR.Ingredients = ingredients[r.ID]
	}

	return modelR, nil
}

// AllWithLimit retrieves x recipes starting from an offset
func (ds SQLDataStore) Recipes(limit, offset, userID int) ([]model.Recipe, error) {
	var recipes []model.Recipe
	var recipeIDs []interface{}
	rows, err := ds.DB.Queryx(
		ds.DB.Rebind(`SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time, r.user_id
		FROM recipe r
		WHERE r.user_id = ?
		ORDER BY r.id
		LIMIT ? OFFSET ?;`),
		userID, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		r := recipe{}
		err = rows.StructScan(&r)
		if err != nil {
			return nil, err
		}

		recipeIDs = append(recipeIDs, r.ID)
		recipes = append(recipes, *r.toModelRecipe())
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(recipes) == 0 {
		return recipes, nil
	}

	ingredientsByRecipe, err := ds.ingredientsForRecipes(recipeIDs...)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
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

// Create creates the specific recipe
func (ds SQLDataStore) RecipeCreate(modelRecipe model.Recipe, userID int) (*model.Recipe, error) {
	tx, err := ds.DB.Beginx()
	if err != nil {
		return nil, err
	}

	r := recipeFromModelRecipe(modelRecipe)
	query := "INSERT INTO recipe (name, instructions, yield, prep_time, cook_time, description, user_id) VALUES (?, ?, ?, ?, ?, ?, ?)" +
		"RETURNING id, name, instructions, yield, prep_time, cook_time, description, user_id;"

	var newR recipe
	err = tx.Get(&newR, tx.Rebind(query), r.Name, r.Instructions, r.Yield, r.PrepTime, r.CookTime, r.Description, userID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	ings, err := createManyIngredients(tx, r.Ingredients, newR.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	modelR := newR.toModelRecipe()
	modelR.Ingredients = ings
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return modelR, err
}

// Update updates the specific recipe
func (ds SQLDataStore) RecipeUpdate(modelRecipe model.Recipe, id, userID int) error {
	var rUserID int
	err := ds.DB.Get(&rUserID, ds.DB.Rebind("SELECT user_id from recipe where id = ?;"), id)
	if err != nil {
		return err
	}

	r := recipeFromModelRecipe(modelRecipe)

	if rUserID != userID {
		return errors.New("cannot update recipe, unauthorized")
	}

	tx, err := ds.DB.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		tx.Rebind("UPDATE recipe SET name = ?, instructions = ?, yield = ?, prep_time = ?, cook_time = ?, description = ? WHERE id = ? and user_id = ?;"),
		r.Name, r.Instructions, r.Yield, r.PrepTime, r.CookTime, r.Description, id, r.UserID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// This isn't exactly efficient but ok for now
	err = deleteAllIngredientsByRecipe(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = createManyIngredients(tx, r.Ingredients, id)
	if err != nil {
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
func (ds SQLDataStore) RecipeDelete(id int, userID int) error {
	result, err := ds.DB.Exec(
		ds.DB.Rebind(`DELETE FROM "recipe" r
		WHERE r.id = ? and r.user_id = ?`), id, userID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("no recipe to delete")
	}

	return nil
}

func createManyIngredients(tx *sqlx.Tx, ingredients []ingredient, recipeID int) ([]model.Ingredient, error) {
	query := `INSERT INTO "ingredient" (name, measure, quantity, recipe_id) VALUES (?, ?, ?, ?)
				RETURNING id, name, measure, quantity, recipe_id;`

	var ings []model.Ingredient
	for _, ing := range ingredients {
		var i ingredient
		err := tx.Get(&i, tx.Rebind(query), ing.Name, ing.Measure, ing.Quantity, recipeID)
		if err != nil {
			return nil, err
		}

		ings = append(ings, *i.toModelIngredient())
	}

	return ings, nil
}

func deleteAllIngredientsByRecipe(tx *sqlx.Tx, recipeID int) error {
	query := "DELETE FROM ingredient WHERE recipe_id = ?;"
	if _, err := tx.Exec(tx.Rebind(query), recipeID); err != nil {
		return err
	}

	return nil
}

func (ds SQLDataStore) ingredientsForRecipes(ids ...interface{}) (map[int][]model.Ingredient, error) {
	m := make(map[int][]model.Ingredient)
	in := strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ",")

	query := fmt.Sprintf(
		`SELECT id, recipe_id, name, measure, quantity
		FROM ingredient
		WHERE recipe_id IN (%v)
		ORDER BY recipe_id, id;`,
		in,
	)

	rows, err := ds.DB.Queryx(ds.DB.Rebind(query), ids...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var i ingredient
		err := rows.StructScan(&i)
		if err != nil {
			return nil, err
		}

		arr := m[i.RecipeID]
		arr = append(arr, model.Ingredient{
			ID:       i.ID,
			RecipeID: i.RecipeID,
			Name:     i.Name,
			Measure:  i.Measure.String,
			Quantity: i.Quantity,
		})
		m[i.RecipeID] = arr
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return m, nil
}
