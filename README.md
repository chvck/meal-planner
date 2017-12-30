# Meal Planner

Meal planning web application written in Golang.

## Setup

### Migrations

Migrations are written with Postgresql targeted but with support for any RDBMS in mind.

The migrations are generic in format so should run with any migration runner that can read from filesystem.

#### Migrate

To use the Migrate CLI with Postgres first install Migrate 

```
$ go get -u -d github.com/mattes/migrate/cli github.com/lib/pq
$ go build -tags 'postgres' -o /usr/local/bin/migrate github.com/mattes/migrate/cli
```

Run the migrations (using an example postgres database string)

```
migrate -database postgres://localhost:5432/meals -source file://migrations/ up

```

