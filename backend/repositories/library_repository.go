package repositories

import (
	"context"
	"fmt"

	"book-recommendation-system/backend/models"
	"github.com/jmoiron/sqlx"
)

// LibraryRepository defines the interface for library data operations.
type LibraryRepository interface {
	GetLibraryByID(ctx context.Context, id string) (*models.Library, error)
	GetAllLibraries(ctx context.Context) ([]models.Library, error)
	CreateLibrary(ctx context.Context, library *models.Library) error
	UpdateLibrary(ctx context.Context, library *models.Library) error
	DeleteLibrary(ctx context.Context, id string) error
}

// libraryRepository implements LibraryRepository using sqlx.
type libraryRepository struct {
	db *sqlx.DB
}

// NewLibraryRepository creates a new LibraryRepository.
func NewLibraryRepository(db *sqlx.DB) LibraryRepository {
	return &libraryRepository{db: db}
}

// GetLibraryByID retrieves a library by its ID.
func (r *libraryRepository) GetLibraryByID(ctx context.Context, id string) (*models.Library, error) {
	var library models.Library
	err := r.db.GetContext(ctx, &library, "SELECT id, name, address, latitude, longitude FROM libraries WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting library by ID: %w", err)
	}
	return &library, nil
}

// GetAllLibraries retrieves all libraries.
func (r *libraryRepository) GetAllLibraries(ctx context.Context) ([]models.Library, error) {
	var libraries []models.Library
	err := r.db.SelectContext(ctx, &libraries, "SELECT id, name, address, latitude, longitude FROM libraries")
	if err != nil {
		return nil, fmt.Errorf("error getting all libraries: %w", err)
	}
	return libraries, nil
}

// CreateLibrary creates a new library.
func (r *libraryRepository) CreateLibrary(ctx context.Context, library *models.Library) error {
	query := `INSERT INTO libraries (id, name, address, latitude, longitude) VALUES (:id, :name, :address, :latitude, :longitude)`
	_, err := r.db.NamedExecContext(ctx, query, library)
	if err != nil {
		return fmt.Errorf("error creating library: %w", err)
	}
	return nil
}

// UpdateLibrary updates an existing library.
func (r *libraryRepository) UpdateLibrary(ctx context.Context, library *models.Library) error {
	query := `UPDATE libraries SET name=:name, address=:address, latitude=:latitude, longitude=:longitude WHERE id=:id`
	_, err := r.db.NamedExecContext(ctx, query, library)
	if err != nil {
		return fmt.Errorf("error updating library: %w", err)
	}
	return nil
}

// DeleteLibrary deletes a library by its ID.
func (r *libraryRepository) DeleteLibrary(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM libraries WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("error deleting library: %w", err)
	}
	return nil
}
