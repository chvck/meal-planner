package server

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"

	"github.com/chvck/meal-planner/proto/model"
	"gopkg.in/couchbase/gocb.v1"
)

// DataStore is used to access data from the underlying store
type DataStore interface {
	User(id string) (*model.User, error)
	Users(limit, offset int) ([]model.User, error)
	UserCreate(u model.User, password []byte) (*model.User, error)
	UserValidatePassword(username string, pw []byte) *model.User
	Recipe(id, userID string) (*model.Recipe, error)
	Recipes(limit, offset int, userID string) ([]model.Recipe, error)
	RecipeCreate(r model.Recipe, userID string) (*model.Recipe, error)
	RecipeUpdate(r model.Recipe, id, userID string) error
	RecipeDelete(id, userID string) error
	PlannerWithRecipeNames(id, userID string) (*model.Planner, error)
	PlannersWithRecipeNames(start, end int, userID string) ([]model.Planner, error)
	PlannerCreate(date int64, mealtime model.Planner_Mealtime, userID string) (*model.Planner, error)
	PlannerAddRecipe(plannerID, recipeID, userID string) error
	PlannerRemoveRecipe(plannerID, recipeID, userID string) error
}

// ConnectionConfig holds configuration details for connecting to a DataStore
type ConnectionConfig struct {
	server   string
	port     int
	username string
	password string
}

// CBDataStore is a type of DataStore backing onto a Couchbase Database
type CBDataStore struct {
	cluster *gocb.Cluster
	bucket  *gocb.Bucket
}

// NewDataStore creates and returns a new DataStore
func NewDataStore(host string, port uint, bucketName, username, password string) (DataStore, error) {
	connString := fmt.Sprintf("http://%s:%d", host, port)
	cluster, err := gocb.Connect(connString)
	if err != nil {
		return nil, err
	}

	err = cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	bucket, err := cluster.OpenBucket(bucketName, "")
	if err != nil {
		return nil, err
	}

	return &CBDataStore{
		cluster: cluster,
		bucket:  bucket,
	}, nil
}

func checkModelID(id, userID string) bool {
	splitID := strings.Split(id, "::")
	user := splitID[1]

	return user == userID
}

type recipe struct {
	*model.Recipe
	Type string `json:"type,omitempty"`
}

// Recipe retrieves a single Recipe by id
func (ds CBDataStore) Recipe(id, userID string) (*model.Recipe, error) {
	r := recipe{}
	_, err := ds.bucket.Get(id, &r)
	if err != nil {
		return nil, err
	}

	if r.UserId != userID {
		return nil, fmt.Errorf("")
	}

	return r.Recipe, nil
}

// Recipes retrieves x recipes starting from an offset
func (ds CBDataStore) Recipes(limit, offset int, userID string) ([]model.Recipe, error) {
	query := gocb.NewN1qlQuery(
		fmt.Sprintf("SELECT id, name, instructions, description, yield, prep_time, cook_time, user_id, ingredients "+
			"FROM `%s` "+
			"WHERE `type` = \"recipe\" AND user_id = $1 "+
			"ORDER BY id "+
			"LIMIT $2 OFFSET $3;", ds.bucket.Name()),
	)

	results, err := ds.bucket.ExecuteN1qlQuery(query, []interface{}{userID, limit, offset})
	if err != nil {
		return nil, err
	}

	r := model.Recipe{}
	var recipes []model.Recipe
	for results.Next(&r) {
		recipes = append(recipes, r)
	}

	if err = results.Close(); err != nil {
		return nil, err
	}

	if len(recipes) == 0 {
		return []model.Recipe{}, nil
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
	_, err := ds.bucket.Replace(modelRecipe.Id, modelRecipe, 0, 0)
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
	if err == gocb.ErrNoResults {
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

func (ds *CBDataStore) PlannerCreate(date int64, mealtime model.Planner_Mealtime, userID string) (*model.Planner, error) {
	key := fmt.Sprintf("planner::%s::%d::%s", userID, date, mealtime)
	newP := new(planner)
	newP.Type = "planner"
	newP.Date = date
	newP.Mealtime = mealtime

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

	_, err := ds.bucket.MutateIn(plannerID, 0, 0).ArrayAppend("recipes", recipeID, true).Execute()

	return err
}

func (ds *CBDataStore) PlannerRemoveRecipe(plannerID, recipeID, userID string) error {
	if !checkModelID(plannerID, userID) || !checkModelID(recipeID, userID) {
		return errors.New("no planner found")
	}

	frag, err := ds.bucket.LookupIn(plannerID).Get("recipe_ids").Execute()
	if err != nil {
		return err
	}

	var recipeIDs []string
	err = frag.ContentByIndex(0, &recipeIDs)
	if err != nil {
		return err
	}

	var newRecipeIDs []string
	for _, id := range recipeIDs {
		if id != recipeID {
			newRecipeIDs = append(newRecipeIDs, id)
		}
	}

	_, err = ds.bucket.MutateIn(plannerID, frag.Cas(), 0).Replace("recipe_ids", newRecipeIDs).Execute()
	return err
}

type user struct {
	model.User
	Type     string `json:"type,omitempty"`
	Password string `json:"password,omitempty"`
}

func (ds *CBDataStore) User(id string) (*model.User, error) {
	u := model.User{}
	_, err := ds.bucket.Get(id, &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (ds *CBDataStore) Users(limit, offset int) ([]model.User, error) {
	var users []model.User
	query := gocb.NewN1qlQuery(`SELECT id, username, email, created_at, updated_at, last_login
		FROM meals
		where type = "user"
		ORDER BY id
		LIMIT $1 OFFSET $2;`)
	results, err := ds.bucket.ExecuteN1qlQuery(query, [2]int{limit, offset})
	if err != nil {
		return nil, err
	}

	u := model.User{}
	for results.Next(&u) {
		users = append(users, u)
	}

	if err = results.Close(); err != nil {
		return nil, err
	}

	return users, nil
}

func (ds *CBDataStore) UserCreate(u model.User, password []byte) (*model.User, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("user::%s", u.Username)
	newU := user{}
	newU.Username = u.Username
	newU.Id = key
	newU.Email = u.Email

	now := time.Now().Unix()
	newU.CreatedAt = now
	newU.UpdatedAt = now
	newU.Password = string(hash)
	newU.Type = "user"

	_, err = ds.bucket.Insert(key, newU, 0)
	if err != nil {
		return nil, err
	}

	return &newU.User, nil
}

func (ds *CBDataStore) UserValidatePassword(username string, pw []byte) *model.User {
	var u user
	_, err := ds.bucket.Get(fmt.Sprintf("user::%s", username), &u)
	if err != nil {
		return nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), pw)
	if err != nil {
		return nil
	}

	return &u.User
}

