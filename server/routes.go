package server
//
//import (
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"strings"
//
//	"github.com/dgrijalva/jwt-go"
//	"github.com/gorilla/context"
//	"github.com/gorilla/mux"
//)
//
//type exception struct {
//	Message string `json:"message"`
//}
//
//var authKey string
//
//func routes(handler *Handler, key string) *mux.Router {
//	authKey = key
//	router := mux.NewRouter()
//
//	router.HandleFunc("/recipe/", validateMiddleware(handler.RecipeIndex, LevelUser)).Methods("GET")
//	router.HandleFunc("/recipe/{id}", validateMiddleware(handler.RecipeByID, LevelUser)).Methods("GET")
//	router.HandleFunc("/recipe/", validateMiddleware(handler.RecipeCreate, LevelUser)).Methods("POST")
//	router.HandleFunc("/recipe/{id}", validateMiddleware(handler.RecipeUpdate, LevelUser)).Methods("PUT")
//	router.HandleFunc("/recipe/{id}", validateMiddleware(handler.RecipeDelete, LevelUser)).Methods("DELETE")
//
//	// check the env and only add these in test in future
//	router.HandleFunc("/testuser/", validateMiddleware(test, LevelUser)).Methods("GET")
//	router.HandleFunc("/testadmin/", validateMiddleware(test, LevelAdmin)).Methods("GET")
//
//	router.HandleFunc("/login/", handler.UserLogin).Methods("POST")
//	router.HandleFunc("/register/", handler.UserCreate).Methods("POST")
//
//	return router
//}
//
//func test(w http.ResponseWriter, req *http.Request) {
//
//}
//
//func validateMiddleware(next http.HandlerFunc, reqLevel float64) http.HandlerFunc {
//	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
//		authorizationHeader := req.Header.Get("authorization")
//		if authorizationHeader == "" {
//			w.WriteHeader(http.StatusUnauthorized)
//			json.NewEncoder(w).Encode(exception{Message: "An authorization header is required"})
//			return
//		}
//
//		bearerToken := strings.Split(authorizationHeader, " ")
//		if len(bearerToken) == 2 {
//			token, err := jwt.Parse(bearerToken[1], parseToken)
//			if err != nil {
//				w.WriteHeader(http.StatusUnauthorized)
//				json.NewEncoder(w).Encode(exception{Message: err.Error()})
//				return
//			}
//
//			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//				if claims["level"].(float64) < reqLevel {
//					w.WriteHeader(http.StatusUnauthorized)
//					json.NewEncoder(w).Encode(exception{Message: "Not authorized"})
//				}
//				u := User{ID: claims["id"].(string), Username: claims["username"].(string)}
//				context.Set(req, "user", u)
//				next(w, req)
//			} else {
//				w.WriteHeader(http.StatusUnauthorized)
//				json.NewEncoder(w).Encode(exception{Message: "Invalid authorization token"})
//			}
//		} else {
//			w.WriteHeader(http.StatusUnauthorized)
//			json.NewEncoder(w).Encode(exception{Message: "Invalid authorization token"})
//		}
//	})
//}
//
//func parseToken(token *jwt.Token) (interface{}, error) {
//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//		return nil, fmt.Errorf("there was an error")
//	}
//	return []byte(authKey), nil
//}

