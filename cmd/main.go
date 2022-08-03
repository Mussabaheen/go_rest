package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mussabaheen/go_rest/pkg/articles"
)

func main() {
	// Define httpRouter as a main router for articles endpoint
	httpRouter := mux.NewRouter().StrictSlash(true)

	// Passing httpRouter to the articles service
	articles.HandleRequests(httpRouter)

	log.Print("http Server running on localhost:8080")
	defer log.Print(http.ListenAndServe(":8080", httpRouter))

}
