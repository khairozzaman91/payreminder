package userhandler

import (
	"encoding/json"
	"net/http"
	"payreminder/domain"
)

type CreateUserRequest struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	BusinessName string `json:"business_name"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	// Decode JSON request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Basic validation
	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Temporary response
	user := domain.User{
		Email:        req.Email,
		BusinessName: req.BusinessName,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)
}
