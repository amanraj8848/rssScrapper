package main

import(
	"net/http")

func handerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "Working! OK"})
}