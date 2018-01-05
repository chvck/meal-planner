package recipe

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/chvck/meal-planner/db"
	"errors"
	"database/sql"
	"reflect"
)

type RecipeTestAdapter struct {
	queries   []string
	bindVars  [][]interface{}
	Row       []interface{}
	QueryRows [][]interface{}
	AllRows [][][]interface{}
	i int
}

type mockRows struct {
	Rows [][]interface{}
	row  []interface{}
	i    int
}

func (mr *mockRows) Next() bool {
	if mr.i < len(mr.Rows) {
		mr.row = mr.Rows[mr.i]
		mr.i++
		return true
	}

	return false
}

func (mr *mockRows) Scan(dest ...interface{}) error {
	if len(dest) != len(mr.row) {
		return errors.New("incorrect number of arguments supplied to Scan")
	}

	for i, col := range mr.row {
		rv := reflect.ValueOf(dest[i])
		rv.Elem().Set(reflect.ValueOf(col))
	}

	return nil
}

func (mr mockRows) Close() error {
	return nil
}

func (r RecipeTestAdapter) Initialize(dbType string, connectionString string) error {
	return nil
}

func (r *RecipeTestAdapter) Query(baseQuery string, bindVars ...interface{}) (db.Rows, error) {
	r.queries = append(r.queries, baseQuery)
	r.bindVars = append(r.bindVars, bindVars)
	r.QueryRows = r.AllRows[r.i]
	rows := &mockRows{Rows: r.QueryRows}

	r.i++
	return rows, nil
}

func (r *RecipeTestAdapter) QueryOne(baseQuery string, bindVars ...interface{}) db.Row {
	//r.query = baseQuery
	//r.bindVars = bindVars
	//r.destTypeOk = false
	//recipe, ok := dest.(*Recipe)
	//if ok {
	//	r.destTypeOk = true
	//
	//	recipe.Id = 1
	//	recipe.Name = "test recipe"
	//}

	return nil
}

func (r RecipeTestAdapter) Exec(baseExec string, bindVars ...interface{}) (int, error) {
	return 0, nil
}

//func TestOne(t *testing.T) {
//	adapter := &RecipeTestAdapter{}
//	recipe, err := One(adapter, 1)
//
//	assert.Nil(t, err)
//	assert.NotNil(t, recipe)
//	assert.Equal(t, 1, recipe.Id)
//	assert.Equal(t, "SELECT * FROM recipe WHERE id = ?", adapter.query)
//	assert.Equal(t, 1, len(adapter.bindVars))
//	assert.Equal(t, 1, adapter.bindVars[0])
//}

func TestAll(t *testing.T) {
	var row1 []interface{}
	var row2 []interface{}
	row1 = append(row1, 1, "name1", "desc1", sql.NullString{String: "inst1"}, sql.NullInt64{Int64: 1}, sql.NullInt64{Int64: 1}, sql.NullInt64{Int64: 1})
	row2 = append(row2, 1, 1, "ing1", "cup", 2)

	var query1Rows [][]interface{}
	query1Rows = append(query1Rows, row1)

	var query2Rows [][]interface{}
	query2Rows = append(query2Rows, row2)

	var allRows [][][]interface{}
	allRows = append(allRows, query1Rows, query2Rows)

	adapter := &RecipeTestAdapter{AllRows: allRows}
	recipesPtr, err := All(adapter)

	assert.Nil(t, err)
	assert.NotNil(t, recipesPtr)
	recipes := *recipesPtr
	assert.Equal(t, 1, len(recipes))
	assert.Equal(t, 1, recipes[0].Id)
	assert.Equal(t, 1, len(recipes[0].Ingredients))
	assert.Equal(t, 1, recipes[0].Ingredients[0].Id)
	assert.Equal(t, 2, len(adapter.queries))
}

//
//func TestAllWithLimit(t *testing.T) {
//	adapter := &RecipeTestAdapter{}
//	recipes, err := AllWithLimit(adapter, 10, 5)
//
//	assert.Nil(t, err)
//	assert.NotNil(t, recipes)
//	assert.Equal(t, 2, len(recipes))
//	assert.Equal(t, 1, recipes[0].Id)
//	assert.Equal(t, 2, recipes[1].Id)
//	assert.Equal(t, "SELECT * FROM recipe LIMIT 10 OFFSET 5", adapter.query)
//	assert.Equal(t, 0, len(adapter.bindVars))
//	assert.True(t, adapter.destTypeOk)
//}
//
//func TestFind(t *testing.T) {
//	adapter := &RecipeTestAdapter{}
//	rf := Finder{}
//	recipes, err := rf.Find(adapter)
//
//	assert.Nil(t, err)
//	assert.NotNil(t, recipes)
//	assert.Equal(t, 2, len(recipes))
//	assert.Equal(t, 1, recipes[0].Id)
//	assert.Equal(t, 2, recipes[1].Id)
//	assert.Equal(t, "SELECT * FROM recipe ", adapter.query)
//	assert.Equal(t, 0, len(adapter.bindVars))
//	assert.True(t, adapter.destTypeOk)
//}
//
//func TestFind_WithWhere(t *testing.T) {
//	adapter := &RecipeTestAdapter{}
//	rf := Finder{}
//	rf.Where("field", "=", "value")
//	recipes, err := rf.Find(adapter)
//
//	assert.Nil(t, err)
//	assert.NotNil(t, recipes)
//	assert.Equal(t, 2, len(recipes))
//	assert.Equal(t, "SELECT * FROM recipe WHERE field = ?", adapter.query)
//	assert.Equal(t, 1, len(adapter.bindVars))
//	assert.Equal(t, "value", adapter.bindVars[0])
//	assert.True(t, adapter.destTypeOk)
//}
//
//func TestFind_WithWhereAndOr(t *testing.T) {
//	adapter := &RecipeTestAdapter{}
//	rf := Finder{}
//	rf.Where("field", "=", "value").AndWhere("field2", "=", "value2").OrWhere("field3", "=", "value3")
//	recipes, err := rf.Find(adapter)
//
//	assert.Nil(t, err)
//	assert.NotNil(t, recipes)
//	assert.Equal(t, 2, len(recipes))
//	assert.Equal(t, "SELECT * FROM recipe WHERE field = ? AND field2 = ? OR field3 = ?", adapter.query)
//	assert.Equal(t, 3, len(adapter.bindVars))
//	assert.Equal(t, "value", adapter.bindVars[0])
//	assert.Equal(t, "value2", adapter.bindVars[1])
//	assert.Equal(t, "value3", adapter.bindVars[2])
//	assert.True(t, adapter.destTypeOk)
//}
