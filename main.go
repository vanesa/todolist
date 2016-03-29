package main // package main is required for a standalone executable.

import (
    	"log"
    	"net/http"
)

func main() {
	// Here we register three routes mapping URL paths to handlers.
	router := NewRouter()
	
    	log.Fatal(http.ListenAndServe(":8080", router))
}