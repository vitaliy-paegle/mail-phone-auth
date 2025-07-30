package user

import (
	"log"
	"mail-phone-auth/internal/api/request"
	"mail-phone-auth/internal/api/response"
	"mail-phone-auth/internal/entity"
	"net/http"
	"strconv"
)

type IAPI interface {
	TestAPI()
}

type Controller struct {
	Router     *http.ServeMux
	Repository *Repository
	API        IAPI
}

func NewController(router *http.ServeMux, repository *Repository, api IAPI) *Controller {
	controller := Controller{
		Router:     router,
		Repository: repository,
		API:        api,
	}

	controller.API.TestAPI()

	controller.Router.HandleFunc("GET /api/user/{id}", controller.Read)
	controller.Router.HandleFunc("GET /api/user/all", controller.ReadAll)
	controller.Router.HandleFunc("POST /api/user", controller.Create)
	controller.Router.HandleFunc("PATCH /api/user/{id}", controller.Update)
	controller.Router.HandleFunc("DELETE /api/user/{id}", controller.Delete)

	return &controller
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {

	body, err := request.DecodeBody[UserCreateRequest](r.Body)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := User{}

	entity.Update(&user, &body)	

	err = c.Repository.Create(&user)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		response.JSON(w, &user, http.StatusCreated)
	}

}

func (c *Controller) Read(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := c.Repository.Read(id)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		response.JSON(w, user, http.StatusOK)
	}

}

func (c *Controller) ReadAll(w http.ResponseWriter, r *http.Request) {

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

	if err != nil {
		limit = -1
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))

	if err != nil {
		offset = 0
	}

	users, err := c.Repository.ReadAll(limit, offset)

	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response.JSON(w, &UserAllResponse{Users: users, Count: len(users)}, http.StatusOK)

}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := request.DecodeBody[UserUpdateRequest](r.Body)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := c.Repository.Read(id)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	entity.Update(user, &body)

	err = c.Repository.Update(user)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	response.JSON(w, user, http.StatusOK)
	
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := c.Repository.Read(id)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.Repository.Delete(id)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		response.JSON(w, user, http.StatusOK)
	}
}
