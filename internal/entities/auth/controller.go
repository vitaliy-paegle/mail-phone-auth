package auth

import "net/http"


type Controller struct {
	router *http.ServeMux
	repository *Repository
}

func NewController(router *http.ServeMux, repository *Repository) *Controller {
	controller := Controller{
		router: router,
		repository: repository,
	}

	return  &controller
}