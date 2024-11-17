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

// Create FeedFollows handler
func (cfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct{
		FeedID uuid.UUID `json:"feed_id"`
	}

	var params parameters
	decoder.Decode(r.Body, &params)

	data, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})
	if err != nil {
		errorHandler(w, http.StatusInternalServerError, "error following feed")
		log.Printf("error following feed: %v", err)
		return
	}
	
	jsonHandler(w, http.StatusCreated, databaseFeedFollowsToFeedFollows(data))

}


// Get user feed handler
func (cfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User){
	feed, err := cfg.DB.GetFeedFollow(r.Context(), user.ID)
	if err != nil{
		if err == sql.ErrNoRows{
			errorHandler(w, http.StatusOK, "user doesn't follow anyone")
			log.Printf("no rows in result set")
		}else{
			errorHandler(w, http.StatusInternalServerError, "error retrieving following feed")
			log.Printf("error retrieving feed follows: %v", err)
		}
		return
	}

	jsonHandler(w, http.StatusFound, databaseFeedsFollowsToFeedsFollows(feed))
}


// Delete feedfollow handler
func (cfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct{
		FeedFollowsID uuid.UUID `json:"feed_follows_id"`
	}

	var params parameters
	decoder.Decode(r.Body, &params)

	feed, err := cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: params.FeedFollowsID,
	})
	if err != nil{
		if err == sql.ErrNoRows{
			errorHandler(w, http.StatusInternalServerError, "cannot remove following feed")
			log.Printf("error couldn't delete feed follow of different user: %v", err)
		}else{
			errorHandler(w, http.StatusInternalServerError, "cannot remove following feed")
			log.Printf("erro couldn't delete feed follow: %v", err)
		}
		return
	}

	jsonHandler(w, http.StatusOK, Message{Text: fmt.Sprintf("successfully deleted feed follow %v", feed)})
}