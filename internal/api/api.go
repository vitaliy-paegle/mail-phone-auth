package api

import (
	"log"
	"mail-phone-auth/internal/entities/user"
	"mail-phone-auth/pkg/postgres"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type API struct {
	Router *http.ServeMux
	Postgres *postgres.Postgres

	UserRepository *user.Repository
	UserHandler *user.Handler
}

func New(postgres *postgres.Postgres) *API {
	api := API{}
	api.Router = http.NewServeMux()
	api.Postgres = postgres

	api.UserRepository = user.NewRepository(api.Postgres)
	api.UserHandler = user.NewHandler(api.Router, api.UserRepository, &api)


	api.Router.HandleFunc("GET /swagger/", api.SwaggerHandler())
	
	return &api
}

func (api *API) SwaggerHandler() http.HandlerFunc {
	return httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	)
}

func (api *API) TestAPI() {
	log.Println("TEST API")
}