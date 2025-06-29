package app

import (
	"encoding/json"
	"log"
	"mail-phone-auth/pkg/http"
	"os"
)

type App struct {
	HttpServer  *http.Server
}

func New() *App {

	const httpServerConfigFilePath = "./config/http_server.json"
	
	httpConfig, err:= initConfig[http.Config](httpServerConfigFilePath)

	if err != nil {
		log.Fatal(err)
	}

	return &App{
		HttpServer: http.New(httpConfig),
	}	
}

func (app *App) Run() {
	go func() {
		app.HttpServer.Run()
	}()
}

func (app *App) Stop() {
	app.HttpServer.Stop()
}

func initConfig[T any](configFilePath string) (*T, error) {
	
	var config T

	fileData, err := os.ReadFile(configFilePath)

	if err != nil {
		return  nil, err
	}

	err = json.Unmarshal(fileData, &config)

	if err != nil {
		return  nil, err
	}

	return &config, nil
}

