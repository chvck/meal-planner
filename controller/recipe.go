package controller

import (
	"net/http"
	"fmt"
	"github.com/chvck/meal-planner/context"
	"github.com/chvck/meal-planner/model/recipe"
	"encoding/json"
	"log"
)

func RecipeIndex(w http.ResponseWriter, r *http.Request) {
	db := context.Database()
	recipes, err := recipe.All(db)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not retrieve recipes", http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(recipes)
	if err != nil {
		http.Error(w, "Could not retrieve recipes", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(js))
}
