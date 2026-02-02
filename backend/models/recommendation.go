package models

import "time"

type Recommendation struct {
	ID          string    `json:"id" db:"id"`
	UserID      string    `json:"user_id" db:"user_id"`
	BookID      string    `json:"book_id" db:"book_id"`
	Score       float64   `json:"score" db:"score"`
	GeneratedAt time.Time `json:"generated_at" db:"generated_at"`
}
