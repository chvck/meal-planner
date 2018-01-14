package server

import (
	"github.com/gorilla/mux"
	"github.com/chvck/meal-planner/controller"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"net/http"
	"strings"
	"encoding/json"
)

type Exception struct {
	Message string `json:"message"`
}

func routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/recipe/", controller.RecipeIndex).Methods("GET")
	router.HandleFunc("/recipe/{id}", controller.RecipeById).Methods("GET")
	router.HandleFunc("/recipe/", controller.RecipeCreate).Methods("POST")

	router.HandleFunc("/ingredient/", controller.IngredientIndex).Methods("GET")

	router.HandleFunc("/login/", controller.UserLogin).Methods("POST")
	router.HandleFunc("/user/", ValidateMiddleware(controller.UserCreate)).Methods("POST")

	return router
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader == "" {
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
			return
		}

		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) == 2 {
			token, err := jwt.Parse(bearerToken[1], parseToken)
			if err != nil {
				json.NewEncoder(w).Encode(Exception{Message: err.Error()})
				return
			}

			if token.Valid {
				next(w, req)
			} else {
				json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
			}
		}
	})
}

func parseToken(token *jwt.Token) (interface{}, error){
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("there was an error")
	}
	return []byte("secret"), nil
}