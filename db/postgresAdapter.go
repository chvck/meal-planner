package db

import "github.com/jmoiron/sqlx"

// PostgresAdapter is a data store adapter that persists to the Postgres database
type PostgresAdapter struct {
	db *sqlx.DB
}

// Initialize sets up the connection to the database
func (p *PostgresAdapter) Initialize(connectionString string) error {
	if db, err := sqlx.Connect("postgres", connectionString); err != nil {
		return err
	} else {
		p.db = db
		return nil
	}
}

// Query performs the specified query and populates the array with retrieved data
func (p *PostgresAdapter) Query(dest interface{}, baseQuery string, bindVars ...interface{}) error {
	query := p.db.Rebind(baseQuery)
	if err := p.db.Select(&dest, query, bindVars); err != nil {
		return err
	} else {
		return nil
	}
}

// Query performs the specified query and populates the interface with retrieved data, will only retrieve a single row
func (p *PostgresAdapter) QueryOne(dest interface{}, baseQuery string, bindVars ...interface{}) error {
	query := p.db.Rebind(baseQuery)
	if err := p.db.Get(&dest, query, bindVars); err != nil {
		return err
	} else {
		return nil
	}
}

// Exec executes a statement
func (p *PostgresAdapter) Exec(baseExec string, bindVars ...interface{}) (int, error) {
	if result, err := p.db.Exec(baseExec, bindVars); err != nil {
		return -1, err
	} else {
		if rows, err := result.RowsAffected(); err != nil {
			return -1, err
		} else {
			return int(rows), nil
		}
	}
}
