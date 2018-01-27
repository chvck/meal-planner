package model

import (
	"github.com/chvck/meal-planner/db"
)

// IDataStoreAdapter is the interface for structs that persist data
type IDataStoreAdapter interface {
	Initialize(dbType string, connectionString string) error
	Query(baseQuery string, bindVars ...interface{}) (db.Rows, error)
	QueryOne(baseQuery string, bindVars ...interface{}) db.Row
	Exec(baseExec string, bindVars ...interface{}) (int, error)
	NewTransaction() (db.Transaction, error)
	DBType() string
}
