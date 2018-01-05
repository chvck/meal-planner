package db

type Rows interface {
	Next() bool
	Scan(...interface{}) error
	Close() error
}

type Row interface {
	Scan(...interface{}) error
}
