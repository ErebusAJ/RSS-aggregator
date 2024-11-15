package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ErebusAJ/rssagg/internal/database"
	"github.com/ErebusAJ/rssagg/internal/decoder"
	"github.com/google/uuid"
)

// CreateFeed handler
// Creates a feed referenced to a authenticated user
// Decodes request body parameters name, url
// Creates feed using the request's params
func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	}

	var params parameters
	decoder.Decode(r.Body, &params)

	data, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title:     params.Title,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		errorHandler(w, http.StatusInternalServerError, "couldn't create feed")
		log.Printf("error creating feed: %v", err)
		return
	}

	jsonHandler(w, http.StatusCreated, databaseFeedToFeed(data))
}

// GetFeed handler
// Uses get request to retrieve all the feeds in db
func (cfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		errorHandler(w, http.StatusNotFound, "error retrieving feeds")
		log.Printf("error retrieving posts: %v", err)
		return
	}

	for _, item := range feeds {
		jsonHandler(w, http.StatusFound, databaseFeedToFeed(item))
	}
}

// DeleteFeed handler
// Deletes feed based on feed_id, and authenticated user
// If another user trie to delete feed throws error
func (cfg *apiConfig) handlerDeleteFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	var params parameters
	decoder.Decode(r.Body, &params)

	data, err := cfg.DB.DeleteFeed(r.Context(), database.DeleteFeedParams{
		ID:     params.FeedID,
		UserID: user.ID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			errorHandler(w, http.StatusForbidden, "cannot delete feed of other users")
			log.Printf("error deleting other user feed: %v", err)
		} else {
			errorHandler(w, http.StatusInternalServerError, "couldn't delete feed")
			log.Printf("error deleting feed: %v", err)
		}
		return
	}

	jsonHandler(w, http.StatusOK, Message{Text: fmt.Sprintf("successfull deleted feed: %v", data)})
}
