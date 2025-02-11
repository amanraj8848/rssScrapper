package main

import(
	"encoding/json"
	"net/http"
	"log"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499{
		log.Println("Responding with 5XX error:", message)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponse{Error: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code) // sets the header to given code 
	w.Write(response)	// write the response
}
