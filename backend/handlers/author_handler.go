package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"book-recommendation-system/backend/models"
	"book-recommendation-system/backend/services"
	"github.com/go-chi/chi/v5"
)

// AuthorHandler handles HTTP requests for authors.
type AuthorHandler struct {
	service services.AuthorService
}

// NewAuthorHandler creates a new AuthorHandler.
func NewAuthorHandler(s services.AuthorService) *AuthorHandler {
	return &AuthorHandler{service: s}
}

// GetAuthorByID handles the request to get an author by their ID.
func (h *AuthorHandler) GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Author ID is required", http.StatusBadRequest)
		return
	}

	author, err := h.service.GetAuthorByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if author == nil {
		http.Error(w, "Author not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

// GetAllAuthors handles the request to get all authors.
func (h *AuthorHandler) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := h.service.GetAllAuthors(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

// CreateAuthor handles the request to create a new author.
func (h *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateAuthor(r.Context(), &author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}

// UpdateAuthor handles the request to update an existing author.
func (h *AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Author ID is required", http.StatusBadRequest)
		return
	}

	var author models.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	author.ID = id // Ensure the ID from the URL is used

	err = h.service.UpdateAuthor(r.Context(), &author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

// DeleteAuthor handles the request to delete an author by their ID.
func (h *AuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Author ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteAuthor(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
