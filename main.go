package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)


func main(){
	fmt.Println("\t Welcome To Rss Aggregator \t")


	//Load the godotenv
	godotenv.Load()
	
	portNo := os.Getenv("PORT")
	if portNo == ""{
		log.Print("error occurred retrieving portNo")
	}


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


	// Server configuration
	// Setting --> Handler, Address
	// Starting up ListenAndServer server
	log.Printf("Starting server on PORT: %v \n", portNo)

	server := &http.Server{
		Handler: router,
		Addr: ":"+portNo,
	}

	err := server.ListenAndServe()
	if err != nil{
		log.Print("error starting server")
	}
	
}