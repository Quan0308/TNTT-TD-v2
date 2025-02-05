package api

import (
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/Quan0308/main-api/core/dtos"
	"github.com/Quan0308/main-api/core/middlewares"
	"github.com/Quan0308/main-api/core/utils"
	authDto "github.com/Quan0308/main-api/dtos"
	messagesEnum "github.com/Quan0308/main-api/enums/messages"
	"github.com/Quan0308/main-api/interfaces"
	"github.com/Quan0308/main-api/types"
)

type AuthAPIImpl struct {
	authHandler         interfaces.AuthHandler
	firebaseAuthService interfaces.FirebaseAuthService
}

func NewAuthAPI(authHandler interfaces.AuthHandler, firebaseAuthService interfaces.FirebaseAuthService) interfaces.AuthAPI {
	return &AuthAPIImpl{authHandler: authHandler, firebaseAuthService: firebaseAuthService}
}

func (api *AuthAPIImpl) LogInHandler(w http.ResponseWriter, r *http.Request) {
	token, ok := r.Context().Value(types.UserTokenKey).(*auth.Token)
	if !ok {
		http.Error(w, "No token found", http.StatusUnauthorized)
		return
	}

	accessToken, err := api.authHandler.LogIn(r.Context(), token.UID)
	if err != nil {
		utils.Response(w, dtos.Response{
			Message:    "Error when get accessToken",
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	data := authDto.LoginRes{
		AccessToken: accessToken,
	}

	response := dtos.Response{
		Data:       data,
		Message:    string(messagesEnum.LoginSuccess),
		StatusCode: http.StatusOK,
	}

	utils.Response(w, response)
}

func (api *AuthAPIImpl) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/auth/login", middlewares.VerifyIdToken(http.HandlerFunc(api.LogInHandler), api.firebaseAuthService))
}
