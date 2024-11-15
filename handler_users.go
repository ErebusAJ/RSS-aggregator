package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ErebusAJ/rssagg/internal/database"
	"github.com/google/uuid"
)


// Create User Handler method of apiConfig 
// Takes arguments ResponseWriter
// Decodes Parmaeter name from r.body
// Calls sqlc CreateUser to create new user in database
// Returns JSON if OK else error
func(cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		errorHandler(w, http.StatusInternalServerError, fmt.Sprintf("couldn't decode parameter: %v", err))
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		errorHandler(w, http.StatusInternalServerError, fmt.Sprintf("error creating user: %v", err))
		log.Printf("error creating user: %v", err)
		return
	}

	jsonHandler(w, http.StatusCreated, databaseUserToUser(user))
}


// GetUsers function
// Retreives all the users form DB
func(cfg *apiConfig) handlerGetUsers(w http.ResponseWriter, r *http.Request){
	users, err := cfg.DB.GetUsers(r.Context())
	if err != nil {
		errorHandler(w, http.StatusInternalServerError, "couldn't retrieve users")
		log.Printf("couldn't retrieve users: %v", err)
		return
	}

	for _, user := range users{
		jsonHandler(w, http.StatusOK, databaseUserToUser(user))
	}
}


// GetUserByApiKey function
// Called through middlewareAuth
// Retreives user by apikey
func(cfg *apiConfig) handlerGetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User){
	user, err := cfg.DB.GetUserByApiKey(r.Context(), user.ApiKey)
	if err != nil{
		errorHandler(w, http.StatusNotFound, "couldn't find user")
		return
	}

	jsonHandler(w, http.StatusFound, user)
}


// DeleteUser handler
// Deletes a user if he's authenticated to do it
// Called by a middleware auth
// Uses apiKey to delete user
func(cfg *apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request, user database.User){
	err := cfg.DB.DeleteUser(r.Context(), user.ApiKey)
	if err != nil {
		errorHandler(w, http.StatusInternalServerError, "couldn't delete user")
		log.Printf("couldn't delete user: %v", err)
		return
	}

	jsonHandler(w, http.StatusOK, Message{Text: "successfully deleted user"})
}