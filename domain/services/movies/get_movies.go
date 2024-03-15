package movies

import (
	"github.com/j-monteiro/movies-api/domain/models"
	"github.com/j-monteiro/movies-api/storage"
)

type GetMoviesService struct {
	Repository storage.MoviesRepositoryInterface
}

type GetMoviesServiceParams struct {
	ID    int
	Title string
}

func (s *GetMoviesService) Perform(search_params GetMoviesServiceParams) []models.Movie {
	return s.Repository.FilterMoviesBy(storage.MovieFilters{
		ID:    search_params.ID,
		Title: search_params.Title,
	})
}
