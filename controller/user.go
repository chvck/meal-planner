package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/chvck/meal-planner/model/user"
	"github.com/chvck/meal-planner/store"
	"github.com/dgrijalva/jwt-go"
)

type userWithPassword struct {
	user.User
	Password string `json:"password"`
}

type loginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type jwtToken struct {
	Token string `json:"token"`
}

// UserLogin is the HTTP handler for logging as user into the system
func UserLogin(w http.ResponseWriter, r *http.Request) {
	var creds loginCredentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
		return
	}

	db := store.Database()
	u := user.ValidatePassword(db, creds.Username, []byte(creds.Password))
	if u == nil {
		http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
		return
	}

	t, err := createToken(u)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	JSONResponse(t, w)
}

// UserCreate is the HTTP handler for creating a user
func UserCreate(w http.ResponseWriter, r *http.Request) {
	db := store.Database()
	var u userWithPassword

	if body, err := ioutil.ReadAll(r.Body); err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid user", http.StatusBadRequest)
		return
	} else {
		if err := json.Unmarshal(body, &u); err != nil {
			log.Println(err.Error())
			http.Error(w, "Invalid user", http.StatusBadRequest)
			return
		}
	}

	err := user.Create(db, u.User, []byte(u.Password))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}
}

func createToken(user *user.User) (*jwtToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  user.Username,
		"email":     user.Email,
		"id":        user.ID,
		"lastLogin": user.LastLogin,
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}
	return &jwtToken{Token: tokenString}, err
}
