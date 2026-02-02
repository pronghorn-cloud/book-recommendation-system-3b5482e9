package models

type Library struct {
	ID        string  `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	Address   string  `json:"address" db:"address"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
}
