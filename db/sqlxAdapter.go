package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"database/sql"
)

// SqlxAdapter is a data store adapter that persists to the Postgres database
type SqlxAdapter struct {
	db *sqlx.DB
}

type SqlxRows struct {
	*sql.Rows
}

type SqlxRow struct {
	*sql.Row
}

// Initialize sets up the connection to the database
func (p *SqlxAdapter) Initialize(dbType string, connectionString string) error {
	if db, err := sqlx.Connect(dbType, connectionString); err != nil {
		return err
	} else {
		p.db = db
		return nil
	}
}

// Query performs the specified query and populates the array with retrieved data
func (p SqlxAdapter) Query(baseQuery string, bindVars ...interface{}) (Rows, error) {
	query := p.db.Rebind(baseQuery)
	if rows, err := p.db.Query(query, bindVars...); err != nil {
		return nil, err
	} else {
		return SqlxRows{rows}, nil
	}
}

// Query performs the specified query and populates the interface with retrieved data, will only retrieve a single row
func (p SqlxAdapter) QueryOne(baseQuery string, bindVars ...interface{}) Row {
	query := p.db.Rebind(baseQuery)
	return SqlxRow{p.db.QueryRow(query, bindVars...)}
}

// Exec executes a statement
func (p SqlxAdapter) Exec(baseExec string, bindVars ...interface{}) (int, error) {
	e := p.db.Rebind(baseExec)
	if result, err := p.db.Exec(e, bindVars...); err != nil {
		return -1, err
	} else {
		if rows, err := result.RowsAffected(); err != nil {
			return -1, err
		} else {
			return int(rows), nil
		}
	}
}
