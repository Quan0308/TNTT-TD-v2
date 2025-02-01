package api

import (
	"net/http"

	"github.com/Quan0308/main-api/core/dtos"
	"github.com/Quan0308/main-api/core/utils"
	authDto "github.com/Quan0308/main-api/dtos"
	messagesEnum "github.com/Quan0308/main-api/enums/messages"
	"github.com/Quan0308/main-api/interfaces"
)

type AuthAPI struct {
	authHandler interfaces.AuthHandler
}

func NewAuthAPI(authHandler interfaces.AuthHandler) *AuthAPI {
	return &AuthAPI{authHandler: authHandler}
}

func (api *AuthAPI) SignInHandler(w http.ResponseWriter, r *http.Request) {
	var payload authDto.LoginDto

	decodeErr := utils.Decode(r, &payload)

	if decodeErr != nil {
		utils.Response(w, dtos.Response{
			Message:    "Invalid Request body",
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	message, validateErr := utils.DoValidate(payload)

	if validateErr != nil {
		utils.Response(w, dtos.Response{
			Message:    message,
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	userId, handlerErr := api.authHandler.LogIn(r.Context(), payload.Email, payload.Password)
	if handlerErr != nil {
		utils.Response(w, dtos.Response{
			Message:    "Error when sign in",
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	data := authDto.LoginRes{
		AccessToken:  userId.String(),
		RefreshToken: userId.String(),
	}

	response := dtos.Response{
		Data:       data,
		Message:    string(messagesEnum.LoginSuccess),
		StatusCode: http.StatusOK,
	}

	utils.Response(w, response)
}

func (api *AuthAPI) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /login", api.SignInHandler)
}
