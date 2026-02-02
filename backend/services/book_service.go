package services

import (
	"context"
	"fmt"

	"book-recommendation-system/backend/models"
	"book-recommendation-system/backend/repositories"
)

// BookService defines the interface for book-related business logic.
type BookService interface {
	GetBookByID(ctx context.Context, id string) (*models.Book, error)
	GetAllBooks(ctx context.Context) ([]models.Book, error)
	CreateBook(ctx context.Context, book *models.Book) error
	UpdateBook(ctx context.Context, book *models.Book) error
	DeleteBook(ctx context.Context, id string) error
}

// bookService implements BookService.
type bookService struct {
	repo repositories.BookRepository
}

// NewBookService creates a new BookService.
func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{repo: repo}
}

// GetBookByID retrieves a book by its ID using the repository.
func (s *bookService) GetBookByID(ctx context.Context, id string) (*models.Book, error) {
	book, err := s.repo.GetBookByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get book by ID: %w", err)
	}
	return book, nil
}

// GetAllBooks retrieves all books using the repository.
func (s *bookService) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	books, err := s.repo.GetAllBooks(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get all books: %w", err)
	}
	return books, nil
}

// CreateBook creates a new book using the repository.
func (s *bookService) CreateBook(ctx context.Context, book *models.Book) error {
	// Add any business logic/validation before creating a book
	if book.ID == "" {
		// Example: Generate a new ID if not provided
		// book.ID = uuid.New().String()
	}

	err := s.repo.CreateBook(ctx, book)
	if err != nil {
		return fmt.Errorf("service: failed to create book: %w", err)
	}
	return nil
}

// UpdateBook updates an existing book using the repository.
func (s *bookService) UpdateBook(ctx context.Context, book *models.Book) error {
	// Add any business logic/validation before updating a book
	err := s.repo.UpdateBook(ctx, book)
	if err != nil {
		return fmt.Errorf("service: failed to update book: %w", err)
	}
	return nil
}

// DeleteBook deletes a book by its ID using the repository.
func (s *bookService) DeleteBook(ctx context.Context, id string) error {
	// Add any business logic/validation before deleting a book
	err := s.repo.DeleteBook(ctx, id)
	if err != nil {
		return fmt.Errorf("service: failed to delete book: %w", err)
	}
	return nil
}
