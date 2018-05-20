package sqldatastore

import (
	"github.com/chvck/meal-planner/model"
	"database/sql"
	"errors"
)

type planner struct {
	ID          int    `db:"id"`
	UserID      int    `db:"user_id"`
	When        int    `db:"when"`
	For         string `db:"for"`
	RecipeNames []model.RecipeName
}

func plannerFromModelPlanner(modelPlan model.Planner) *planner {
	p := planner{}
	p.ID = modelPlan.ID
	p.UserID = modelPlan.UserID
	p.When = modelPlan.When
	p.For = modelPlan.For

	return &p
}

func (p planner) toModelPlanner() *model.Planner {
	modelPlan := model.Planner{}
	modelPlan.ID = p.ID
	modelPlan.UserID = p.UserID
	modelPlan.When = p.When
	modelPlan.For = p.For
	modelPlan.RecipeNames = p.RecipeNames

	return &modelPlan
}

func (ds *SQLDataStore) PlannerWithRecipeNames(when int, mealtime string, userID int) (*model.Planner, error) {
	query := `SELECT p.id, p.when, p.for, p.user_id, r.id, r.name FROM planner p
		JOIN planner_to_recipe pr ON p.id = pr.planner_id
		JOIN recipe r ON pr.recipe_id = r.id WHERE p.when = ? AND p.for = ? and p.user_id = ?;`

	rows, err := ds.DB.Queryx(ds.DB.Rebind(query), when, mealtime, userID)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	p := planner{}
	defer rows.Close()
	for rows.Next() {
		var pWithRN struct {
			planner
			model.RecipeName
		}
		{
		}
		err = rows.Scan(&pWithRN.planner.ID, &pWithRN.planner.When, &pWithRN.planner.For, &pWithRN.planner.UserID,
			&pWithRN.RecipeName.ID, &pWithRN.RecipeName.Name)
		if err != nil {
			return nil, err
		}

		if p.ID == 0 {
			p = pWithRN.planner
		}

		p.RecipeNames = append(p.RecipeNames, pWithRN.RecipeName)
	}

	return p.toModelPlanner(), nil
}

func (ds *SQLDataStore) PlannersWithRecipeNames(start, end, userID int) ([]model.Planner, error) {
	query := `SELECT p.id, p.when, p.for, p.user_id, r.id, r.name FROM planner p
		JOIN planner_to_recipe pr ON p.id = pr.planner_id
		JOIN recipe r ON pr.recipe_id = r.id WHERE p.when BETWEEN ? AND ? and user_id = ?
		ORDER BY p.id;`

	rows, err := ds.DB.Queryx(ds.DB.Rebind(query), start, end, userID)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var planners []model.Planner
	p := planner{}
	defer rows.Close()
	for rows.Next() {
		var pWithRN struct {
			planner
			model.RecipeName
		}
		{
		}
		err = rows.Scan(&pWithRN.planner.ID, &pWithRN.planner.When, &pWithRN.planner.For, &pWithRN.planner.UserID,
			&pWithRN.RecipeName.ID, &pWithRN.RecipeName.Name)
		if err != nil {
			return nil, err
		}

		if p.ID != pWithRN.planner.ID && p.ID != 0 {
			planners = append(planners, *p.toModelPlanner())
			p = pWithRN.planner
		}

		p.RecipeNames = append(p.RecipeNames, pWithRN.RecipeName)
	}

	if p.ID != 0 {
		planners = append(planners, *p.toModelPlanner())
	}

	return planners, nil
}

func (ds *SQLDataStore) PlannerCreate(when int, mealtime string, userID int) (int, error) {
	query := `INSERT INTO "planner" (when, mealtime, user_id) VALUES (?, ?, ?) RETURNING id;`

	var pID int
	err := ds.DB.Get(&pID, ds.DB.Rebind(query), when, mealtime, userID)
	if err != nil {
		return 0, err
	}

	return pID, nil
}

func (ds *SQLDataStore) PlannerAddRecipe(plannerID, recipeID, userID int) error {
	var checkID int
	err := ds.DB.Get(&checkID, ds.DB.Rebind("SELECT id FROM planner WHERE id = ? AND user_id = ?"), plannerID, userID)
	if err == sql.ErrNoRows {
		return errors.New("planner could not be found")
	}

	err = ds.DB.Get(&checkID, ds.DB.Rebind("SELECT id FROM recipe WHERE id = ? AND user_id = ?"), recipeID, userID)
	if err == sql.ErrNoRows {
		return errors.New("recipe could not be found")
	}

	query := `INSERT INTO "planner_to_recipe" ("planner_id", "recipe_id") VALUES (?, ?) RETURNING id;`
	_, err = ds.DB.Exec(ds.DB.Rebind(query), plannerID, recipeID)
	if err != nil {
		return err
	}

	return nil
}

func (ds *SQLDataStore) PlannerRemoveRecipe(plannerID, recipeID, userID int) error {
	var checkID int
	err := ds.DB.Get(&checkID, ds.DB.Rebind("SELECT id FROM planner WHERE id = ? AND user_id = ?"), plannerID, userID)
	if err == sql.ErrNoRows {
		return errors.New("planner could not be found")
	}

	err = ds.DB.Get(&checkID, ds.DB.Rebind("SELECT id FROM recipe WHERE id = ? AND user_id = ?"), recipeID, userID)
	if err == sql.ErrNoRows {
		return errors.New("recipe could not be found")
	}

	query := `DELETE pr FROM "planner_to_recipe" WHERE planner_id = ? AND recipe_id = ?;`

	if _, err := ds.DB.Exec(ds.DB.Rebind(query), plannerID, userID, recipeID); err != nil {
		return err
	}

	return nil
}
