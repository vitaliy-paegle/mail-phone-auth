package http_server

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Config struct {
	Port string `json:"port"  validate:"required"`
}

// http_server.json
// {
// 	"port": "1000"
// }

type HttpServer struct {
	Config *Config
	Server *http.Server
	Router *http.ServeMux
}

func New(config *Config, router *http.ServeMux, staticFileSystem *http.FileSystem) *HttpServer {

	const readTimeout = 5
	const writeTimeout = 5

	server := http.Server{
		Addr:         ":" + config.Port,
		Handler:      router,
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
	}

	httpServer := HttpServer{
		Config: config,
		Server: &server,
		Router: router,
	}

	httpServer.Router.Handle("/static/", http.FileServer(*staticFileSystem))
	// httpServer.Router.Handle("/data/", http.StripPrefix("/data",  http.FileServer(http.Dir("./static"))))

	return &httpServer
}

func (httpServer *HttpServer) Run() error {
	log.Println("Server is listening port: " + httpServer.Config.Port)
	err := httpServer.Server.ListenAndServe()
	if err != nil {
		return err
	} else {
		return nil
	}

}

func (httpServer *HttpServer) Stop() {
	httpServer.Server.Shutdown(context.Background())
}
