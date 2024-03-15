package api

import (
	"net/http"

	"github.com/j-monteiro/movies-api/api/controllers"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", controllers.GetMovies)

	return router
}
