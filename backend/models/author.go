package models

type Author struct {
	ID        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Biography string `json:"biography" db:"biography"`
}
