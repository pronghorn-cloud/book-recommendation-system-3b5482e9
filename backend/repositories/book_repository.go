package repositories

import (
	"context"
	"fmt"

	"book-recommendation-system/backend/models"
	"github.com/jmoiron/sqlx"
)

// BookRepository defines the interface for book data operations.
type BookRepository interface {
	GetBookByID(ctx context.Context, id string) (*models.Book, error)
	GetAllBooks(ctx context.Context) ([]models.Book, error)
	CreateBook(ctx context.Context, book *models.Book) error
	UpdateBook(ctx context.Context, book *models.Book) error
	DeleteBook(ctx context.Context, id string) error
}

// bookRepository implements BookRepository using sqlx.
type bookRepository struct {
	db *sqlx.DB
}

// NewBookRepository creates a new BookRepository.
func NewBookRepository(db *sqlx.DB) BookRepository {
	return &bookRepository{db: db}
}

// GetBookByID retrieves a book by its ID.
func (r *bookRepository) GetBookByID(ctx context.Context, id string) (*models.Book, error) {
	var book models.Book
	err := r.db.GetContext(ctx, &book, "SELECT id, title, author, isbn, description, cover_image_url, genre, publication_year FROM books WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting book by ID: %w", err)
	}
	return &book, nil
}

// GetAllBooks retrieves all books.
func (r *bookRepository) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	err := r.db.SelectContext(ctx, &books, "SELECT id, title, author, isbn, description, cover_image_url, genre, publication_year FROM books")
	if err != nil {
		return nil, fmt.Errorf("error getting all books: %w", err)
	}
	return books, nil
}

// CreateBook creates a new book.
func (r *bookRepository) CreateBook(ctx context.Context, book *models.Book) error {
	query := `INSERT INTO books (id, title, author, isbn, description, cover_image_url, genre, publication_year) VALUES (:id, :title, :author, :isbn, :description, :cover_image_url, :genre, :publication_year)`
	_, err := r.db.NamedExecContext(ctx, query, book)
	if err != nil {
		return fmt.Errorf("error creating book: %w", err)
	}
	return nil
}

// UpdateBook updates an existing book.
func (r *bookRepository) UpdateBook(ctx context.Context, book *models.Book) error {
	query := `UPDATE books SET title=:title, author=:author, isbn=:isbn, description=:description, cover_image_url=:cover_image_url, genre=:genre, publication_year=:publication_year WHERE id=:id`
	_, err := r.db.NamedExecContext(ctx, query, book)
	if err != nil {
		return fmt.Errorf("error updating book: %w", err)
	}
	return nil
}

// DeleteBook deletes a book by its ID.
func (r *bookRepository) DeleteBook(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM books WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("error deleting book: %w", err)
	}
	return nil
}
