package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/j-monteiro/movies-api/domain/services/movies"
)

type Movie struct {
	Title    string
	Director string
	Year     int
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	service := movies.GetMoviesService{}
	movies := service.Perform(movies.GetMoviesServiceParams{})

	json.NewEncoder(w).Encode(movies)
}
