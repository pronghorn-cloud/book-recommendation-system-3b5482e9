package services

import (
	"context"
	"fmt"

	"book-recommendation-system/backend/models"
	"book-recommendation-system/backend/repositories"
)

// AuthorService defines the interface for author-related business logic.
type AuthorService interface {
	GetAuthorByID(ctx context.Context, id string) (*models.Author, error)
	GetAllAuthors(ctx context.Context) ([]models.Author, error)
	CreateAuthor(ctx context.Context, author *models.Author) error
	UpdateAuthor(ctx context.Context, author *models.Author) error
	DeleteAuthor(ctx context.Context, id string) error
}

// authorService implements AuthorService.
type authorService struct {
	repo repositories.AuthorRepository
}

// NewAuthorService creates a new AuthorService.
func NewAuthorService(repo repositories.AuthorRepository) AuthorService {
	return &authorService{repo: repo}
}

// GetAuthorByID retrieves an author by its ID using the repository.
func (s *authorService) GetAuthorByID(ctx context.Context, id string) (*models.Author, error) {
	author, err := s.repo.GetAuthorByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get author by ID: %w", err)
	}
	return author, nil
}

// GetAllAuthors retrieves all authors using the repository.
func (s *authorService) GetAllAuthors(ctx context.Context) ([]models.Author, error) {
	authors, err := s.repo.GetAllAuthors(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get all authors: %w", err)
	}
	return authors, nil
}

// CreateAuthor creates a new author using the repository.
func (s *authorService) CreateAuthor(ctx context.Context, author *models.Author) error {
	// Add any business logic/validation before creating an author
	err := s.repo.CreateAuthor(ctx, author)
	if err != nil {
		return fmt.Errorf("service: failed to create author: %w", err)
	}
	return nil
}

// UpdateAuthor updates an existing author using the repository.
func (s *authorService) UpdateAuthor(ctx context.Context, author *models.Author) error {
	// Add any business logic/validation before updating an author
	err := s.repo.UpdateAuthor(ctx, author)
	if err != nil {
		return fmt.Errorf("service: failed to update author: %w", err)
	}
	return nil
}

// DeleteAuthor deletes an author by its ID using the repository.
func (s *authorService) DeleteAuthor(ctx context.Context, id string) error {
	// Add any business logic/validation before deleting an author
	err := s.repo.DeleteAuthor(ctx, id)
	if err != nil {
		return fmt.Errorf("service: failed to delete author: %w", err)
	}
	return nil
}
