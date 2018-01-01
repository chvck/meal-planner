package recipe

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type RecipeTestAdapter struct {
	query      string
	bindVars   []interface{}
	destTypeOk bool
}

func (r RecipeTestAdapter) Initialize(connectionString string) error {
	return nil
}

func (r *RecipeTestAdapter) Query(dest interface{}, baseQuery string, bindVars ...interface{}) error {
	r.query = baseQuery
	r.bindVars = bindVars
	r.destTypeOk = false
	recipes, ok := dest.(*[]Recipe)
	if ok {
		r.destTypeOk = true

		*recipes = append(*recipes, Recipe{Id: 1, Name: "test"})
		*recipes = append(*recipes, Recipe{Id: 2, Name: "test2"})
	}

	return nil
}

func (r *RecipeTestAdapter) QueryOne(dest interface{}, baseQuery string, bindVars ...interface{}) error {
	r.query = baseQuery
	r.bindVars = bindVars
	r.destTypeOk = false
	recipe, ok := dest.(*Recipe)
	if ok {
		r.destTypeOk = true

		recipe.Id = 1
		recipe.Name = "test recipe"
	}

	return nil
}

func (r RecipeTestAdapter) Exec(baseExec string, bindVars ...interface{}) (int, error) {
	return 0, nil
}

func TestOne(t *testing.T) {
	adapter := &RecipeTestAdapter{}
	recipe, err := One(adapter, 1)

	assert.Nil(t, err)
	assert.NotNil(t, recipe)
	assert.Equal(t, 1, recipe.Id)
	assert.Equal(t, "SELECT * FROM recipe WHERE id = ?", adapter.query)
	assert.Equal(t, 1, len(adapter.bindVars))
	assert.Equal(t, 1, adapter.bindVars[0])
	assert.True(t, adapter.destTypeOk)
}

func TestAll(t *testing.T) {
	adapter := &RecipeTestAdapter{}
	recipes, err := All(adapter)

	assert.Nil(t, err)
	assert.NotNil(t, recipes)
	assert.Equal(t, 2, len(recipes))
	assert.Equal(t, 1, recipes[0].Id)
	assert.Equal(t, 2, recipes[1].Id)
	assert.Equal(t, "SELECT * FROM recipe", adapter.query)
	assert.Equal(t, 0, len(adapter.bindVars))
	assert.True(t, adapter.destTypeOk)
}

func TestAllWithLimit(t *testing.T) {
	adapter := &RecipeTestAdapter{}
	recipes, err := AllWithLimit(adapter, 10, 5)

	assert.Nil(t, err)
	assert.NotNil(t, recipes)
	assert.Equal(t, 2, len(recipes))
	assert.Equal(t, 1, recipes[0].Id)
	assert.Equal(t, 2, recipes[1].Id)
	assert.Equal(t, "SELECT * FROM recipe LIMIT 10 OFFSET 5", adapter.query)
	assert.Equal(t, 0, len(adapter.bindVars))
	assert.True(t, adapter.destTypeOk)
}

func TestFind(t *testing.T) {
	adapter := &RecipeTestAdapter{}
	rf := Finder{}
	recipes, err := rf.Find(adapter)

	assert.Nil(t, err)
	assert.NotNil(t, recipes)
	assert.Equal(t, 2, len(recipes))
	assert.Equal(t, 1, recipes[0].Id)
	assert.Equal(t, 2, recipes[1].Id)
	assert.Equal(t, "SELECT * FROM recipe ", adapter.query)
	assert.Equal(t, 0, len(adapter.bindVars))
	assert.True(t, adapter.destTypeOk)
}

func TestFind_WithWhere(t *testing.T) {
	adapter := &RecipeTestAdapter{}
	rf := Finder{}
	rf.Where("field", "=", "value")
	recipes, err := rf.Find(adapter)

	assert.Nil(t, err)
	assert.NotNil(t, recipes)
	assert.Equal(t, 2, len(recipes))
	assert.Equal(t, "SELECT * FROM recipe WHERE field = ?", adapter.query)
	assert.Equal(t, 1, len(adapter.bindVars))
	assert.Equal(t, "value", adapter.bindVars[0])
	assert.True(t, adapter.destTypeOk)
}

func TestFind_WithWhereAndOr(t *testing.T) {
	adapter := &RecipeTestAdapter{}
	rf := Finder{}
	rf.Where("field", "=", "value").AndWhere("field2", "=", "value2").OrWhere("field3", "=", "value3")
	recipes, err := rf.Find(adapter)

	assert.Nil(t, err)
	assert.NotNil(t, recipes)
	assert.Equal(t, 2, len(recipes))
	assert.Equal(t, "SELECT * FROM recipe WHERE field = ? AND field2 = ? OR field3 = ?", adapter.query)
	assert.Equal(t, 3, len(adapter.bindVars))
	assert.Equal(t, "value", adapter.bindVars[0])
	assert.Equal(t, "value2", adapter.bindVars[1])
	assert.Equal(t, "value3", adapter.bindVars[2])
	assert.True(t, adapter.destTypeOk)
}
