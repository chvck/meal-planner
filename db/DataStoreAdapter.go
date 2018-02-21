package db

// DataStoreAdapter is the interface for structs that persist data
type DataStoreAdapter interface {
	Initialize(dbType string, connectionString string) error
	Query(baseQuery string, bindVars ...interface{}) (Rows, error)
	QueryOne(baseQuery string, bindVars ...interface{}) Row
	Exec(baseExec string, bindVars ...interface{}) (int, error)
	NewTransaction() (Transaction, error)
}
