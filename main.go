package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ErebusAJ/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Struct for DB queries
type apiConfig struct{
	DB *database.Queries
}

func main(){
	fmt.Println("\t Welcome To RSS Aggregator \t")


	//Load the godotenv
	godotenv.Load()
	
	portNo := os.Getenv("PORT")
	if portNo == ""{
		log.Fatal("error occurred retrieving portNo")
	}


	// Database Connection
	// Retrieving database url from .env
	// Establishing db connection using sql package
	// New database db
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == ""{
		log.Print("error retrieving database url")
	}
	
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Printf("error connecting to database: %v", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	//Scrapper
	go startScrapping(db, 10, time.Minute)


	// Creating Router for our server
	router := chi.NewRouter()

	// Setting up the chi-router configuration using the chi-cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))


	// Adding router paths
	router.Get("/test", testingHandler)


	// Database Routers
	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)

	//Users 
	v1Router.Get("/users", apiCfg.handlerGetUsers)
	v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handlerGetUserByApiKey))
	v1Router.Post("/user", apiCfg.handlerCreateUser)
	v1Router.Delete("/user", apiCfg.middlewareAuth(apiCfg.handlerDeleteUser))
	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetUserPosts))
	
	//Feeds
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	v1Router.Post("/feed", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Delete("/feed", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeed))

	//FeedFollows
	v1Router.Get("/feed-follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1Router.Post("/feed-follow", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollows))
	v1Router.Delete("/feed-follow", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollows))

	// Server configuration
	// Setting --> Handler, Address
	// Starting up ListenAndServer server
	log.Printf("Starting server on PORT: %v \n", portNo)

	server := &http.Server{
		Handler: router,
		Addr: ":"+portNo,
	}

	err = server.ListenAndServe()
	if err != nil{
		log.Print("error starting server")
	}
	
}