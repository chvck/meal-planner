package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/chvck/meal-planner/model/ingredient"
	"github.com/chvck/meal-planner/store"
)

func IngredientIndex(w http.ResponseWriter, r *http.Request) {
	db := store.Database()
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
