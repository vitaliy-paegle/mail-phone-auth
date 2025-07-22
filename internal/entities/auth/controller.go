package auth

import (
	"crypto/sha256"
	"fmt"
	"log"
	"mail-phone-auth/internal/api/request"
	"mail-phone-auth/internal/api/response"
	"mail-phone-auth/pkg/jino_mail"
	"net/http"
	"time"
)

type Controller struct {
	router     *http.ServeMux
	repository *Repository
	jinoMail   *jino_mail.JinoMail
}

func NewController(router *http.ServeMux, repository *Repository, jinoMail *jino_mail.JinoMail) *Controller {
	controller := Controller{
		router:     router,
		repository: repository,
		jinoMail:   jinoMail,
	}

	controller.router.HandleFunc("POST /api/auth/email/code", controller.EmailCode)
	controller.router.HandleFunc("POST /api/auth/email/confirm", controller.EmailConfirm)

	return &controller
}

func (c *Controller) EmailCode(w http.ResponseWriter, r *http.Request) {

	const timeOut = time.Duration(20) * time.Second

	body, err := request.DecodeBody[AuthEmailCodeRequest](r.Body)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authCopy := c.repository.ReadLastAuthByEmail(body.Email)

	if authCopy != nil {

		duration := time.Since(authCopy.CreatedAt)

		if duration < timeOut {

			waitingTime := (timeOut - duration).Round(time.Second).String()

			message := fmt.Sprintf("Повторная отправка кода возможна через: %s", waitingTime)
			response.Error(w, message, http.StatusInternalServerError)

			return
		}
	}

	code := RandomCode(4)
	hashSum := sha256.Sum256([]byte(code))
	hashString := fmt.Sprintf("%x", hashSum)

	auth := Auth{
		Email: body.Email,
		Code:  hashString,
	}

	auth.CreatedAt = time.Now()
	auth.UpdatedAt = time.Now()

	err = c.repository.CreateEmailAuth(&auth)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.jinoMail.SendCode(auth.Email, code)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.JSON(w, &auth, http.StatusOK)

}

func (controller *Controller) EmailConfirm(w http.ResponseWriter, r *http.Request) {

}
