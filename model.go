package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

// Ingredient is the model for the ingredient table
type Ingredient struct {
	ID       string          `json:"id,omitempty"`
	RecipeID string          `json:"recipeId,omitempty"`
	Name     string          `json:"name,omitempty"`
	Measure  string          `json:"measure,omitempty"`
	Quantity decimal.Decimal `json:"quantity,omitempty"`
}

// String is the string representation of an ingredient.Ingredient
func (i Ingredient) String() string {
	return fmt.Sprintf("%v %v %v", i.Quantity, i.Measure, i.Name)
}

// Planner is the model for the planner table
type Planner struct {
	ID          string       `json:"id,omitempty"`
	UserID      string       `json:"userID,omitempty"`
	When        int          `json:"when,omitempty"`
	For         string       `json:"for,omitempty"`
	RecipeNames []RecipeName `json:"recipes,omitempty"`
}

type RecipeName struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Validate checks that the planner is valid
func (p Planner) Validate() []error {
	var errs []error
	if p.When == 0 {
		errs = append(errs, errors.New("when cannot be empty"))
	}
	if p.For == "" {
		errs = append(errs, errors.New("for cannot be empty"))
	}

	return errs
}

// Recipe is the model for the recipe table
type Recipe struct {
	ID           string       `json:"id,omitempty"`
	UserID       string       `json:"userId,omitempty"`
	Name         string       `json:"name,omitempty"`
	Instructions string       `json:"instructions,omitempty"`
	Yield        int          `json:"yield,omitempty"`
	PrepTime     int          `json:"prepTime,omitempty"`
	CookTime     int          `json:"cookTime,omitempty"`
	Description  string       `json:"description,omitempty"`
	Ingredients  []Ingredient `json:"ingredients,omitempty"`
}

// Validate checks that the recipe is valid
func (r Recipe) Validate() []error {
	var errs []error
	if r.Name == "" {
		errs = append(errs, errors.New("name cannot be empty"))
	}
	if r.Instructions == "" {
		errs = append(errs, errors.New("instructions cannot be empty"))
	}

	return errs
}

// User is the model for the user table
type User struct {
	ID        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	CreatedAt int    `json:"createdAt,omitempty"`
	UpdatedAt int    `json:"updatedAt,omitempty"`
	LastLogin int    `json:"lastLogin,omitempty"`
}

// Levels are for user access levels
const (
	LevelUser  = 1.0
	LevelAdmin = 2.0
)

// Validate checks that the user is valid
func (u User) Validate() []error {
	var errs []error
	if u.Username == "" {
		errs = append(errs, errors.New("username cannot be empty"))
	}
	if u.Email == "" {
		errs = append(errs, errors.New("email cannot be empty"))
	}

	return errs
}

// ValidatePassword checks that the password is valid
func ValidatePassword(pwd string) error {
	if len(pwd) < 8 {
		return errors.New("password must be longer than 8 characters")
	}

	return nil
}

