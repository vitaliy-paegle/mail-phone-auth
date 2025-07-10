package api

import (
	"log"
	"mail-phone-auth/internal/entities/auth"
	"mail-phone-auth/internal/entities/user"
	"mail-phone-auth/pkg/jwt"
	"mail-phone-auth/pkg/postgres"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

//@title OpenApiTitle

type API struct {
	Router *http.ServeMux
	postgres *postgres.Postgres
	jwt *jwt.JWT

	authRepository *auth.Repository
	authController *auth.Controller

	userRepository *user.Repository
	userHandler *user.Handler
}

func New(postgres *postgres.Postgres, jwt *jwt.JWT) *API {
	api := API{}
	api.Router = http.NewServeMux()
	api.postgres = postgres
	api.jwt = jwt

	api.authRepository = auth.NewRepository(api.postgres)
	api.authController = auth.NewController(api.Router, api.authRepository)

	api.userRepository = user.NewRepository(api.postgres)
	api.userHandler = user.NewHandler(api.Router, api.userRepository, &api)


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