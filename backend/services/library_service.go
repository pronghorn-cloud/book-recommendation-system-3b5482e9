package services

import (
	"context"
	"fmt"

	"book-recommendation-system/backend/models"
	"book-recommendation-system/backend/repositories"
)

// LibraryService defines the interface for library-related business logic.
type LibraryService interface {
	GetLibraryByID(ctx context.Context, id string) (*models.Library, error)
	GetAllLibraries(ctx context.Context) ([]models.Library, error)
	CreateLibrary(ctx context.Context, library *models.Library) error
	UpdateLibrary(ctx context.Context, library *models.Library) error
	DeleteLibrary(ctx context.Context, id string) error
}

// libraryService implements LibraryService.
type libraryService struct {
	repo repositories.LibraryRepository
}

// NewLibraryService creates a new LibraryService.
func NewLibraryService(repo repositories.LibraryRepository) LibraryService {
	return &libraryService{repo: repo}
}

// GetLibraryByID retrieves a library by its ID using the repository.
func (s *libraryService) GetLibraryByID(ctx context.Context, id string) (*models.Library, error) {
	library, err := s.repo.GetLibraryByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get library by ID: %w", err)
	}
	return library, nil
}

// GetAllLibraries retrieves all libraries using the repository.
func (s *libraryService) GetAllLibraries(ctx context.Context) ([]models.Library, error) {
	libraries, err := s.repo.GetAllLibraries(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get all libraries: %w", err)
	}
	return libraries, nil
}

// CreateLibrary creates a new library using the repository.
func (s *libraryService) CreateLibrary(ctx context.Context, library *models.Library) error {
	// Add any business logic/validation before creating a library
	err := s.repo.CreateLibrary(ctx, library)
	if err != nil {
		return fmt.Errorf("service: failed to create library: %w", err)
	}
	return nil
}

// UpdateLibrary updates an existing library using the repository.
func (s *libraryService) UpdateLibrary(ctx context.Context, library *models.Library) error {
	// Add any business logic/validation before updating a library
	err := s.repo.UpdateLibrary(ctx, library)
	if err != nil {
		return fmt.Errorf("service: failed to update library: %w", err)
	}
	return nil
}

// DeleteLibrary deletes a library by its ID using the repository.
func (s *libraryService) DeleteLibrary(ctx context.Context, id string) error {
	// Add any business logic/validation before deleting a library
	err := s.repo.DeleteLibrary(ctx, id)
	if err != nil {
		return fmt.Errorf("service: failed to delete library: %w", err)
	}
	return nil
}
