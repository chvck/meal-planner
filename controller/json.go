package controller

import (
	"encoding/json"
	"net/http"
	"log"
)

func JsonResponse(response interface{}, w http.ResponseWriter) {
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
