package api

import (
	"encoding/json"
	"net/http"

	"github.com/Quan0308/main-api/handlers"
	"github.com/gorilla/mux"
)

type AuthAPI struct {
	authHandler handlers.AuthHandler
}

func NewAuthAPI(authHandler handlers.AuthHandler) *AuthAPI {
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

func (api *AuthAPI) RegisterRoutes(server *mux.Router) {
	server.HandleFunc("/register", api.SignUpHandler).Methods(http.MethodPost)
}
