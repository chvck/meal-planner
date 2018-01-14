package ingredient

import (
	"gopkg.in/guregu/null.v3"
	"github.com/chvck/meal-planner/model"
	"fmt"
)

type Ingredient struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
}

type IngredientWithProps struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Measure  null.String
	Quantity int
}

func (i IngredientWithProps) String() string {
	return fmt.Sprintf("%v %v %v", i.Quantity, i.Measure, i.Name)
}

func NewIngredient() *Ingredient {
	return &Ingredient{Id: -1}
}

// All retrieves all ingredients
func All(dataStore model.IDataStoreAdapter) (*[]Ingredient, error) {
	return AllWithLimit(dataStore, "NULL", 0)
}

// AllWithLimit retrieves x ingredients starting from an offset
// limit is expected to a positive int or string NULL (for no limit)
func AllWithLimit(dataStore model.IDataStoreAdapter, limit interface{}, offset int) (*[]Ingredient, error) {
	var ingredients []Ingredient
	if rows, err := dataStore.Query(fmt.Sprintf(
		`SELECT id, name
		FROM ingredient
		ORDER BY id
		LIMIT %v OFFSET %v;`,
		limit,
		offset,
	)); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		for rows.Next() {
			i := NewIngredient()
			rows.Scan(&i.Id, &i.Name)

			ingredients = append(ingredients, *i)
		}
	}

	return &ingredients, nil
}
