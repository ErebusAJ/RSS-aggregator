package main

import (
	"net/http"
)


func testingHandler(w http.ResponseWriter, r *http.Request){
	// jsonHandler(w, 200, struct{Message string}{Message: "JSON handler works!!"})
	errorHandler(w, 400, "cannot execute")
}
