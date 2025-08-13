package app

import (
	"log"
	static "mail-phone-auth"
	"mail-phone-auth/internal/api"
	"mail-phone-auth/internal/app/files"
	"mail-phone-auth/internal/middleware"
	"mail-phone-auth/pkg/exolve"
	"mail-phone-auth/pkg/http_server"
	"mail-phone-auth/pkg/jino_mail"
	"mail-phone-auth/pkg/jwt"
	"mail-phone-auth/pkg/postgres"
	"net/http"
	"os"
)

type App struct {
	staticFileSystem *http.FileSystem
	api              *api.API
	httpServer       *http_server.HttpServer
	postgres         *postgres.Postgres
	jwt              *jwt.JWT
	jinoMail         *jino_mail.JinoMail
	exolve           *exolve.Exolve
	middleware *middleware.Middleware
}

func NewApp() *App {

	const httpServerConfigFilePath = "./config/http_server.json"
	const postgresCongigFilePath = "./config/postgres.json"
	const jwtCongigFilePath = "./config/jwt.json"
	const jinoMailFilePath = "./config/jino_mail.json"
	const exolveFilePath = "./config/exolve.json"
	const logFilePath = "./"

	app := App{}

	//  Set Logs file
	app.setLogFile(logFilePath)

	// Static File System:
	app.staticFileSystem = static.New()

	// JWT:
	jwtConfig, err := files.InitConfig[jwt.Config](jwtCongigFilePath)
	if err != nil {
		log.Fatal("JWT CONFIG ERROR: ", err)
	}

	app.jwt = jwt.New(jwtConfig)

	//JinoEmail:
	jinoMailConfig, err := files.InitConfig[jino_mail.Config](jinoMailFilePath)
	if err != nil {
		log.Fatal("JINO MAIL CONFIG ERROR: ", err)
	}

	app.jinoMail = jino_mail.New(jinoMailConfig, false)

	//Postgres:
	postgresConfig, err := files.InitConfig[postgres.Config](postgresCongigFilePath)
	if err != nil {
		log.Fatal("POSTGRES CONFIG ERROR: ", err)
	}

	app.postgres, err = postgres.NewPostgres(postgresConfig)
	if err != nil {
		log.Fatal("POSTGRES INIT ERROR: ", err)
	}

	//API:
	app.api = api.New(app.postgres, app.jwt, app.jinoMail)

	//Middleware:
	app.middleware = middleware.New(app.jwt)

	//HttpServer:
	httpConfig, err := files.InitConfig[http_server.Config](httpServerConfigFilePath)
	if err != nil {
		log.Fatal("HTTP SERVER CONFIG ERROR: ", err)
	}

	app.httpServer = http_server.New(httpConfig, app.api.Router, app.staticFileSystem, app.middleware)

	//Exolve:
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

func (app *App) setLogFile(filePath string) {

	_, err := os.Stat(filePath + "logs")

	if err != nil {
		err = nil
		err = os.Mkdir(filePath+"logs", 0755)
		if err != nil {
			log.Panic(err)
		}
	}

	logFile, err := os.Create(filePath + "logs/logs")

	if err != nil {
		log.Panic(err)
	}

	log.SetOutput(logFile)	

}
