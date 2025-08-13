package file

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"mail-phone-auth/internal/api/response"
	"mail-phone-auth/internal/entity"
	"mail-phone-auth/pkg/jwt"
	"mail-phone-auth/pkg/postgres"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

type Controller struct {
	Router    *http.ServeMux
	Postgres  *postgres.Postgres
	Entity    *entity.Entity[File, FileData]
	StorePath string
	JWT       *jwt.JWT
}

func NewController(router *http.ServeMux, postgres *postgres.Postgres, jwt *jwt.JWT) *Controller {

	controller := Controller{
		Router:   router,
		Postgres: postgres,
		JWT:      jwt,
	}

	controller.Entity = entity.New[File, FileData](controller.Postgres)
	controller.StorePath = "./files_store"

	controller.CreateStore()

	controller.Router.HandleFunc("POST /api/file", controller.Create)
	controller.Router.HandleFunc("GET /api/file/{id}", controller.Entity.Read)
	controller.Router.HandleFunc("GET /api/file/all", controller.Entity.ReadAll)
	controller.Router.HandleFunc("DELETE /api/file/{id}", controller.Entity.Delete)
	controller.Router.HandleFunc("GET /file/{link}", controller.Download)

	return &controller
}

func (c *Controller) CreateStore() {
	_, err := os.Stat(c.StorePath)
	if err != nil {
		err = nil
		err = os.Mkdir(c.StorePath, 0755)
		if err != nil {
			log.Panic(err)
		}
	}
}

func FileHash(file multipart.File) (string, error) {

	hash := sha256.New()
	_, err := io.Copy(hash, file)

	if err != nil {
		return "", err
	}

	hashInBytes := hash.Sum(nil)

	return hex.EncodeToString(hashInBytes), nil
}

func (c *Controller) GetUserID(r *http.Request) (*uint, error) {

	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return nil, errors.New("auth header empty")
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {

		return nil, errors.New("authorization header must be a Bearer token")
	}

	tokenString := authHeader[len(prefix):]

	tokenData, err := c.JWT.ParseToken(tokenString)

	if err != nil {
		return nil, err
	}

	var userID uint = tokenData.UserID
	return &userID, nil

}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {

	userID, err := c.GetUserID(r)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = r.ParseMultipartForm(32 << 20)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	hash, err := FileHash(file)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = file.Seek(0, 0)

	if err != nil {
		log.Println("Failed to seek file:", err)
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(header.Filename, hash)

	fileExt := strings.Split(filepath.Ext(header.Filename), ".")[1]

	fileName := hash + "." + fileExt

	var copy File

	result := c.Postgres.Table("files").Where("hash = ?", hash).Where("deleted_at is NULL").First(&copy)

	c.Entity.ReadRelatedData(reflect.ValueOf(&copy))

	if result.Error == nil {
		response.JSON(w, &copy, http.StatusOK)
		return
	}

	f, err := os.Create(c.StorePath + "/" + fileName)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer f.Close()

	written, err := io.Copy(f, file)

	log.Println("Written: ", written)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var fileData File = File{}

	fileData.CreatedAt = time.Now()
	fileData.UpdatedAt = time.Now()
	fileData.Name = strings.TrimSuffix(header.Filename, filepath.Ext(header.Filename))
	fileData.Extension = &fileExt
	fileData.Hash = hash
	fileData.Link = "/file/" + fileName
	fileData.UserID = *userID

	// json, _ := json.Marshal(fileData)

	// log.Println(string(json))

	result = c.Postgres.DB.Create(&fileData)

	if result.Error != nil {
		log.Println(result.Error.Error())
		response.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	c.Entity.ReadRelatedData(reflect.ValueOf(&fileData))

	response.JSON(w, &fileData, http.StatusOK)

}

func (c *Controller) Download(w http.ResponseWriter, r *http.Request) {
	link := r.PathValue("link")

	var filePath string = c.StorePath + "/" + link

	log.Println(filePath)

	http.ServeFile(w, r, filePath)

}
