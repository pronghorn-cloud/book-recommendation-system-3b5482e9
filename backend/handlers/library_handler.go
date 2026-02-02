package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"book-recommendation-system/backend/models"
	"book-recommendation-system/backend/services"
	"github.com/go-chi/chi/v5"
)

// LibraryHandler handles HTTP requests for libraries.
type LibraryHandler struct {
	service services.LibraryService
}

// NewLibraryHandler creates a new LibraryHandler.
func NewLibraryHandler(s services.LibraryService) *LibraryHandler {
	return &LibraryHandler{service: s}
}

// GetLibraryByID handles the request to get a library by its ID.
func (h *LibraryHandler) GetLibraryByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Library ID is required", http.StatusBadRequest)
		return
	}

	library, err := h.service.GetLibraryByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if library == nil {
		http.Error(w, "Library not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(library)
}

// GetAllLibraries handles the request to get all libraries.
func (h *LibraryHandler) GetAllLibraries(w http.ResponseWriter, r *http.Request) {
	libraries, err := h.service.GetAllLibraries(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(libraries)
}

// CreateLibrary handles the request to create a new library.
func (h *LibraryHandler) CreateLibrary(w http.ResponseWriter, r *http.Request) {
	var library models.Library
	err := json.NewDecoder(r.Body).Decode(&library)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateLibrary(r.Context(), &library)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(library)
}

// UpdateLibrary handles the request to update an existing library.
func (h *LibraryHandler) UpdateLibrary(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Library ID is required", http.StatusBadRequest)
		return
	}

	var library models.Library
	err := json.NewDecoder(r.Body).Decode(&library)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	library.ID = id // Ensure the ID from the URL is used

	err = h.service.UpdateLibrary(r.Context(), &library)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(library)
}

// DeleteLibrary handles the request to delete a library by its ID.
func (h *LibraryHandler) DeleteLibrary(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Library ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteLibrary(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
