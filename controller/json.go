package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSONError is a wrapper for error/errors that can be parsed into JSON
type JSONError struct {
	Error  error   `json:"error"`
	Errors []error `json:"errors"`
}

// JSONResponse parses the data into JSON and writes it into the response
func JSONResponse(response interface{}, w http.ResponseWriter) {
	JSONResponseWithCode(response, w, 200)
}

// JSONResponseWithCode parses the data into JSON and writes it into the response
// returning response with status code supplied
func JSONResponseWithCode(response interface{}, w http.ResponseWriter, statusCode int) {
	js, err := json.Marshal(response)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
