package model

// IDataStoreAdapter is the interface for structs that persist data
type IDataStoreAdapter interface {
	Initialize(connectionString string) error
	Query(dest interface{}, baseQuery string, bindVars ...interface{}) error
	QueryOne(dest interface{}, baseQuery string, bindVars ...interface{}) error
	Exec(baseExec string, bindVars ...interface{}) (int, error)
}
