package model

// Planner is the model for the planner table
type Planner struct {
	ID      int      `db:"id" json:"id"`
	UserID  int      `db:"user_id" json:"user_id"`
	When    int      `db:"when" json:"when"`
	For     string   `db:"for" json:"for"`
	Menus   []Menu   `json:"menus"`
	Recipes []Recipe `json:"recipes"`
}
