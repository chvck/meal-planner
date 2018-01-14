package db

type Transaction interface {
	Exec(baseExec string, bindVars ...interface{}) (int, error)
	QueryOne(baseQuery string, bindVars ...interface{}) Row
	Commit() error
	Rollback() error
}
