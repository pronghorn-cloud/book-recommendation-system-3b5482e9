package repositories

import (
	"context"
	"fmt"

	"book-recommendation-system/backend/models"
	"github.com/jmoiron/sqlx"
)

// GenreRepository defines the interface for genre data operations.
type GenreRepository interface {
	GetGenreByID(ctx context.Context, id string) (*models.Genre, error)
	GetAllGenres(ctx context.Context) ([]models.Genre, error)
	CreateGenre(ctx context.Context, genre *models.Genre) error
	UpdateGenre(ctx context.Context, genre *models.Genre) error
	DeleteGenre(ctx context.Context, id string) error
}

// genreRepository implements GenreRepository using sqlx.
type genreRepository struct {
	db *sqlx.DB
}

// NewGenreRepository creates a new GenreRepository.
func NewGenreRepository(db *sqlx.DB) GenreRepository {
	return &genreRepository{db: db}
}

// GetGenreByID retrieves a genre by its ID.
func (r *genreRepository) GetGenreByID(ctx context.Context, id string) (*models.Genre, error) {
	var genre models.Genre
	err := r.db.GetContext(ctx, &genre, "SELECT id, name FROM genres WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting genre by ID: %w", err)
	}
	return &genre, nil
}

// GetAllGenres retrieves all genres.
func (r *genreRepository) GetAllGenres(ctx context.Context) ([]models.Genre, error) {
	var genres []models.Genre
	err := r.db.SelectContext(ctx, &genres, "SELECT id, name FROM genres")
	if err != nil {
		return nil, fmt.Errorf("error getting all genres: %w", err)
	}
	return genres, nil
}

// CreateGenre creates a new genre.
func (r *genreRepository) CreateGenre(ctx context.Context, genre *models.Genre) error {
	query := `INSERT INTO genres (id, name) VALUES (:id, :name)`
	_, err := r.db.NamedExecContext(ctx, query, genre)
	if err != nil {
		return fmt.Errorf("error creating genre: %w", err)
	}
	return nil
}

// UpdateGenre updates an existing genre.
func (r *genreRepository) UpdateGenre(ctx context.Context, genre *models.Genre) error {
	query := `UPDATE genres SET name=:name WHERE id=:id`
	_, err := r.db.NamedExecContext(ctx, query, genre)
	if err != nil {
		return fmt.Errorf("error updating genre: %w", err)
	}
	return nil
}

// DeleteGenre deletes a genre by its ID.
func (r *genreRepository) DeleteGenre(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM genres WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("error deleting genre: %w", err)
	}
	return nil
}
