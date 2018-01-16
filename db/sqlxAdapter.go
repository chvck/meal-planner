package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"           // required by sqlx
	_ "github.com/mattn/go-sqlite3" // required by sqlx
)

// SqlxAdapter is a data store adapter that utilises sqlx
type SqlxAdapter struct {
	db *sqlx.DB
}

// SqlxRows is an implementation of Rows that utilises sqlx
type SqlxRows struct {
	*sql.Rows
}

// SqlxRow is an implementation of Row that utilises sqlx
type SqlxRow struct {
	*sql.Row
}

// SqlxTransaction is an implementation of Transation that utilises sqlx
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

// QueryOne performs the specified query and returns a single row
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

// InitializeWithDb initializes the adapter with a pre-connected database - primarily for testing
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

// Query performs the specified query and returns a set of rows
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

// QueryOne performs the specified query and returns a single row
func (p SqlxAdapter) QueryOne(baseQuery string, bindVars ...interface{}) Row {
	query := p.db.Rebind(baseQuery)
	return SqlxRow{p.db.QueryRow(query, bindVars...)}
}

// Exec executes a statement and returns number of affected rows
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
