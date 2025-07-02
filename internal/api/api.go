package api

import (
	"mail-phone-auth/internal/entities/user"
	"net/http"
)

type API struct {
	Router *http.ServeMux
	UserHandler *user.Handler
}

func New() *API {
	api := API{}

	api.Router = http.NewServeMux()
	api.UserHandler = user.NewHandler(api.Router)

	return &api
}
