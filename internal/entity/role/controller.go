package role

import (
	"mail-phone-auth/internal/entity"
	"mail-phone-auth/pkg/postgres"
	"net/http"
)

type Controller struct {
	Router   *http.ServeMux
	Postgres *postgres.Postgres
	Entity   *entity.Entity[Role, RoleData]
}

func NewController(router *http.ServeMux, postgres *postgres.Postgres) *Controller {

	controller := Controller{
		Router:   router,
		Postgres: postgres,
	}

	controller.Entity = entity.New[Role, RoleData](controller.Postgres)

	controller.Router.HandleFunc("GET /api/role/{id}", controller.Entity.Read)
	controller.Router.HandleFunc("GET /api/role/all", controller.Entity.ReadAll)
	controller.Router.HandleFunc("POST /api/role", controller.Entity.Create)
	controller.Router.HandleFunc("PUT /api/role/{id}", controller.Entity.Update)
	controller.Router.HandleFunc("DELETE /api/role/{id}", controller.Entity.Delete)

	return &controller
}