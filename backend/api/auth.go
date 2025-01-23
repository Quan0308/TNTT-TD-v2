package api

import (
	"encoding/json"
	"net/http"

	core "github.com/Quan0308/main-api/core/interfaces"
)

type AuthAPI struct {
	authHandler core.AuthHandler
}

func NewAuthAPI(authHandler core.AuthHandler) *AuthAPI {
	return &AuthAPI{authHandler: authHandler}
}

func (api *AuthAPI) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	msg, err := api.authHandler.SignUp(payload.Email, payload.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(msg))
}

func (api *AuthAPI) SignInHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userId, err := api.authHandler.SignIn(payload.Email, payload.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"userId": userId.String()}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (api *AuthAPI) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /register", api.SignUpHandler)
	mux.HandleFunc("POST /login", api.SignInHandler)

}
