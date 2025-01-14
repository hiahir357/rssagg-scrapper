package main

import (
	"fmt"
	"net/http"

	"github.com/hiahir357/rssagg/internal/auth"
	"github.com/hiahir357/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(&r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %s", err))
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 404, fmt.Sprintf("Couldn't get user: %s", err))
		}
		handler(w, r, user)
	}
}
