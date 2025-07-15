package api

import (
	"log"
	"mail-phone-auth/internal/api/swagger"
	"mail-phone-auth/internal/entities/auth"
	"mail-phone-auth/internal/entities/user"
	"mail-phone-auth/pkg/jwt"
	"mail-phone-auth/pkg/postgres"
	"net/http"
)

//	@title			API Пользователи и авторизация
//	@version		1.0
//	@description	---

type API struct {
	Router *http.ServeMux
	postgres *postgres.Postgres
	jwt *jwt.JWT

	swagger *swagger.Swagger

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

	api.swagger = swagger.New(api.Router)

	api.authRepository = auth.NewRepository(api.postgres)
	api.authController = auth.NewController(api.Router, api.authRepository)

	api.userRepository = user.NewRepository(api.postgres)
	api.userHandler = user.NewHandler(api.Router, api.userRepository, &api)
	
	return &api
}


func (api *API) TestAPI() {
	log.Println("TEST API")
}