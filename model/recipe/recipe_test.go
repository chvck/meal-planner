package recipe

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/chvck/meal-planner/db"
	"errors"
	"database/sql"
	"reflect"
	"gopkg.in/guregu/null.v3"
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

type mockRow struct {
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


func (mr *mockRow) Scan(dest ...interface{}) error {
	if len(dest) != len(mr.row) {
		return errors.New("incorrect number of arguments supplied to Scan")
	}

	for i, col := range mr.row {
		rv := reflect.ValueOf(dest[i])
		rv.Elem().Set(reflect.ValueOf(col))
	}

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
	r.queries = append(r.queries, baseQuery)
	r.bindVars = append(r.bindVars, bindVars)

	return &mockRow{row: r.Row}
}

func (r RecipeTestAdapter) Exec(baseExec string, bindVars ...interface{}) (int, error) {
	return 0, nil
}

func TestNewRecipe(t *testing.T) {
	r := NewRecipe()
	assert.Equal(t, -1, r.Id)
}

func TestOne(t *testing.T) {
	i1 := ingredientWithProps{
		Id:       1,
		Name:     "ing1",
		Measure:  null.StringFrom("meas1"),
		Quantity: 12,
	}

	var iCol []ingredientWithProps
	iCol = append(iCol, i1)

	r1 := recipe{
		Id:           1,
		Name:         "name1",
		Description:  null.String{sql.NullString{String: "desc1"}},
		Instructions: "inst1",
		Yield:        null.Int{sql.NullInt64{Int64: 1}},
		PrepTime:     null.Int{sql.NullInt64{Int64: 30}},
		CookTime:     null.Int{sql.NullInt64{Int64: 5}},
		Ingredients:  iCol,
	}

	var rRow1 []interface{}
	var iRow1 []interface{}
	rRow1 = append(rRow1, r1.Id, r1.Name, r1.Instructions, r1.Description, r1.Yield, r1.PrepTime, r1.CookTime)
	iRow1 = append(iRow1, r1.Id, i1.Id, i1.Name, i1.Measure, i1.Quantity)

	var query2Rows [][]interface{}
	query2Rows = append(query2Rows, iRow1)

	var allRows [][][]interface{}
	allRows = append(allRows, query2Rows)

	adapter := &RecipeTestAdapter{Row: rRow1, AllRows: allRows}
	recipePtr, err := One(adapter, 1)

	assert.Nil(t, err)
	assert.NotNil(t, recipePtr)
	assertRecipe(t, &r1, recipePtr)
	assert.Equal(t, 2, len(adapter.queries))
}

func TestAll(t *testing.T) {
	i1 := ingredientWithProps{
		Id:       1,
		Name:     "ing1",
		Measure:  null.StringFrom("meas1"),
		Quantity: 12,
	}

	var iCol []ingredientWithProps
	iCol = append(iCol, i1)

	r1 := recipe{
		Id:           1,
		Name:         "name1",
		Description:  null.String{sql.NullString{String: "desc1"}},
		Instructions: "inst1",
		Yield:        null.Int{sql.NullInt64{Int64: 1}},
		PrepTime:     null.Int{sql.NullInt64{Int64: 30}},
		CookTime:     null.Int{sql.NullInt64{Int64: 5}},
		Ingredients:  iCol,
	}
	r2 := recipe{
		Id:           2,
		Name:         "name2",
		Description:  null.String{sql.NullString{String: "desc2"}},
		Instructions: "inst2",
		Yield:        null.Int{sql.NullInt64{Int64: 1}},
		PrepTime:     null.Int{sql.NullInt64{Int64: 30}},
		CookTime:     null.Int{sql.NullInt64{Int64: 5}},
		Ingredients:  iCol,
	}

	var rRow1 []interface{}
	var rRow2 []interface{}
	var iRow1 []interface{}
	var iRow2 []interface{}
	rRow1 = append(rRow1, r1.Id, r1.Name, r1.Instructions, r1.Description, r1.Yield, r1.PrepTime, r1.CookTime)
	rRow2 = append(rRow2, r2.Id, r2.Name, r2.Instructions, r2.Description, r2.Yield, r2.PrepTime, r2.CookTime)
	iRow1 = append(iRow1, r1.Id, i1.Id, i1.Name, i1.Measure, i1.Quantity)
	iRow2 = append(iRow2, r2.Id, i1.Id, i1.Name, i1.Measure, i1.Quantity)

	var query1Rows [][]interface{}
	query1Rows = append(query1Rows, rRow1, rRow2)

	var query2Rows [][]interface{}
	query2Rows = append(query2Rows, iRow1, iRow2)

	var allRows [][][]interface{}
	allRows = append(allRows, query1Rows, query2Rows)

	adapter := &RecipeTestAdapter{AllRows: allRows}
	recipesPtr, err := All(adapter)

	assert.Nil(t, err)
	assert.NotNil(t, recipesPtr)
	recipes := *recipesPtr
	assert.Equal(t, 2, len(recipes))
	assertRecipe(t, &r1, &recipes[0])
	assertRecipe(t, &r2, &recipes[1])
	assert.Equal(t, 2, len(adapter.queries))
}
