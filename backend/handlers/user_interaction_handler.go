package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"book-recommendation-system/backend/models"
	"book-recommendation-system/backend/services"
	"github.com/go-chi/chi/v5"
)

// UserInteractionHandler handles HTTP requests for user interactions.
type UserInteractionHandler struct {
	service services.UserInteractionService
}

// NewUserInteractionHandler creates a new UserInteractionHandler.
	return &UserInteractionHandler{service: s}
}

// GetAllUserInteractions handles the request to get all user interactions.
func (h *UserInteractionHandler) GetAllUserInteractions(w http.ResponseWriter, r *http.Request) {
	userInteractions, err := h.service.GetAllUserInteractions(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userInteractions)
}

// GetUserInteractionByID handles the request to get a user interaction by its ID.
	return &UserInteractionHandler{service: s}
}

// GetUserInteractionByID handles the request to get a user interaction by its ID.
func (h *UserInteractionHandler) GetUserInteractionByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "User Interaction ID is required", http.StatusBadRequest)
		return
	}

	userInteraction, err := h.service.GetUserInteractionByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if userInteraction == nil {
		http.Error(w, "User Interaction not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userInteraction)
}

// GetUserInteractionsByUserID handles the request to get user interactions for a specific user ID.
func (h *UserInteractionHandler) GetUserInteractionsByUserID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	userInteractions, err := h.service.GetUserInteractionsByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userInteractions)
}

// CreateUserInteraction handles the request to create a new user interaction.
func (h *UserInteractionHandler) CreateUserInteraction(w http.ResponseWriter, r *http.Request) {
	var userInteraction models.UserInteraction
	err := json.NewDecoder(r.Body).Decode(&userInteraction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateUserInteraction(r.Context(), &userInteraction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userInteraction)
}

// UpdateUserInteraction handles the request to update an existing user interaction.
func (h *UserInteractionHandler) UpdateUserInteraction(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "User Interaction ID is required", http.StatusBadRequest)
		return
	}

	var userInteraction models.UserInteraction
	err := json.NewDecoder(r.Body).Decode(&userInteraction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userInteraction.ID = id // Ensure the ID from the URL is used

	err = h.service.UpdateUserInteraction(r.Context(), &userInteraction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userInteraction)
}

// DeleteUserInteraction handles the request to delete a user interaction by its ID.
func (h *UserInteractionHandler) DeleteUserInteraction(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "User Interaction ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteUserInteraction(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
