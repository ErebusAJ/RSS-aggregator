package main

import (
	"time"

	"github.com/ErebusAJ/rssagg/internal/database"
	"github.com/google/uuid"
)


// A structure and a function to follow json naming convention for user
type Users struct{
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAit time.Time `json:"updated_at"`
}

func databaseUserToUser(user database.User) Users{
	return Users{
		ID: user.ID,
		Name: user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAit: user.UpdatedAt,
	}
}


// Handling naming conventions for feeds table
type Feed struct{
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title string `json:"title"`
	URL string `json:"url"`
	UserID uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(feed database.Feed) Feed{
	return Feed{
		ID: feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Title: feed.Title,
		URL: feed.Url,
		UserID: feed.UserID,
	}
}


// A struct for msg
type Message struct{
	Text string `json:"message"`
}