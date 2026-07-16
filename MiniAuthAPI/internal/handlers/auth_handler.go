package handlers

import (
	"MiniProject/internal/auth"
	"MiniProject/internal/models"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	service *auth.AuthService
}

func NewAuthHandeler(service *auth.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(
	w http.ResponseWriter,
	r *http.Request,
) {
	var request models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(
		request.Email,
		request.Password,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := models.LoginResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
