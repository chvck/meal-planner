package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
	"strings"
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

type SqlxTransaction struct {
	db *sqlx.DB
	tx *sql.Tx
}

// Exec executes a statement as a part of this Transaction
func (tx SqlxTransaction) Exec(baseExec string, bindVars ...interface{}) (int, error) {
	e := tx.db.Rebind(baseExec)
	if result, err := tx.tx.Exec(e, bindVars...); err != nil {
		return -1, err
	} else {
		if rows, err := result.RowsAffected(); err != nil {
			return -1, err
		} else {
			return int(rows), nil
		}
	}
}

// Query performs the specified query and populates the interface with retrieved data, will only retrieve a single row
func (tx SqlxTransaction) QueryOne(baseQuery string, bindVars ...interface{}) Row {
	query := tx.db.Rebind(baseQuery)
	return SqlxRow{tx.tx.QueryRow(query, bindVars...)}
}

// Commit the current Transaction
func (tx SqlxTransaction) Commit() error {
	return tx.tx.Commit()
}

// Rollback the current Transaction
func (tx SqlxTransaction) Rollback() error {
	return tx.tx.Rollback()
}

// Initialize the adapter with a pre-connected database - primarily for testing
func (p *SqlxAdapter) InitializeWithDb(db *sqlx.DB) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf("database connection does not work %v", err.Error())
	}
	p.db = db

	return nil
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
	if strings.Contains(baseQuery, "?") {
		baseQuery = p.db.Rebind(baseQuery)
	}
	if rows, err := p.db.Query(baseQuery, bindVars...); err != nil {
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

// NewTransaction returns a new database Transaction
func (p SqlxAdapter) NewTransaction() (Transaction, error) {
	if tx, err := p.db.Begin(); err != nil {
		return nil, err
	} else {
		return &SqlxTransaction{db: p.db, tx: tx}, err
	}
}
