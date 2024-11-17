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
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseUserToUser(user database.User) Users{
	return Users{
		ID: user.ID,
		Name: user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func databaseUsersToUsers(users []database.User) []Users{
	user := []Users{}
	for _, item := range users{
		user = append(user, databaseUserToUser(item))
	}
	return user
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

func databaseFeedsToFeeds(dbfeed []database.Feed) []Feed{
	feeds := []Feed{}
	for _, dbFeed := range dbfeed{
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}

	return feeds
}

// A struct for msg
type Message struct{
	Text string `json:"message"`
}


// Struct and function for feed_follows table
type FeedFollows struct{
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID uuid.UUID `json:"user_id"`
	FeedID uuid.UUID `json:"feed_d"`
}

func databaseFeedFollowsToFeedFollows(feedFollows database.FeedFollow) FeedFollows{
	return FeedFollows{
		ID: feedFollows.ID,
		CreatedAt: feedFollows.CreatedAt,
		UpdatedAt: feedFollows.UpdatedAt,
		UserID: feedFollows.UserID,
		FeedID: feedFollows.FeedID,
	}
}

func databaseFeedsFollowsToFeedsFollows(feedsFollows []database.FeedFollow) []FeedFollows{
	feedFollows := []FeedFollows{}
	for _, item := range feedsFollows{
		feedFollows = append(feedFollows, databaseFeedFollowsToFeedFollows(item))
	}

	return feedFollows
}