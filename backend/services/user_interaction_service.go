package services

import (
	"context"
	"fmt"

	"book-recommendation-system/backend/models"
	"book-recommendation-system/backend/repositories"
)

// UserInteractionService defines the interface for user interaction-related business logic.
type UserInteractionService interface {
	GetAllUserInteractions(ctx context.Context) ([]models.UserInteraction, error)
	GetUserInteractionsByUserID(ctx context.Context, userID string) ([]models.UserInteraction, error)
	GetUserInteractionsByUserID(ctx context.Context, userID string) ([]models.UserInteraction, error)
	CreateUserInteraction(ctx context.Context, userInteraction *models.UserInteraction) error
	UpdateUserInteraction(ctx context.Context, userInteraction *models.UserInteraction) error
	DeleteUserInteraction(ctx context.Context, id string) error
}

// userInteractionService implements UserInteractionService.
type userInteractionService struct {
	repo repositories.UserInteractionRepository
}

// NewUserInteractionService creates a new UserInteractionService.
func NewUserInteractionService(repo repositories.UserInteractionRepository) UserInteractionService {
	return &userInteractionService{repo: repo}
}

// GetUserInteractionByID retrieves a user interaction by its ID using the repository.
func (s *userInteractionService) GetUserInteractionByID(ctx context.Context, id string) (*models.UserInteraction, error) {
	userInteraction, err := s.repo.GetUserInteractionByID(ctx, id)
	if err != nil {
// GetAllUserInteractions retrieves all user interactions using the repository.
func (s *userInteractionService) GetAllUserInteractions(ctx context.Context) ([]models.UserInteraction, error) {
	userInteractions, err := s.repo.GetAllUserInteractions(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get all user interactions: %w", err)
	}
	return userInteractions, nil
}

// GetUserInteractionsByUserID retrieves user interactions for a specific user using the repository.
	}
	return userInteraction, nil
}

// GetUserInteractionsByUserID retrieves user interactions for a specific user using the repository.
func (s *userInteractionService) GetUserInteractionsByUserID(ctx context.Context, userID string) ([]models.UserInteraction, error) {
	userInteractions, err := s.repo.GetUserInteractionsByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get user interactions by user ID: %w", err)
	}
	return userInteractions, nil
}

// CreateUserInteraction creates a new user interaction using the repository.
func (s *userInteractionService) CreateUserInteraction(ctx context.Context, userInteraction *models.UserInteraction) error {
	// Add any business logic/validation before creating a user interaction
	err := s.repo.CreateUserInteraction(ctx, userInteraction)
	if err != nil {
		return fmt.Errorf("service: failed to create user interaction: %w", err)
	}
	return nil
}

// UpdateUserInteraction updates an existing user interaction using the repository.
func (s *userInteractionService) UpdateUserInteraction(ctx context.Context, userInteraction *models.UserInteraction) error {
	// Add any business logic/validation before updating a user interaction
	err := s.repo.UpdateUserInteraction(ctx, userInteraction)
	if err != nil {
		return fmt.Errorf("service: failed to update user interaction: %w", err)
	}
	return nil
}

// DeleteUserInteraction deletes a user interaction by its ID using the repository.
func (s *userInteractionService) DeleteUserInteraction(ctx context.Context, id string) error {
	// Add any business logic/validation before deleting a user interaction
	err := s.repo.DeleteUserInteraction(ctx, id)
	if err != nil {
		return fmt.Errorf("service: failed to delete user interaction: %w", err)
	}
	return nil
}
