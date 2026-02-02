package services

import (
	"context"
	"fmt"

	"book-recommendation-system/backend/models"
	"book-recommendation-system/backend/repositories"
)

// RecommendationService defines the interface for recommendation-related business logic.
type RecommendationService interface {
	GetAllRecommendations(ctx context.Context) ([]models.Recommendation, error)
	GetRecommendationsByUserID(ctx context.Context, userID string) ([]models.Recommendation, error)
	GetRecommendationsByUserID(ctx context.Context, userID string) ([]models.Recommendation, error)
	CreateRecommendation(ctx context.Context, recommendation *models.Recommendation) error
	UpdateRecommendation(ctx context.Context, recommendation *models.Recommendation) error
	DeleteRecommendation(ctx context.Context, id string) error
}

// recommendationService implements RecommendationService.
type recommendationService struct {
	repo repositories.RecommendationRepository
}

// NewRecommendationService creates a new RecommendationService.
func NewRecommendationService(repo repositories.RecommendationRepository) RecommendationService {
	return &recommendationService{repo: repo}
}

// GetRecommendationByID retrieves a recommendation by its ID using the repository.
func (s *recommendationService) GetRecommendationByID(ctx context.Context, id string) (*models.Recommendation, error) {
	recommendation, err := s.repo.GetRecommendationByID(ctx, id)
	if err != nil {
// GetAllRecommendations retrieves all recommendations using the repository.
func (s *recommendationService) GetAllRecommendations(ctx context.Context) ([]models.Recommendation, error) {
	recommendations, err := s.repo.GetAllRecommendations(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get all recommendations: %w", err)
	}
	return recommendations, nil
}

// GetRecommendationsByUserID retrieves recommendations for a specific user using the repository.
	}
	return recommendation, nil
}

// GetRecommendationsByUserID retrieves recommendations for a specific user using the repository.
func (s *recommendationService) GetRecommendationsByUserID(ctx context.Context, userID string) ([]models.Recommendation, error) {
	recommendations, err := s.repo.GetRecommendationsByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get recommendations by user ID: %w", err)
	}
	return recommendations, nil
}

// CreateRecommendation creates a new recommendation using the repository.
func (s *recommendationService) CreateRecommendation(ctx context.Context, recommendation *models.Recommendation) error {
	// Add any business logic/validation before creating a recommendation
	err := s.repo.CreateRecommendation(ctx, recommendation)
	if err != nil {
		return fmt.Errorf("service: failed to create recommendation: %w", err)
	}
	return nil
}

// UpdateRecommendation updates an existing recommendation using the repository.
func (s *recommendationService) UpdateRecommendation(ctx context.Context, recommendation *models.Recommendation) error {
	// Add any business logic/validation before updating a recommendation
	err := s.repo.UpdateRecommendation(ctx, recommendation)
	if err != nil {
		return fmt.Errorf("service: failed to update recommendation: %w", err)
	}
	return nil
}

// DeleteRecommendation deletes a recommendation by its ID using the repository.
func (s *recommendationService) DeleteRecommendation(ctx context.Context, id string) error {
	// Add any business logic/validation before deleting a recommendation
	err := s.repo.DeleteRecommendation(ctx, id)
	if err != nil {
		return fmt.Errorf("service: failed to delete recommendation: %w", err)
	}
	return nil
}
