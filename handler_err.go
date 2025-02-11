package main

import(
	"net/http")

func handlerErr(w http.ResponseWriter, r *http.Request) {
	// respondWithJSON(w, http.StatusOK, map[string]string{"status": "Working! OK"})
	respondWithError(w,400,"Something went wrong")
}