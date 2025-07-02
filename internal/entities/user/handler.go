package user

import (
	"log"
	"mail-phone-auth/internal/api/request"
	"net/http"
)

type Handler struct{
	Router *http.ServeMux
}

func NewHandler(router *http.ServeMux) *Handler {
	handler := Handler{Router: router}

	handler.Router.HandleFunc("POST /api/user", handler.Create)
	handler.Router.HandleFunc("GET /api/user/{id}", handler.Read)
	handler.Router.HandleFunc("PUT /api/user/{id}", handler.Update)
	handler.Router.HandleFunc("DELETE /api/user/{id}", handler.Delete)

	return &handler
}

func (handler *Handler) Create(w http.ResponseWriter, r *http.Request) {
	log.Println(request.DecodeBody[UserCreateRequest](r.Body))
}

func (handler *Handler) Read(w http.ResponseWriter, r *http.Request) {
	
}

func (handler *Handler) Update(w http.ResponseWriter, r *http.Request) {
	
}

func (handler *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	
}