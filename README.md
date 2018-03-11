# Meal Planner

Meal planning web application written in Golang.

## Setup

### Migrations

Migrations are written with Postgresql targeted.

The migrations are generic in format so should run with any migration runner that can read from filesystem.

#### Migrate

To use the Migrate CLI with Postgres first install Migrate.

```
$ go get -u -d github.com/mattes/migrate/cli github.com/lib/pq
$ go build -tags 'postgres' -o /usr/local/bin/migrate github.com/mattes/migrate/cli
```

Run the migrations (using an example postgres database string)

```
migrate -database postgres://localhost:5432/meals -source file://migrations/ up

```

## Tests

Tests are not unit tests and are not exhaustive. They are designed to run end to end; sending a HTTP request
and then testing the HTTP responses and what's in the database. The tests are designed to give a decent level of confidence
that a change will not break the system but things like error handling in all circumstances are not tested. The general principle
behind this is to try to be dogmatic and actually create something whilst having a good idea of what a change is going to do without
manual testing.