package main

import (
	"log"
	"net/http"

	"github.com/ErebusAJ/rssagg/internal/auth"
	"github.com/ErebusAJ/rssagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func(cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		api, err := auth.GetApiKey(r.Header)
		if err != nil{
			errorHandler(w, http.StatusInternalServerError, "couldn't retrieve api key")
			log.Printf("couldn't retrieve api from header: %v", err)
			return
		}

		user, err := cfg.DB.GetUserByApiKey(r.Context(), api)
		if err != nil{
			errorHandler(w, http.StatusNotFound, "couldn't retreive user")
			log.Printf("couldn't retrieve user: %v", err)
			return
		}
		
		handler(w, r, user)
	}
}