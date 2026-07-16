package handlers

import (
	"encoding/json"
	"net/http"
)

type ProfileHandler struct{}

func NewProfileHander() *ProfileHandler {
	return &ProfileHandler{}
}

func (h *ProfileHandler) Profile(
	w http.ResponseWriter,
	r *http.Request,
) {
	response := map[string]any{
		"id":    1,
		"email": "bhargav@gmail.com",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
