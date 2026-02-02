package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"book-recommendation-system/backend/models"
	"book-recommendation-system/backend/services"
	"github.com/go-chi/chi/v5"
)

// RecommendationHandler handles HTTP requests for recommendations.
type RecommendationHandler struct {
	service services.RecommendationService
}

// NewRecommendationHandler creates a new RecommendationHandler.
	return &RecommendationHandler{service: s}
}

// GetAllRecommendations handles the request to get all recommendations.
func (h *RecommendationHandler) GetAllRecommendations(w http.ResponseWriter, r *http.Request) {
	recommendations, err := h.service.GetAllRecommendations(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendations)
}

// GetRecommendationByID handles the request to get a recommendation by its ID.
	return &RecommendationHandler{service: s}
}

// GetRecommendationByID handles the request to get a recommendation by its ID.
func (h *RecommendationHandler) GetRecommendationByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Recommendation ID is required", http.StatusBadRequest)
		return
	}

	recommendation, err := h.service.GetRecommendationByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if recommendation == nil {
		http.Error(w, "Recommendation not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendation)
}

// GetRecommendationsByUserID handles the request to get recommendations for a specific user ID.
func (h *RecommendationHandler) GetRecommendationsByUserID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	recommendations, err := h.service.GetRecommendationsByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendations)
}

// CreateRecommendation handles the request to create a new recommendation.
func (h *RecommendationHandler) CreateRecommendation(w http.ResponseWriter, r *http.Request) {
	var recommendation models.Recommendation
	err := json.NewDecoder(r.Body).Decode(&recommendation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateRecommendation(r.Context(), &recommendation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(recommendation)
}

// UpdateRecommendation handles the request to update an existing recommendation.
func (h *RecommendationHandler) UpdateRecommendation(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Recommendation ID is required", http.StatusBadRequest)
		return
	}

	var recommendation models.Recommendation
	err := json.NewDecoder(r.Body).Decode(&recommendation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recommendation.ID = id // Ensure the ID from the URL is used

	err = h.service.UpdateRecommendation(r.Context(), &recommendation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendation)
}

// DeleteRecommendation handles the request to delete a recommendation by its ID.
func (h *RecommendationHandler) DeleteRecommendation(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Recommendation ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteRecommendation(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
