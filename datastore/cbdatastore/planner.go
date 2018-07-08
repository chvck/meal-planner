package cbdatastore

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/chvck/gocb"
	"github.com/chvck/meal-planner/model"
)

type planner struct {
	*model.Planner
	Type string `json:"type,omitempty"`
}

func (ds *CBDataStore) PlannerWithRecipeNames(id, userID string) (*model.Planner, error) {
	if !checkModelID(id, userID) {
		return nil, errors.New("no planner found")
	}

	p := planner{}
	_, err := ds.bucket.Get(id, &p)
	if err != nil {
		return nil, err
	}

	return p.Planner, nil
}

func (ds *CBDataStore) PlannersWithRecipeNames(start, end int, userID string) ([]model.Planner, error) {
	query := gocb.NewN1qlQuery(`SELECT id, when, for, userID, recipes
		FROM meals
		WHERE type = "planner" AND user_id = $1
		AND when BETWEEN $2 AND $3
		ORDER BY id;`)

	results, err := ds.bucket.ExecuteN1qlQuery(query, [3]interface{}{userID, start, end})
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var planners []model.Planner
	p := planner{}
	for results.Next(&p) {
		planners = append(planners, *p.Planner)
	}

	if err = results.Close(); err != nil {
		return nil, err
	}

	return planners, nil
}

func (ds *CBDataStore) PlannerCreate(when int, mealtime, userID string) (*model.Planner, error) {
	key := fmt.Sprintf("planner::%s::%d::%s", userID, when, mealtime)
	newP := new(planner)
	newP.Type = "planner"
	newP.When = when
	newP.For = mealtime

	_, err := ds.bucket.Insert(key, newP, 0)
	if err != nil {
		return nil, err
	}

	return newP.Planner, nil
}

func (ds *CBDataStore) PlannerAddRecipe(plannerID, recipeID, userID string) error {
	if !checkModelID(plannerID, userID) || !checkModelID(recipeID, userID) {
		return errors.New("no planner found")
	}

	frag, err := ds.bucket.LookupIn(recipeID).Get("name").Execute()
	if err != nil {
		return err
	}
	var name string
	err = frag.Content("name", &name)
	if err != nil {
		return err
	}
	ds.bucket.MutateIn(plannerID, 0, 0).ArrayAppend("recipes", model.RecipeName{
		ID:   recipeID,
		Name: name,
	}, true)

	return nil
}

func (ds *CBDataStore) PlannerRemoveRecipe(plannerID, recipeID, userID string) error {
	if !checkModelID(plannerID, userID) || !checkModelID(recipeID, userID) {
		return errors.New("no planner found")
	}

	var p planner
	cas, err := ds.bucket.Get(plannerID, &p)
	if err != nil {
		return err
	}

	newNames := make([]model.RecipeName, len(p.RecipeNames)-1)
	for i, rn := range p.RecipeNames {
		if rn.ID != recipeID {
			newNames[i] = rn
		}
	}

	p.RecipeNames = newNames
	_, err = ds.bucket.Replace(plannerID, p, cas, 0)
	if err != nil {
		return err
	}
	return nil
}
