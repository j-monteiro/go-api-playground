package main

import (
	"log"
	"net/http"

	"github.com/j-monteiro/movies-api/api"
	"github.com/j-monteiro/movies-api/api/middlewares"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := api.NewRouter()

	chainedMiddleware := middlewares.Chain(middlewares.SetupCors, middlewares.JSONContentType)
	http.ListenAndServe(":8080", chainedMiddleware(router))
}
