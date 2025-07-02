package app

import (
	"log"
	"mail-phone-auth/internal/api"
	"mail-phone-auth/internal/app/files"
	"mail-phone-auth/pkg/http_server"
	"mail-phone-auth/pkg/postgres"
)

type App struct {
	API *api.API
	HttpServer *http_server.HttpServer
	Postgres *postgres.Postgres
}

func NewApp() *App {

	const httpServerConfigFilePath = "./config/http_server.json"
	const postgresCongigFilePath = "./config/postgres.json"

	app := App{}

	// Create Handler:
	app.API = api.New()

	// Create HttpServer:
	httpConfig, err:= files.InitConfig[http_server.Config](httpServerConfigFilePath)
	if err != nil {log.Fatal(err)}

	app.HttpServer =  http_server.New(httpConfig, app.API.Router)

	//Create Poatgres:
	postgresConfig, err := files.InitConfig[postgres.Config](postgresCongigFilePath)
	if err != nil {log.Fatal(err)}

	app.Postgres, err = postgres.NewPostgres(postgresConfig)
	if err != nil {log.Fatal(err)}

	return &app
}



func (app *App) Run() {
	go func() {
		err := app.HttpServer.Run()
		if err != nil {log.Fatal(err)}
	}()
}

func (app *App) Stop() {
	app.HttpServer.Stop()
}
