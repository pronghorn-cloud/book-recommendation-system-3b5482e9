package models

type Book struct {
	ID              string `json:"id" db:"id"`
	Title           string `json:"title" db:"title"`
	Author          string `json:"author" db:"author"`
	ISBN            string `json:"isbn" db:"isbn"`
	Description     string `json:"description" db:"description"`
	CoverImageURL   string `json:"cover_image_url" db:"cover_image_url"`
	Genre           string `json:"genre" db:"genre"`
	PublicationYear int    `json:"publication_year" db:"publication_year"`
}
