package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSONError is a wrapper for error/errors that can be parsed into JSON
type JSONError struct {
	Errors []string `json:"errors"`
}

// NewJSONError creates a new JSONError from a single error
func NewJSONError(err error) JSONError {
	errs := []error{err}
	return NewJSONErrors(errs)
}

// NewJSONErrors creates a new JSONError from a collection of errors
func NewJSONErrors(errs []error) JSONError {
	jsonErr := JSONError{}
	for _, err := range errs {
		jsonErr.Errors = append(jsonErr.Errors, err.Error())
	}

	return jsonErr
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
