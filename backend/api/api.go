package api

import (
	"net/http"

	"github.com/Quan0308/main-api/interfaces"
)

func RegisterRoutes(server *http.ServeMux, container interfaces.Container) {
	authAPI := container.GetAuthAPI()
	authAPI.RegisterRoutes(server)
}
