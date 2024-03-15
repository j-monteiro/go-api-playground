package movies

import (
	"testing"

	"github.com/j-monteiro/movies-api/domain/models"
	"github.com/j-monteiro/movies-api/storage"
)

type MockMoviesRepository struct{}

func (m *MockMoviesRepository) GetAllMovies() []models.Movie {
	return []models.Movie{
		{ID: 1, Title: "Movie 1"},
		{ID: 2, Title: "Movie 2"},
	}
}

func (m *MockMoviesRepository) GetMovieByID(id int) *models.Movie {
	return &models.Movie{ID: 1, Title: "Movie 1"}
}

func (m *MockMoviesRepository) FilterMoviesBy(params storage.MovieFilters) []models.Movie {
	return []models.Movie{
		{ID: 1, Title: "Movie 1"},
		{ID: 2, Title: "Movie 2"},
	}
}

func TestGetMoviesAll(t *testing.T) {
	service := GetMoviesService{
		Repository: &MockMoviesRepository{},
	}

	movies := service.Perform(GetMoviesServiceParams{})
	if len(movies) != 2 {
		t.Error("Expected 2 movies, got", len(movies))
	}
	if movies[0].ID != 1 {
		t.Error("Expected movie ID 1, got", movies[0].ID)
	}
	if movies[1].ID != 2 {
		t.Error("Expected movie ID 2, got", movies[1].ID)
	}
}

func TestGetMovieByID(t *testing.T) {
	service := GetMoviesService{
		Repository: &MockMoviesRepository{},
	}

	movies := service.Perform(GetMoviesServiceParams{ID: 1})
	if len(movies) != 1 {
		t.Error("Expected 1 movie, got", len(movies))
	}
	if movies[0].ID != 1 {
		t.Error("Expected movie ID 1, got", movies[0].ID)
	}
}

func TestGetMoviesByTitle(t *testing.T) {
	service := GetMoviesService{
		Repository: &MockMoviesRepository{},
	}

	movies := service.Perform(GetMoviesServiceParams{Title: "Movie 1"})
	if len(movies) != 2 {
		t.Error("Expected 2 movies, got", len(movies))
	}
	if movies[0].ID != 1 {
		t.Error("Expected movie ID 1, got", movies[0].ID)
	}
	if movies[1].ID != 2 {
		t.Error("Expected movie ID 2, got", movies[1].ID)
	}
}

func TestGetMoviesByTitleAndID(t *testing.T) {
	service := GetMoviesService{
		Repository: &MockMoviesRepository{},
	}

	movies := service.Perform(GetMoviesServiceParams{ID: 1, Title: "Movie 1"})
	if len(movies) != 1 {
		t.Error("Expected 1 movie, got", len(movies))
	}
	if movies[0].ID != 1 {
		t.Error("Expected movie ID 1, got", movies[0].ID)
	}
}
