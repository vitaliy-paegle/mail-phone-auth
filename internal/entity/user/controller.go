package user

import (
	"mail-phone-auth/internal/entity"
	"mail-phone-auth/pkg/postgres"
	"net/http"
)

type Controller struct {
	Router   *http.ServeMux
	Postgres *postgres.Postgres
	Entity   *entity.Entity[User, UserData]
}

func NewController(router *http.ServeMux, postgres *postgres.Postgres) *Controller {

	controller := Controller{
		Router:   router,
		Postgres: postgres,
	}

	controller.Entity = entity.New[User, UserData](controller.Postgres)

	controller.Router.HandleFunc("GET /api/user/{id}", controller.Entity.Read)
	controller.Router.HandleFunc("GET /api/user/all", controller.Entity.ReadAll)
	controller.Router.HandleFunc("POST /api/user", controller.Entity.Create)
	controller.Router.HandleFunc("PUT /api/user/{id}", controller.Entity.Update)
	controller.Router.HandleFunc("DELETE /api/user/{id}", controller.Entity.Delete)

	return &controller
}
