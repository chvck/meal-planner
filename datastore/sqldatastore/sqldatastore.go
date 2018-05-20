package sqldatastore

import "github.com/jmoiron/sqlx"

// SQLDataStore is a type of DataStore backing onto a SQL Database
type SQLDataStore struct{
	*sqlx.DB
}

// NewSQLDataStore creates and returns a new SQLDataStore
func NewSQLDataStore(dbType, connString string) (*SQLDataStore, error) {
	db, err := sqlx.Connect(dbType, connString)
	if err != nil {
		return nil, err
	}

	dataStore := SQLDataStore{db}

	return &dataStore, nil
}