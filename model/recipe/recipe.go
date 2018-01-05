package recipe

import (
	"database/sql"
	"fmt"
	"github.com/chvck/meal-planner/model"
	"strings"
)

// recipe is the struct representing a recipe
type recipe struct {
	Id           int            `db:"id"`
	Name         string         `db:"name"`
	Instructions string         `db:"instructions"`
	Yield        sql.NullInt64  `db:"yield"`
	PrepTime     sql.NullInt64  `db:"prep_time"`
	CookTime     sql.NullInt64  `db:"cook_time"`
	Description  sql.NullString `db:"description"`
	Ingredients  []ingredient
}

type ingredient struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Measure  string
	Quantity int
}

func (i ingredient) String() string {
	return fmt.Sprintf("%v %v %v", i.Quantity, i.Measure, i.Name)
}

func NewRecipe() *recipe {
	return &recipe{Id: -1}
}

// Find executes a search for recipes using the where string built within the Finder
func FindByIngredientNames(dataStore model.IDataStoreAdapter, names ...interface{}) (*[]recipe, error) {
	if len(names) == 0 {
		var recipes []recipe
		return &recipes, nil
	}

	m := make(map[int]*recipe)
	var ids []interface{}
	where := "i.name = ?"
	for i := 0; i < len(names[1:]); i++ {
		where = fmt.Sprintf("%v OR i.name = ?", where)
	}
	query := fmt.Sprintf(
		`SELECT DISTINCT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time
		FROM ingredient i
		JOIN recipe_to_ingredient ri ON ri.ingredient_id = i.id
		JOIN recipe r ON r.id = ri.recipe_id
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
	}

	if len(m) == 0 {
		var recipes []recipe
		return &recipes, nil
	}

	recipes := make([]recipe, 0, len(m))
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

// One retrieves a single recipe by id
func One(dataStore model.IDataStoreAdapter, id int) (*recipe, error) {
	r := NewRecipe()
	row := dataStore.QueryOne(
		`SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time
		FROM recipe r
		WHERE r.id = ?;`,
		id,
	)

	if err := row.Scan(&r.Id, &r.Name, &r.Instructions, &r.Description, &r.Yield, &r.PrepTime, &r.CookTime); err != nil {
		return nil, err
	}

	var ids []interface{}
	ids = append(ids, r.Id)

	if ingredients, err := ingredientsByRecipe(dataStore, ids...); err != nil {
		return nil, err
	} else {
		r.Ingredients = ingredients[r.Id]
	}

	return r, nil
}

// All retrieves all recipes
func All(dataStore model.IDataStoreAdapter) (*[]recipe, error) {
	return AllWithLimit(dataStore, "NULL", 0)
}

// AllWithLimit retrieves x recipes starting from an offset
// limit is expected to a positive int or string NULL (for no limit)
func AllWithLimit(dataStore model.IDataStoreAdapter, limit interface{}, offset int) (*[]recipe, error) {
	m := make(map[int]*recipe)
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
	}

	if len(m) == 0 {
		var recipes []recipe
		return &recipes, nil
	}

	recipes := make([]recipe, 0, len(m))
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

func ingredientsByRecipe(dataStore model.IDataStoreAdapter, ids ...interface{}) (map[int][]ingredient, error) {
	m := make(map[int][]ingredient)
	in := strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ",")

	query := fmt.Sprintf(
		`SELECT ri.recipe_id, i.id, i.name, m.name, quantity
		FROM recipe_to_ingredient ri
		JOIN ingredient i on i.id = ri.ingredient_id
		JOIN measure m on m.id = ri.measure_id
		WHERE ri.recipe_id IN (%v)
		ORDER BY ri.recipe_id;`,
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
				mName   string
				q       int
			)
			if err := rows.Scan(&rId, &ingId, &ingName, &mName, &q); err != nil {
				return nil, err
			}

			arr := m[rId]
			i := ingredient{Id: ingId, Name: ingName, Measure: mName, Quantity: q}
			arr = append(arr, i)
			m[rId] = arr
		}
	}

	return m, nil
}

// Save persists the specific recipe
// TODO: Consider adding a Save function to the IDataStoreAdapter which uses reflection to accept an interface and then iterate over fields for updates/saves
func Save(dataStore model.IDataStoreAdapter, r recipe) error {
	if r.Id == 0 {
		_, err := dataStore.Exec(
			"INSERT INTO recipe (name, instructions, yield, prep_time, cook_time, description) VALUES (?, ?, ?, ?, ?, ?);",
			r.Name, r.Instructions, r.Yield, r.PrepTime, r.CookTime, r.Description)
		return err
	} else {
		_, err := dataStore.Exec(
			"UPDATE name SET name = ?, instructions = ?, yield = ?, prep_time = ?, cook_time = ?, description = ? WHERE id = ?;",
			r.Name, r.Instructions, r.Yield, r.PrepTime, r.CookTime, r.Description, r.Id)

		return err
	}

}
