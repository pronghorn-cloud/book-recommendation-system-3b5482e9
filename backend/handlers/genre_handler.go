package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"book-recommendation-system/backend/models"
	"book-recommendation-system/backend/services"
	"github.com/go-chi/chi/v5"
)

// GenreHandler handles HTTP requests for genres.
type GenreHandler struct {
	service services.GenreService
}

// NewGenreHandler creates a new GenreHandler.
func NewGenreHandler(s services.GenreService) *GenreHandler {
	return &GenreHandler{service: s}
}

// GetGenreByID handles the request to get a genre by its ID.
func (h *GenreHandler) GetGenreByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Genre ID is required", http.StatusBadRequest)
		return
	}

	genre, err := h.service.GetGenreByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if genre == nil {
		http.Error(w, "Genre not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genre)
}

// GetAllGenres handles the request to get all genres.
func (h *GenreHandler) GetAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := h.service.GetAllGenres(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genres)
}

// CreateGenre handles the request to create a new genre.
func (h *GenreHandler) CreateGenre(w http.ResponseWriter, r *http.Request) {
	var genre models.Genre
	err := json.NewDecoder(r.Body).Decode(&genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateGenre(r.Context(), &genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(genre)
}

// UpdateGenre handles the request to update an existing genre.
func (h *GenreHandler) UpdateGenre(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Genre ID is required", http.StatusBadRequest)
		return
	}

	var genre models.Genre
	err := json.NewDecoder(r.Body).Decode(&genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	genre.ID = id // Ensure the ID from the URL is used

	err = h.service.UpdateGenre(r.Context(), &genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genre)
}

// DeleteGenre handles the request to delete a genre by its ID.
func (h *GenreHandler) DeleteGenre(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Genre ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteGenre(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
