package storage

import (
	"github.com/j-monteiro/movies-api/domain/models"
)

type MoviesRepositoryInterface interface {
	GetAllMovies() []models.Movie
	GetMovieByID(id int) *models.Movie
	FilterMoviesBy(search_params MovieFilters) []models.Movie
}

type MoviesRepository struct {
}

type MovieFilters struct {
	ID    int
	Title string
}

func NewMoviesRepository() *MoviesRepository {
	return &MoviesRepository{}
}

// define a movies constant
var movies_storage = []models.Movie{
	{ID: 1, Title: "The Shawshank Redemption", Description: "Two imprisoned"},
	{ID: 2, Title: "The Godfather", Description: "The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son."},
	{ID: 3, Title: "The Dark Knight", Description: "When the menace known as the Joker emerges from his mysterious past, he wreaks havoc and chaos on the people of Gotham."},
}

func (r *MoviesRepository) GetAllMovies() []models.Movie {
	return movies_storage
}

func (r *MoviesRepository) GetMovieByID(id int) *models.Movie {
	for _, movie := range movies_storage {
		if movie.ID == id {
			return &movie
		}
	}

	return nil
}

func (r *MoviesRepository) FilterMoviesBy(search_params MovieFilters) []models.Movie {
	if search_params.ID == 0 && search_params.Title == "" {
		return movies_storage
	}

	if search_params.ID != 0 {
		return []models.Movie{*r.GetMovieByID(search_params.ID)}
	}

	movies := []models.Movie{}
	for _, movie := range movies_storage {
		if movie.Title == search_params.Title {
			movies = append(movies, movie)
		}
	}

	return movies
}
