package main

import (
	"fmt"
	"net/http"

	"github.com/amanraj8848/rssScrapper/internal/auth"
	"github.com/amanraj8848/rssScrapper/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKEY(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, 400, "User not found")
			return
		}

		handler(w, r, user)
	}
}