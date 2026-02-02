package repositories

import (
	"context"
	"fmt"

	"book-recommendation-system/backend/models"
	"github.com/jmoiron/sqlx"
)

// UserInteractionRepository defines the interface for user interaction data operations.
type UserInteractionRepository interface {
	GetAllUserInteractions(ctx context.Context) ([]models.UserInteraction, error)
	GetUserInteractionsByUserID(ctx context.Context, userID string) ([]models.UserInteraction, error)
	GetUserInteractionsByUserID(ctx context.Context, userID string) ([]models.UserInteraction, error)
	CreateUserInteraction(ctx context.Context, userInteraction *models.UserInteraction) error
	UpdateUserInteraction(ctx context.Context, userInteraction *models.UserInteraction) error
	DeleteUserInteraction(ctx context.Context, id string) error
}

// userInteractionRepository implements UserInteractionRepository using sqlx.
type userInteractionRepository struct {
	db *sqlx.DB
}

// NewUserInteractionRepository creates a new UserInteractionRepository.
func NewUserInteractionRepository(db *sqlx.DB) UserInteractionRepository {
	return &userInteractionRepository{db: db}
}

// GetUserInteractionByID retrieves a user interaction by its ID.
func (r *userInteractionRepository) GetUserInteractionByID(ctx context.Context, id string) (*models.UserInteraction, error) {
	var userInteraction models.UserInteraction
	err := r.db.GetContext(ctx, &userInteraction, "SELECT id, user_id, book_id, interaction_type, timestamp FROM user_interactions WHERE id=$1", id)
// GetAllUserInteractions retrieves all user interactions.
func (r *userInteractionRepository) GetAllUserInteractions(ctx context.Context) ([]models.UserInteraction, error) {
	var userInteractions []models.UserInteraction
	err := r.db.SelectContext(ctx, &userInteractions, "SELECT id, user_id, book_id, interaction_type, timestamp FROM user_interactions")
	if err != nil {
		return nil, fmt.Errorf("error getting all user interactions: %w", err)
	}
	return userInteractions, nil
}

// GetUserInteractionsByUserID retrieves user interactions for a specific user.
		return nil, fmt.Errorf("error getting user interaction by ID: %w", err)
	}
	return &userInteraction, nil
}

// GetUserInteractionsByUserID retrieves user interactions for a specific user.
func (r *userInteractionRepository) GetUserInteractionsByUserID(ctx context.Context, userID string) ([]models.UserInteraction, error) {
	var userInteractions []models.UserInteraction
	err := r.db.SelectContext(ctx, &userInteractions, "SELECT id, user_id, book_id, interaction_type, timestamp FROM user_interactions WHERE user_id=$1", userID)
	if err != nil {
		return nil, fmt.Errorf("error getting user interactions by user ID: %w", err)
	}
	return userInteractions, nil
}

// CreateUserInteraction creates a new user interaction.
func (r *userInteractionRepository) CreateUserInteraction(ctx context.Context, userInteraction *models.UserInteraction) error {
	query := `INSERT INTO user_interactions (id, user_id, book_id, interaction_type, timestamp) VALUES (:id, :user_id, :book_id, :interaction_type, :timestamp)`
	_, err := r.db.NamedExecContext(ctx, query, userInteraction)
	if err != nil {
		return fmt.Errorf("error creating user interaction: %w", err)
	}
	return nil
}

// UpdateUserInteraction updates an existing user interaction.
func (r *userInteractionRepository) UpdateUserInteraction(ctx context.Context, userInteraction *models.UserInteraction) error {
	query := `UPDATE user_interactions SET user_id=:user_id, book_id=:book_id, interaction_type=:interaction_type, timestamp=:timestamp WHERE id=:id`
	_, err := r.db.NamedExecContext(ctx, query, userInteraction)
	if err != nil {
		return fmt.Errorf("error updating user interaction: %w", err)
	}
	return nil
}

// DeleteUserInteraction deletes a user interaction by its ID.
func (r *userInteractionRepository) DeleteUserInteraction(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM user_interactions WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("error deleting user interaction: %w", err)
	}
	return nil
}
