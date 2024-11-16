package main

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Description string    `xml:"description"`
		Link        string    `xml:"link"`
		Language 	string	  `xml:"language"`
		Item		[]RSSItem `xml:"item"`	

	} `xml:"channel"`
}

type RSSItem struct{
	Title		string `xml:"title"`	 
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PublishDate string `xml:"pubDate"`
}

func urlToFeed(url string) (RSSFeed, error){
	httpClient := http.Client{
		Timeout: 15 * time.Second,
	}

	response, err := httpClient.Get(url)
	if err != nil{
		log.Printf("error couldn't get url %v", err)
		return RSSFeed{}, err
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil{
		log.Printf("error coudn't read response body: %v", err)
		return RSSFeed{}, err
	}

	rssFeed := RSSFeed{}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil{
		log.Printf("error unmarshling xml: %v", err)
		return RSSFeed{}, err
	}

	return rssFeed, nil
}
