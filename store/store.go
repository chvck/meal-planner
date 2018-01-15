package store

import "github.com/chvck/meal-planner/model"

var DB model.IDataStoreAdapter

func StoreDb(db model.IDataStoreAdapter) {
	DB = db
}

func Database() model.IDataStoreAdapter {
	return DB
}
