package controller

import (
	"log"
	"encoding/json"
	"fmt"
	"github.com/chvck/meal-planner/context"
	"net/http"
	"github.com/chvck/meal-planner/model/ingredient"
)

func IngredientIndex(w http.ResponseWriter, r *http.Request) {
	db := context.Database()
	ingredients, err := ingredient.All(db)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not retrieve ingredients", http.StatusNotFound)
		return
	}

	js, err := json.Marshal(ingredients)
	if err != nil {
		http.Error(w, "Could not retrieve ingredients", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(js))
}