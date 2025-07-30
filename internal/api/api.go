package api

import (
	"log"
	"mail-phone-auth/internal/entities/auth"
	"mail-phone-auth/internal/entities/user"
	"mail-phone-auth/pkg/jino_mail"
	"mail-phone-auth/pkg/jwt"
	"mail-phone-auth/pkg/postgres"
	"net/http"

	"github.com/swaggest/swgui"
	"github.com/swaggest/swgui/v5emb"
)

type API struct {
	Router         *http.ServeMux
	postgres       *postgres.Postgres
	jwt            *jwt.JWT
	jinoMail       *jino_mail.JinoMail
	authRepository *auth.Repository
	authController *auth.Controller
	userRepository *user.Repository
	userController    *user.Controller
}

func New(postgres *postgres.Postgres, jwt *jwt.JWT, jinoMail *jino_mail.JinoMail) *API {
	api := API{}
	api.Router = http.NewServeMux()
	api.postgres = postgres
	api.jwt = jwt
	api.jinoMail = jinoMail

	api.userRepository = user.NewRepository(api.postgres)
	api.userController = user.NewController(api.Router, api.userRepository, &api)

	api.authRepository = auth.NewRepository(api.postgres)
	api.authController = auth.NewController(api.Router, api.authRepository, api.jinoMail, api.jwt, api.userRepository)



	api.OpenAPIconnect()

	return &api
}

func (api *API) OpenAPIconnect() {

	config := swgui.Config{}

	// Show SCHEMAS: {"defaultModelsExpandDepth": "1"} 
	config.SettingsUI= map[string]string{"defaultModelsExpandDepth": "-1"} 
	config.Title = "OpenAPI"
	config.SwaggerJSON = "/static/openapi.json"
	config.BasePath = "/"

	handler := v5emb.NewHandlerWithConfig(config)

	api.Router.Handle("/", handler)
}

func (api *API) TestAPI() {
	log.Println("TEST API")
}
