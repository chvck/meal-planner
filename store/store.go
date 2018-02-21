package store

import "github.com/chvck/meal-planner/db"

var database db.DataStoreAdapter

// StoreDb stores the database adapter
func StoreDb(db db.DataStoreAdapter) {
	database = db
}

// Database gets the database adapter
func Database() db.DataStoreAdapter {
	return database
}
