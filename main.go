package main

import (
	"github.com/gorilla/handlers"
	"log"
	"net/http"

	"github.com/chuckbutler/gix/gifs"
)

func main() {
	router := gifs.NewRouter()

	// These lines are important in order to allow access from a front-end
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// launch server with CORS validations
	log.Fatal(http.ListenAndServe(":9000",
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
