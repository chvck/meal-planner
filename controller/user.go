package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/chvck/meal-planner/model"
	"github.com/dgrijalva/jwt-go"
)

type userWithPassword struct {
	model.User
	Password string `json:"password,omitzero"`
}

type loginCredentials struct {
	Username string `json:"username,omitzero"`
	Password string `json:"password,omitzero"`
}

type jwtToken struct {
	Token string `json:"token,omitzero"`
}

// UserLogin is the HTTP handler for logging as user into the system
func (sc StandardController) UserLogin(w http.ResponseWriter, r *http.Request) {
	var creds loginCredentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	u := sc.ds.UserValidatePassword(creds.Username, []byte(creds.Password))
	if u == nil {
		err := errors.New("invalid user credentials provided")
		JSONResponseWithCode(NewJSONError(err), w, http.StatusUnauthorized)
		return
	}

	t, err := createToken(u, sc.authKey)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	JSONResponse(t, w)
}

// UserCreate is the HTTP handler for creating a user
func (sc StandardController) UserCreate(w http.ResponseWriter, r *http.Request) {
	var u userWithPassword

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &u); err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	if errs := u.Validate(); len(errs) > 0 {
		JSONResponseWithCode(NewJSONErrors(errs), w, http.StatusBadRequest)
		return
	}

	if err := model.ValidatePassword(u.Password); err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	created, err := sc.ds.UserCreate(u.User, []byte(u.Password))
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	JSONResponseWithCode(created, w, 201)
}

func createToken(user *model.User, key string) (*jwtToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  user.Username,
		"email":     user.Email,
		"id":        user.ID,
		"lastLogin": user.LastLogin,
		"level":     model.LevelUser,
	})
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return nil, err
	}
	return &jwtToken{Token: tokenString}, err
}
