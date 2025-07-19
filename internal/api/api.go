package api

import (
	"log"
	"mail-phone-auth/internal/entities/auth"
	"mail-phone-auth/internal/entities/user"
	"mail-phone-auth/pkg/jwt"
	"mail-phone-auth/pkg/postgres"
	"net/http"

	"github.com/swaggest/swgui/v5emb"
)

type API struct {
	Router         *http.ServeMux
	postgres       *postgres.Postgres
	jwt            *jwt.JWT
	authRepository *auth.Repository
	authController *auth.Controller
	userRepository *user.Repository
	userHandler    *user.Handler
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

	api.OpenAPIconnect()

	return &api
}

func (api *API) OpenAPIconnect() {
	api.Router.Handle("/openapi/", v5emb.New("OpenAPI", "/static/openapi.json", "/openapi/"))
}

func (api *API) TestAPI() {
	log.Println("TEST API")
}
