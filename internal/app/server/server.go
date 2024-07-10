package server

import (
	"log"
	"net/http"

	"github.com/neogan74/url-shortner/internal/app/config"
	"github.com/neogan74/url-shortner/internal/app/handlers"
	"github.com/neogan74/url-shortner/internal/app/services"
)

type Server struct {
	config  config.Config
	service *services.Shortner
}

func (s *Server) Run() {
	r := handlers.NewRouter(s.service)

	httpServer := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: r,
	}
	log.Fatal(httpServer.ListenAndServe())
}

func New(config config.Config, service *services.Shortner) *Server {
	return &Server{
		config:  config,
		service: service,
	}
}
