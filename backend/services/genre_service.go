package services

import (
	"context"
	"fmt"

	"book-recommendation-system/backend/models"
	"book-recommendation-system/backend/repositories"
)

// GenreService defines the interface for genre-related business logic.
type GenreService interface {
	GetGenreByID(ctx context.Context, id string) (*models.Genre, error)
	GetAllGenres(ctx context.Context) ([]models.Genre, error)
	CreateGenre(ctx context.Context, genre *models.Genre) error
	UpdateGenre(ctx context.Context, genre *models.Genre) error
	DeleteGenre(ctx context.Context, id string) error
}

// genreService implements GenreService.
type genreService struct {
	repo repositories.GenreRepository
}

// NewGenreService creates a new GenreService.
func NewGenreService(repo repositories.GenreRepository) GenreService {
	return &genreService{repo: repo}
}

// GetGenreByID retrieves a genre by its ID using the repository.
func (s *genreService) GetGenreByID(ctx context.Context, id string) (*models.Genre, error) {
	genre, err := s.repo.GetGenreByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get genre by ID: %w", err)
	}
	return genre, nil
}

// GetAllGenres retrieves all genres using the repository.
func (s *genreService) GetAllGenres(ctx context.Context) ([]models.Genre, error) {
	genres, err := s.repo.GetAllGenres(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get all genres: %w", err)
	}
	return genres, nil
}

// CreateGenre creates a new genre using the repository.
func (s *genreService) CreateGenre(ctx context.Context, genre *models.Genre) error {
	// Add any business logic/validation before creating a genre
	err := s.repo.CreateGenre(ctx, genre)
	if err != nil {
		return fmt.Errorf("service: failed to create genre: %w", err)
	}
	return nil
}

// UpdateGenre updates an existing genre using the repository.
func (s *genreService) UpdateGenre(ctx context.Context, genre *models.Genre) error {
	// Add any business logic/validation before updating a genre
	err := s.repo.UpdateGenre(ctx, genre)
	if err != nil {
		return fmt.Errorf("service: failed to update genre: %w", err)
	}
	return nil
}

// DeleteGenre deletes a genre by its ID using the repository.
func (s *genreService) DeleteGenre(ctx context.Context, id string) error {
	// Add any business logic/validation before deleting a genre
	err := s.repo.DeleteGenre(ctx, id)
	if err != nil {
		return fmt.Errorf("service: failed to delete genre: %w", err)
	}
	return nil
}
