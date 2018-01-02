package db

type Rows interface {
	Next() bool
	Scan(...interface{}) error
}

type Row interface {
	Scan(...interface{}) error
}
