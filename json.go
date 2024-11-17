package main

import (
	"encoding/json"
	"log"
	"net/http"
)


// JSON Handler to encode the incoming payload to JSON
// Takes arguments:
// --> Response Write : Constructs HTTP response
// --> Code : status code
// --> Payload : incoming data struct
// Return --> JSON encoded data
func jsonHandler(w http.ResponseWriter, code int, payload interface{}) {
    // Marshal the payload to JSON
    data, err := json.Marshal(payload)
    if err != nil {
        log.Printf("error encoding to JSON: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Set the Content-Type header
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)

    // Write the JSON response
    _, err = w.Write(data)
    if err != nil {
        log.Printf("error writing response: %v", err)
    }
}



// Error handler to return an error executing any reponse handler
func errorHandler(w http.ResponseWriter, code int, msg string){
	if code > 499{
		log.Printf("server side erro 5xx: %v", msg)
		return
	}

	type errResponse struct{
		Error string `json:"error"`
	}

	jsonHandler(w, code, errResponse{Error: msg})
}