package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSONResponse parses the data into JSON and writes it into the response
func JSONResponse(response interface{}, w http.ResponseWriter) {
	js, err := json.Marshal(response)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
