package recipe

import (
	"fmt"
	"github.com/chvck/meal-planner/model"
	"strings"
	"gopkg.in/guregu/null.v3"
	"sort"
	"github.com/chvck/meal-planner/model/ingredient"
)

// Recipe is the struct representing a Recipe
type Recipe struct {
	Id           int         `db:"id"`
	Name         string      `db:"name"`
	Instructions string      `db:"instructions"`
	Yield        null.Int    `db:"yield"`
	PrepTime     null.Int    `db:"prep_time"`
	CookTime     null.Int    `db:"cook_time"`
	Description  null.String `db:"description"`
	Ingredients  []ingredient.Ingredient
}

func NewRecipe() *Recipe {
	return &Recipe{Id: -1, Ingredients: []ingredient.Ingredient{}}
}

// Find executes a search for recipes using the where string built within the Finder
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
			rows.Scan(&r.Id, &r.Name, &r.Instructions, &r.Description, &r.Yield, &r.PrepTime, &r.CookTime)

			m[r.Id] = r
			ids = append(ids, r.Id)
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
		for rId, i := range ingredients {
			r := m[rId]

			r.Ingredients = i
			recipes = append(recipes, *r)
		}
	}

	return &recipes, nil
}

// One retrieves a single Recipe by id
func One(dataStore model.IDataStoreAdapter, id int) (*Recipe, error) {
	row := dataStore.QueryOne(
		`SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time
		FROM recipe r
		WHERE r.id = ?;`,
		id,
	)

	r := NewRecipe()
	if err := row.Scan(&r.Id, &r.Name, &r.Instructions, &r.Description, &r.Yield, &r.PrepTime, &r.CookTime); err != nil {
		return nil, err
	}

	var ids []interface{}
	ids = append(ids, r.Id)

	if ingredients, err := ingredientsByRecipe(dataStore, ids...); err != nil {
		return nil, err
	} else {
		if ingredients[r.Id] != nil {
			r.Ingredients = ingredients[r.Id]
		}
	}

	return r, nil
}

// All retrieves all recipes
func All(dataStore model.IDataStoreAdapter) (*[]Recipe, error) {
	return AllWithLimit(dataStore, "NULL", 0)
}

// AllWithLimit retrieves x recipes starting from an offset
// limit is expected to a positive int or string NULL (for no limit)
func AllWithLimit(dataStore model.IDataStoreAdapter, limit interface{}, offset int) (*[]Recipe, error) {
	m := make(map[int]*Recipe)
	var ids []interface{}
	if rows, err := dataStore.Query(fmt.Sprintf(
		`SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time
		FROM recipe r
		ORDER BY r.id
		LIMIT %v OFFSET %v;`,
		limit,
		offset,
	)); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		for rows.Next() {
			r := NewRecipe()
			rows.Scan(&r.Id, &r.Name, &r.Instructions, &r.Description, &r.Yield, &r.PrepTime, &r.CookTime)

			m[r.Id] = r
			ids = append(ids, r.Id)
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}
	}

	if len(m) == 0 {
		var recipes []Recipe
		return &recipes, nil
	}

	if ingredients, err := ingredientsByRecipe(dataStore, ids...); err != nil {
		return nil, err
	} else {
		for rId, i := range ingredients {
			r := m[rId]

			r.Ingredients = i
		}
	}
	recipes := make([]Recipe, 0, len(m))
	for _, recipe := range m {
		recipes = append(recipes, *recipe)
	}
	sort.SliceStable(recipes, func(i, j int) bool {
		return recipes[i].Id < recipes[j].Id
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
				rId     int
				ingId   int
				ingName string
				mName   null.String
				q       int
			)
			if err := rows.Scan(&ingId, &rId, &ingName, &mName, &q); err != nil {
				return nil, err
			}

			arr := m[rId]
			i := ingredient.Ingredient{Id: ingId, RecipeId: rId, Name: ingName, Measure: mName, Quantity: q}
			arr = append(arr, i)
			m[rId] = arr
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}
	}

	return m, nil
}

// Create creates the specific Recipe
func Create(dataStore model.IDataStoreAdapter, r Recipe) error {
	tx, err := dataStore.NewTransaction()
	if err != nil {
		return err
	}

	row := tx.QueryOne(
		"INSERT INTO recipe (name, instructions, yield, prep_time, cook_time, description) VALUES (?, ?, ?, ?, ?, ?) RETURNING id;",
		r.Name, r.Instructions, r.Yield, r.PrepTime, r.CookTime, r.Description)

	var recipeId int
	if err = row.Scan(&recipeId); err != nil {
		tx.Rollback()
		return err
	}

	if err := ingredient.CreateMany(tx, r.Ingredients, r.Id); err != nil {
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
func Update(dataStore model.IDataStoreAdapter, r Recipe) error {
	tx, err := dataStore.NewTransaction()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(
		"UPDATE recipe SET name = ?, instructions = ?, yield = ?, prep_time = ?, cook_time = ?, description = ? WHERE id = ?;",
		r.Name, r.Instructions, r.Yield, r.PrepTime, r.CookTime, r.Description, r.Id); err != nil {
		tx.Rollback()
		return err
	}

	// This isn't exactly efficient but ok for now
	if err := ingredient.DeleteAllByRecipe(tx, r.Id); err != nil {
		tx.Rollback()
		return err
	}

	if err := ingredient.CreateMany(tx, r.Ingredients, r.Id); err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

	return err

}
