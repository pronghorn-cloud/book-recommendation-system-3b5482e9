package repositories

import (
	"context"
	"fmt"

	"book-recommendation-system/backend/models"
	"github.com/jmoiron/sqlx"
)

// RecommendationRepository defines the interface for recommendation data operations.
type RecommendationRepository interface {
	GetAllRecommendations(ctx context.Context) ([]models.Recommendation, error)
	GetRecommendationsByUserID(ctx context.Context, userID string) ([]models.Recommendation, error)
	GetRecommendationsByUserID(ctx context.Context, userID string) ([]models.Recommendation, error)
	CreateRecommendation(ctx context.Context, recommendation *models.Recommendation) error
	UpdateRecommendation(ctx context.Context, recommendation *models.Recommendation) error
	DeleteRecommendation(ctx context.Context, id string) error
}

// recommendationRepository implements RecommendationRepository using sqlx.
type recommendationRepository struct {
	db *sqlx.DB
}

// NewRecommendationRepository creates a new RecommendationRepository.
func NewRecommendationRepository(db *sqlx.DB) RecommendationRepository {
	return &recommendationRepository{db: db}
}

// GetRecommendationByID retrieves a recommendation by its ID.
func (r *recommendationRepository) GetRecommendationByID(ctx context.Context, id string) (*models.Recommendation, error) {
	var recommendation models.Recommendation
	err := r.db.GetContext(ctx, &recommendation, "SELECT id, user_id, book_id, score, generated_at FROM recommendations WHERE id=$1", id)
// GetAllRecommendations retrieves all recommendations.
func (r *recommendationRepository) GetAllRecommendations(ctx context.Context) ([]models.Recommendation, error) {
	var recommendations []models.Recommendation
	err := r.db.SelectContext(ctx, &recommendations, "SELECT id, user_id, book_id, score, generated_at FROM recommendations")
	if err != nil {
		return nil, fmt.Errorf("error getting all recommendations: %w", err)
	}
	return recommendations, nil
}

// GetRecommendationsByUserID retrieves recommendations for a specific user.
		return nil, fmt.Errorf("error getting recommendation by ID: %w", err)
	}
	return &recommendation, nil
}

// GetRecommendationsByUserID retrieves recommendations for a specific user.
func (r *recommendationRepository) GetRecommendationsByUserID(ctx context.Context, userID string) ([]models.Recommendation, error) {
	var recommendations []models.Recommendation
	err := r.db.SelectContext(ctx, &recommendations, "SELECT id, user_id, book_id, score, generated_at FROM recommendations WHERE user_id=$1", userID)
	if err != nil {
		return nil, fmt.Errorf("error getting recommendations by user ID: %w", err)
	}
	return recommendations, nil
}

// CreateRecommendation creates a new recommendation.
func (r *recommendationRepository) CreateRecommendation(ctx context.Context, recommendation *models.Recommendation) error {
	query := `INSERT INTO recommendations (id, user_id, book_id, score, generated_at) VALUES (:id, :user_id, :book_id, :score, :generated_at)`
	_, err := r.db.NamedExecContext(ctx, query, recommendation)
	if err != nil {
		return fmt.Errorf("error creating recommendation: %w", err)
	}
	return nil
}

// UpdateRecommendation updates an existing recommendation.
func (r *recommendationRepository) UpdateRecommendation(ctx context.Context, recommendation *models.Recommendation) error {
	query := `UPDATE recommendations SET user_id=:user_id, book_id=:book_id, score=:score, generated_at=:generated_at WHERE id=:id`
	_, err := r.db.NamedExecContext(ctx, query, recommendation)
	if err != nil {
		return fmt.Errorf("error updating recommendation: %w", err)
	}
	return nil
}

// DeleteRecommendation deletes a recommendation by its ID.
func (r *recommendationRepository) DeleteRecommendation(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM recommendations WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("error deleting recommendation: %w", err)
	}
	return nil
}
