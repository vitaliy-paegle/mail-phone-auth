package app

import (
	"log"
	static "mail-phone-auth"
	"mail-phone-auth/internal/api"
	"mail-phone-auth/internal/app/files"
	"mail-phone-auth/pkg/exolve"
	"mail-phone-auth/pkg/http_server"
	"mail-phone-auth/pkg/jino_mail"
	"mail-phone-auth/pkg/jwt"
	"mail-phone-auth/pkg/postgres"
	"net/http"
)

type App struct {
	staticFileSystem *http.FileSystem
	api              *api.API
	httpServer       *http_server.HttpServer
	postgres         *postgres.Postgres
	jwt              *jwt.JWT
	jinoMail         *jino_mail.JinoMail
	exolve           *exolve.Exolve
}

func NewApp() *App {

	const httpServerConfigFilePath = "./config/http_server.json"
	const postgresCongigFilePath = "./config/postgres.json"
	const jwtCongigFilePath = "./config/jwt.json"
	const jinoMailFilePath = "./config/jino_mail.json"
	const exolveFilePath = "./config/exolve.json"

	app := App{}

	// Create Static File System

	app.staticFileSystem = static.New()

	// Create JWT:
	jwtConfig, err := files.InitConfig[jwt.Config](jwtCongigFilePath)
	if err != nil {
		log.Fatal("JWT CONFIG ERROR: ", err)
	}

	app.jwt = jwt.New(jwtConfig)

	// Create JinoEmail:
	jinoMailConfig, err := files.InitConfig[jino_mail.Config](jinoMailFilePath)
	if err != nil {
		log.Fatal("JINO MAIL CONFIG ERROR: ", err)
	}

	app.jinoMail = jino_mail.New(jinoMailConfig, false)

	// Create Poatgres:
	postgresConfig, err := files.InitConfig[postgres.Config](postgresCongigFilePath)
	if err != nil {
		log.Fatal("POSTGRES CONFIG ERROR: ", err)
	}

	app.postgres, err = postgres.NewPostgres(postgresConfig)
	if err != nil {
		log.Fatal("POSTGRES INIT ERROR: ", err)
	}

	// Create API:
	app.api = api.New(app.postgres, app.jwt, app.jinoMail)

	// Create HttpServer:
	httpConfig, err := files.InitConfig[http_server.Config](httpServerConfigFilePath)
	if err != nil {
		log.Fatal("HTTP SERVER CONFIG ERROR: ", err)
	}

	app.httpServer = http_server.New(httpConfig, app.api.Router, app.staticFileSystem)

	// Create Exolve:
	exolveConfig, err := files.InitConfig[exolve.Config](exolveFilePath)
	if err != nil {
		log.Fatal("EXOLVE CONFIG ERROR: ", err)
	}

	app.exolve = exolve.New(exolveConfig, false)

	return &app
}

func (app *App) Run() {
	go func() {
		err := app.httpServer.Run()
		if err != nil {
			log.Println(err)
		}
	}()
}

func (app *App) Stop() {
	app.httpServer.Stop()
}
