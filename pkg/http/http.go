package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Server struct{
	Config Config
	HttpServer *http.Server
	Handler *http.ServeMux
}


func New(config *Config) *Server {

	const readTimeout = 5
	const writeTimeout = 5

	handler := http.NewServeMux()

	httpServer := http.Server{
		Addr: config.Host + ":" + config.Port,
		Handler: http.NewServeMux(),
		ReadTimeout: readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
	}
 
	return  &Server{
		Config: *config,
		HttpServer: &httpServer,
		Handler: handler,
	}
}

func (server *Server) Run() {
	log.Println("Server is listening port: " + server.Config.Port)
	err := server.HttpServer.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (server *Server) Stop() {
	server.HttpServer.Shutdown(context.Background())
}
