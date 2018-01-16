package store

import "github.com/chvck/meal-planner/model"

var database model.IDataStoreAdapter

// StoreDb stores the database adapter
func StoreDb(db model.IDataStoreAdapter) {
	database = db
}

// Database gets the database adapter
func Database() model.IDataStoreAdapter {
	return database
}
