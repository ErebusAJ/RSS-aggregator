package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/ErebusAJ/rssagg/internal/database"
	"github.com/google/uuid"
)

func startScrapping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration){
	log.Printf("Scrapping on %v goroutines every %v seconds", concurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)
	// for loop that executes every time a value is sent accross ticker.C channel
	// it executes every timeBetweenRequest seconds
	for ;; <-ticker.C{
		feeds, err := db.GetNextFeedToFetchFrom(context.Background(), int32(concurrency))
		if err != nil{
			log.Printf("error coouldn't get next feed to fetch from: %v", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds{
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()

	}
}


func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed){
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil{
		log.Printf("error marking feed as fetched: %v", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil{
		log.Printf("error fetching feed: %v", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		// handling nullable description
		description := sql.NullString{} 
		if item.Description != ""{
			description.String = item.Description
			description.Valid = true
		}

		// handling published date
		publisher_at, err := time.Parse(time.RFC3339, item.PublishDate)
		if err != nil {
			log.Printf("error parsing date: %v", err)
			continue
		}


		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title : item.Title,
			Description: description,
			Link: item.Link,
			PublishedAt: publisher_at,
			FeedID: feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key"){
				continue
			}
			log.Printf("error saving post to db: %v", err)
		}
	}
	log.Printf("Feed %v collected %v posts", feed.Title, len(rssFeed.Channel.Item))
}
