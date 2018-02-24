package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/chvck/meal-planner/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type exception struct {
	Message string `json:"message"`
}

func routes(handler *Handler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/recipe/", validateMiddleware(handler.RecipeIndex)).Methods("GET")
	router.HandleFunc("/recipe/{id}", validateMiddleware(handler.RecipeByID)).Methods("GET")
	router.HandleFunc("/recipe/", validateMiddleware(handler.RecipeCreate)).Methods("POST")
	router.HandleFunc("/recipe/{id}", validateMiddleware(handler.RecipeUpdate)).Methods("POST")
	router.HandleFunc("/recipe/{id}", validateMiddleware(handler.RecipeDelete)).Methods("DELETE")

	router.HandleFunc("/menu/{id}", validateMiddleware(handler.MenuByID)).Methods("GET")

	router.HandleFunc("/login/", handler.UserLogin).Methods("POST")
	router.HandleFunc("/register/", handler.UserCreate).Methods("POST")

	return router
}

func validateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader == "" {
			json.NewEncoder(w).Encode(exception{Message: "An authorization header is required"})
			return
		}

		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) == 2 {
			token, err := jwt.Parse(bearerToken[1], parseToken)
			if err != nil {
				json.NewEncoder(w).Encode(exception{Message: err.Error()})
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				u := model.User{ID: int(claims["id"].(float64)), Username: claims["username"].(string)}
				context.Set(req, "user", u)
				next(w, req)
			} else {
				json.NewEncoder(w).Encode(exception{Message: "Invalid authorization token"})
			}
		}
	})
}

func parseToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("there was an error")
	}
	return []byte("secret"), nil
}
