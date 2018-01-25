package recipe

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"

	"github.com/chvck/meal-planner/model"
	"github.com/chvck/meal-planner/model/ingredient"
	"gopkg.in/guregu/null.v3"
)

// Recipe is the model for the recipe table
type Recipe struct {
	ID           int         `db:"id"`
	UserID       int         `db:"user_id"`
	Name         string      `db:"name"`
	Instructions string      `db:"instructions"`
	Yield        null.Int    `db:"yield"`
	PrepTime     null.Int    `db:"prep_time"`
	CookTime     null.Int    `db:"cook_time"`
	Description  null.String `db:"description"`
	Ingredients  []ingredient.Ingredient
}

// NewRecipe creates a new Recipe
func NewRecipe() *Recipe {
	return &Recipe{ID: -1, Ingredients: []ingredient.Ingredient{}}
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
		`SELECT DISTINCT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time
		FROM ingredient i
		JOIN recipe r ON r.id = i.recipe_id
		WHERE %v;`,
		where,
	)

	if rows, err := dataStore.Query(query, names...); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		for rows.Next() {
			r := NewRecipe()
			rows.Scan(&r.ID, &r.Name, &r.Instructions, &r.Description, &r.Yield, &r.PrepTime, &r.CookTime)

			m[r.ID] = r
			ids = append(ids, r.ID)
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}
	}

	if len(m) == 0 {
		var recipes []Recipe
		return &recipes, nil
	}

	recipes := make([]Recipe, 0, len(m))
	if ingredients, err := ingredientsByRecipe(dataStore, ids...); err != nil {
		return nil, err
	} else {
		for rID, i := range ingredients {
			r := m[rID]

			r.Ingredients = i
			recipes = append(recipes, *r)
		}
	}

	return &recipes, nil
}

// One retrieves a single Recipe by id
func One(dataStore model.IDataStoreAdapter, id int, userID int) (*Recipe, error) {
	row := dataStore.QueryOne(
		`SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time
		FROM recipe r
		WHERE r.id = ? and r.user_id = ?;`,
		id,
		userID,
	)

	r := NewRecipe()
	if err := row.Scan(&r.ID, &r.Name, &r.Instructions, &r.Description, &r.Yield, &r.PrepTime, &r.CookTime); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var ids []interface{}
	ids = append(ids, r.ID)

	if ingredients, err := ingredientsByRecipe(dataStore, ids...); err == nil {
		if ingredients[r.ID] != nil {
			r.Ingredients = ingredients[r.ID]
		}
		return r, nil
	} else {
		return nil, err
	}
}

// All retrieves all recipes
func All(dataStore model.IDataStoreAdapter, userID int) (*[]Recipe, error) {
	return AllWithLimit(dataStore, "NULL", 0, userID)
}

// AllWithLimit retrieves x recipes starting from an offset
// limit is expected to a positive int or string NULL (for no limit)
func AllWithLimit(dataStore model.IDataStoreAdapter, limit interface{}, offset int, userID int) (*[]Recipe, error) {
	m := make(map[int]*Recipe)
	var ids []interface{}
	if rows, err := dataStore.Query(fmt.Sprintf(
		`SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time
		FROM recipe r
		WHERE r.user_id = ?
		ORDER BY r.id
		LIMIT %v OFFSET %v;`,
		limit,
		offset,
	), userID); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		for rows.Next() {
			r := NewRecipe()
			rows.Scan(&r.ID, &r.Name, &r.Instructions, &r.Description, &r.Yield, &r.PrepTime, &r.CookTime)

			m[r.ID] = r
			ids = append(ids, r.ID)
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}
	}

	if len(m) == 0 {
		recipes := make([]Recipe, 0, len(m))
		return &recipes, nil
	}

	if ingredients, err := ingredientsByRecipe(dataStore, ids...); err != nil {
		return nil, err
	} else {
		for rID, i := range ingredients {
			r := m[rID]

			r.Ingredients = i
		}
	}
	recipes := make([]Recipe, 0, len(m))
	for _, recipe := range m {
		recipes = append(recipes, *recipe)
	}
	sort.SliceStable(recipes, func(i, j int) bool {
		return recipes[i].ID < recipes[j].ID
	})

	return &recipes, nil
}

func ingredientsByRecipe(dataStore model.IDataStoreAdapter, ids ...interface{}) (map[int][]ingredient.Ingredient, error) {
	m := make(map[int][]ingredient.Ingredient)
	in := strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ",")

	query := fmt.Sprintf(
		`SELECT id, recipe_id, name, measure, quantity
		FROM ingredient
		WHERE recipe_id IN (%v)
		ORDER BY recipe_id;`,
		in,
	)

	if rows, err := dataStore.Query(query, ids...); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		for rows.Next() {
			var (
				rID     int
				ingID   int
				ingName string
				mName   null.String
				q       int
			)
			if err := rows.Scan(&ingID, &rID, &ingName, &mName, &q); err != nil {
				return nil, err
			}

			arr := m[rID]
			i := ingredient.Ingredient{ID: ingID, RecipeID: rID, Name: ingName, Measure: mName, Quantity: q}
			arr = append(arr, i)
			m[rID] = arr
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}
	}

	return m, nil
}

// Create creates the specific Recipe
func Create(dataStore model.IDataStoreAdapter, r Recipe, userID int) error {
	tx, err := dataStore.NewTransaction()
	if err != nil {
		return err
	}

	row := tx.QueryOne(
		"INSERT INTO recipe (name, instructions, yield, prep_time, cook_time, description, user_id) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING id;",
		r.Name, r.Instructions, r.Yield, r.PrepTime, r.CookTime, r.Description, userID)

	var recipeID int
	if err = row.Scan(&recipeID); err != nil {
		tx.Rollback()
		return err
	}

	if err := ingredient.CreateMany(tx, r.Ingredients, r.ID); err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Update updates the specific Recipe
func Update(dataStore model.IDataStoreAdapter, r Recipe, userID int) error {
	tx, err := dataStore.NewTransaction()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(
		"UPDATE recipe SET name = ?, instructions = ?, yield = ?, prep_time = ?, cook_time = ?, description = ? WHERE id = ? and user_id = ?;",
		r.Name, r.Instructions, r.Yield, r.PrepTime, r.CookTime, r.Description, r.ID, userID); err != nil {
		tx.Rollback()
		return err
	}

	// This isn't exactly efficient but ok for now
	if err := ingredient.DeleteAllByRecipe(tx, r.ID); err != nil {
		tx.Rollback()
		return err
	}

	if err := ingredient.CreateMany(tx, r.Ingredients, r.ID); err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
