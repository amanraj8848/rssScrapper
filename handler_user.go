package main

import (
	"encoding/json"
	"net/http"
	"time"
	"fmt"
	"github.com/google/uuid"
	"github.com/amanraj8848/rssScrapper/internal/database"
)
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// Read the request body into a json decoder
	decoder := json.NewDecoder(r.Body)
	// Define a struct to hold the expected JSON parameters
	type parameters struct {
		Name string `json:"name"`
	}
	// Decode the request body into the parameters struct
	params := parameters{}
	err := decoder.Decode(&params)
	// If there is an error, return an error response to the client
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}
func (apiCfg *apiConfig) handlerGetUserByAPIKey	(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, (user))
}
func (apiCfg *apiConfig) handlerGetPostsForUser	(w http.ResponseWriter, r *http.Request, user database.User) {
	post,err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err!= nil{
		respondWithError(w, 400, fmt.Sprintf("Couldn't get posts for user: %v", err))
		return
	}
	respondWithJSON(w, 200, databasePostsToPosts(post))
}