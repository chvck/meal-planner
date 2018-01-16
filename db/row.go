package db

// Rows is the representation of a set of database rows
type Rows interface {
	Next() bool
	Scan(...interface{}) error
	Close() error
	Err() error
}

// Row is the representation of a single database row
type Row interface {
	Scan(...interface{}) error
}
