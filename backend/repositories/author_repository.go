package repositories

import (
	"context"
	"fmt"

	"book-recommendation-system/backend/models"
	"github.com/jmoiron/sqlx"
)

// AuthorRepository defines the interface for author data operations.
type AuthorRepository interface {
	GetAuthorByID(ctx context.Context, id string) (*models.Author, error)
	GetAllAuthors(ctx context.Context) ([]models.Author, error)
	CreateAuthor(ctx context.Context, author *models.Author) error
	UpdateAuthor(ctx context.Context, author *models.Author) error
	DeleteAuthor(ctx context.Context, id string) error
}

// authorRepository implements AuthorRepository using sqlx.
type authorRepository struct {
	db *sqlx.DB
}

// NewAuthorRepository creates a new AuthorRepository.
func NewAuthorRepository(db *sqlx.DB) AuthorRepository {
	return &authorRepository{db: db}
}

// GetAuthorByID retrieves an author by its ID.
func (r *authorRepository) GetAuthorByID(ctx context.Context, id string) (*models.Author, error) {
	var author models.Author
	err := r.db.GetContext(ctx, &author, "SELECT id, name, biography FROM authors WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting author by ID: %w", err)
	}
	return &author, nil
}

// GetAllAuthors retrieves all authors.
func (r *authorRepository) GetAllAuthors(ctx context.Context) ([]models.Author, error) {
	var authors []models.Author
	err := r.db.SelectContext(ctx, &authors, "SELECT id, name, biography FROM authors")
	if err != nil {
		return nil, fmt.Errorf("error getting all authors: %w", err)
	}
	return authors, nil
}

// CreateAuthor creates a new author.
func (r *authorRepository) CreateAuthor(ctx context.Context, author *models.Author) error {
	query := `INSERT INTO authors (id, name, biography) VALUES (:id, :name, :biography)`
	_, err := r.db.NamedExecContext(ctx, query, author)
	if err != nil {
		return fmt.Errorf("error creating author: %w", err)
	}
	return nil
}

// UpdateAuthor updates an existing author.
func (r *authorRepository) UpdateAuthor(ctx context.Context, author *models.Author) error {
	query := `UPDATE authors SET name=:name, biography=:biography WHERE id=:id`
	_, err := r.db.NamedExecContext(ctx, query, author)
	if err != nil {
		return fmt.Errorf("error updating author: %w", err)
	}
	return nil
}

// DeleteAuthor deletes an author by its ID.
func (r *authorRepository) DeleteAuthor(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM authors WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("error deleting author: %w", err)
	}
	return nil
}
