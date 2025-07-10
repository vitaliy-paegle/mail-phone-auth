package user

import (
	"log"
	"mail-phone-auth/internal/api/request"
	"mail-phone-auth/internal/api/response"
	"net/http"
	"strconv"
)


type IAPI interface{
	TestAPI()
}


type Handler struct{
	Router *http.ServeMux
	Repository *Repository
	API IAPI
}

func NewHandler(router *http.ServeMux, repository *Repository, api IAPI) *Handler {
	handler := Handler{
		Router: router,
		Repository: repository,
		API: api,
	}

	handler.API.TestAPI()

	handler.Router.HandleFunc("GET /api/user/{id}", handler.Read)
	handler.Router.HandleFunc("GET /api/user/all", handler.ReadAll)
	handler.Router.HandleFunc("POST /api/user", handler.Create)	
	handler.Router.HandleFunc("PATCH /api/user/{id}", handler.Update)
	handler.Router.HandleFunc("DELETE /api/user/{id}", handler.Delete)

	return &handler
}

func (handler *Handler) Create(w http.ResponseWriter, r *http.Request) {

	body, err := request.DecodeBody[UserCreateRequest](r.Body)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := NewUser(body)

	err = handler.Repository.Create(user)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		response.JSON(w, user, http.StatusCreated)
	}

}

// Get User
// @summary  Get User by ID
// @router   /api/user/{id} [get]
// @tags     user
// @success  200
// @failure  400

func (handler *Handler) Read(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := handler.Repository.Read(id)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		response.JSON(w, user, http.StatusOK)
	}

}

func (handler *Handler) ReadAll(w http.ResponseWriter, r *http.Request) {

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

	if err != nil {limit = -1}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))

	if err != nil {offset = 0}

	users := handler.Repository.ReadAll(limit, offset)

	response.JSON(w, &users, http.StatusOK)

}

func (handler *Handler) Update(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := request.DecodeBody[UserUpdateRequest](r.Body)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// user := User{
	// 	Model: gorm.Model{ID: uint(id)},
	// 	Name: body.Name,
	// 	Phone: body.Phone,		
	// }

	user, err := handler.Repository.Read(id)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Name = body.Name
	user.Phone = body.Phone

	err =  handler.Repository.Update(user)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		response.JSON(w, user, http.StatusOK)
	}
}

func (handler *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := handler.Repository.Read(id)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.Repository.Delete(id)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		response.JSON(w, user, http.StatusOK)
	}
}