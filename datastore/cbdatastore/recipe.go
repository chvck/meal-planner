package cbdatastore

import (
	"fmt"

	"github.com/chvck/gocb"
	"github.com/chvck/meal-planner/model"
)

type recipe struct {
	*model.Recipe
	Type string `json:"type,omitempty"`
}

// One retrieves a single model.Recipe by id
func (ds CBDataStore) Recipe(id, userID string) (*model.Recipe, error) {
	r := recipe{}
	_, err := ds.bucket.Get(id, &r)
	if err != nil {
		return nil, err
	}

	if r.UserID != userID {
		return nil, fmt.Errorf("")
	}

	return r.Recipe, nil
}

// AllWithLimit retrieves x recipes starting from an offset
func (ds CBDataStore) Recipes(limit, offset int, userID string) ([]model.Recipe, error) {
	var recipes []model.Recipe
	query := gocb.NewN1qlQuery(`SELECT id, name, instructions, description, yield, prep_time, cook_time, user_id, ingredients
		FROM meals
		WHERE type = "recipe" AND user_id = $1
		ORDER BY id
		LIMIT $2 OFFSET $3;`)

	results, err := ds.bucket.ExecuteN1qlQuery(query, [3]interface{}{userID, limit, offset})
	if err != nil {
		return nil, err
	}

	r := model.Recipe{}
	for results.Next(&r) {
		recipes = append(recipes, r)
	}

	if err = results.Close(); err != nil {
		return nil, err
	}

	if len(recipes) == 0 {
		return recipes, nil
	}

	return recipes, nil
}

// Create creates the specific recipe
func (ds CBDataStore) RecipeCreate(modelRecipe model.Recipe, userID string) (*model.Recipe, error) {
	key := fmt.Sprintf("recipe::%s::%s", userID, modelRecipe.Name)
	newR := new(recipe)
	newR.Recipe = &modelRecipe
	newR.Type = "recipe"

	_, err := ds.bucket.Insert(key, newR, 0)
	if err != nil {
		return nil, err
	}

	return newR.Recipe, nil
}

// Update updates the specific recipe
func (ds CBDataStore) RecipeUpdate(modelRecipe model.Recipe, id, userID string) error {
	_, err := ds.bucket.Replace(modelRecipe.ID, modelRecipe, 0, 0)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the specific recipe
func (ds CBDataStore) RecipeDelete(id, userID string) error {
	frag, err := ds.bucket.LookupIn(id).Get("user_id").Execute()
	if err != nil {
		return err
	}

	var recipeUserID string
	err = frag.Content("user_id", &recipeUserID)
	if err != nil {
		return err
	}

	if recipeUserID != userID {
		return fmt.Errorf("recipe does not exist")
	}

	_, err = ds.bucket.Remove(id, 0)
	if err != nil {
		return err
	}

	return nil
}
