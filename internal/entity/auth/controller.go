package auth

import (
	"crypto/sha256"
	"fmt"
	"log"
	"mail-phone-auth/internal/api/request"
	"mail-phone-auth/internal/api/response"
	"mail-phone-auth/internal/entity/user"
	"mail-phone-auth/pkg/jino_mail"
	"mail-phone-auth/pkg/jwt"
	"net/http"
	"time"
)

type Controller struct {
	router     *http.ServeMux
	repository *Repository
	jinoMail   *jino_mail.JinoMail
	jwt        *jwt.JWT
}

func NewController(
	router *http.ServeMux,
	repository *Repository,
	jinoMail *jino_mail.JinoMail,
	jwt *jwt.JWT,
) *Controller {
	controller := Controller{
		router:     router,
		repository: repository,
		jinoMail:   jinoMail,
		jwt:        jwt,
	}

	controller.router.HandleFunc("POST /api/auth/email/code", controller.EmailCode)
	controller.router.HandleFunc("POST /api/auth/email/confirm", controller.EmailConfirm)
	controller.router.HandleFunc("POST /api/auth/refresh", controller.Refresh)

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

	response.ResponseStatus(w, http.StatusOK)

}

func (c *Controller) EmailConfirm(w http.ResponseWriter, r *http.Request) {

	const codeValidPeriod = time.Duration(60) * time.Second

	body, err := request.DecodeBody[AuthEmailConfirmRequest](r.Body)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	codeHashString := fmt.Sprintf("%x", sha256.Sum256([]byte(body.Code)))

	auth := c.repository.ReadLastAuthByEmail(body.Email)

	if auth == nil {
		response.Error(w, "error: Запрос на авторизацию не найден", http.StatusBadRequest)
		return
	}

	duration := time.Since(auth.CreatedAt)

	if duration > codeValidPeriod {

		message := fmt.Sprintf("error: Код авторизации устарел. Период действия: %s", codeValidPeriod)
		response.Error(w, message, http.StatusBadRequest)
		return
	}

	if codeHashString != auth.Code {
		response.Error(w, "error: Hеверный код подтверждения", http.StatusBadRequest)
		return
	}

	userData := user.User{}

	result := c.repository.postgres.DB.Table("users").
		Where("email = ?", body.Email).
		First(&userData)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {

			userData.Email = body.Email
			result = c.repository.postgres.DB.Create(&userData)

			if result.Error != nil {
				response.Error(w, result.Error.Error(), http.StatusInternalServerError)
				return
			}

		} else {
			response.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
	}

	tokensSet, err := c.jwt.CreateTokens(userData.ID)

	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
	}

	responseData := AuthJwtTokens{
		AccessToken:  tokensSet.Access,
		RefreshToken: tokensSet.Refresh,
	}

	c.repository.Delete(int(auth.ID))

	response.JSON(w, &responseData, http.StatusOK)
}

func (c *Controller) Refresh(w http.ResponseWriter, r *http.Request) {
	body, err := request.DecodeBody[AuthRefreshRequest](r.Body)

	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokensSet, err := c.jwt.UpdateTokens(body.RefreshToken)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseData := AuthJwtTokens{
		AccessToken:  tokensSet.Access,
		RefreshToken: tokensSet.Refresh,
	}

	response.JSON(w, &responseData, http.StatusOK)

}
