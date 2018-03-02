package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/chvck/meal-planner/model"
	"github.com/chvck/meal-planner/service"
	"github.com/dgrijalva/jwt-go"
)

// UserController is the interface for a controller than handles user endpoints
type UserController interface {
	UserLogin(w http.ResponseWriter, r *http.Request)
	UserCreate(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	service service.UserService
}

// NewUserController creates a new user controller
func NewUserController(service service.UserService) UserController {
	return &userController{service: service}
}

type userWithPassword struct {
	model.User
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
func (uc userController) UserLogin(w http.ResponseWriter, r *http.Request) {
	var creds loginCredentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
		return
	}

	u := uc.service.ValidatePassword(creds.Username, []byte(creds.Password))
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
func (uc userController) UserCreate(w http.ResponseWriter, r *http.Request) {
	var u userWithPassword

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid user", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &u); err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid user", http.StatusBadRequest)
		return
	}

	created, err := uc.service.Create(u.User, []byte(u.Password))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	JSONResponse(created, w)
}

func createToken(user *model.User) (*jwtToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  user.Username,
		"email":     user.Email,
		"id":        user.ID,
		"lastLogin": user.LastLogin,
		"level":     model.LevelUser,
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}
	return &jwtToken{Token: tokenString}, err
}
