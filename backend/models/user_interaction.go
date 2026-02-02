package models

import "time"

type UserInteraction struct {
	ID              string    `json:"id" db:"id"`
	UserID          string    `json:"user_id" db:"user_id"`
	BookID          string    `json:"book_id" db:"book_id"`
	InteractionType string    `json:"interaction_type" db:"interaction_type"` // e.g., "view", "click", "rating"
	Timestamp       time.Time `json:"timestamp" db:"timestamp"`
}
