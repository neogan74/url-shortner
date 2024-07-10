package main

import (
	"github.com/neogan74/url-shortner/internal/app/config"
	"github.com/neogan74/url-shortner/internal/app/server"
	"github.com/neogan74/url-shortner/internal/app/services"
	"github.com/neogan74/url-shortner/internal/app/services/generator"
	"github.com/neogan74/url-shortner/internal/app/storage"
)

func main() {
	cfg := config.New()
	repo := storage.NewMemStorage()
	gen := &generator.HashGenerator{}
	service := services.New(repo, gen, cfg)
	srv := server.New(cfg, service)

	srv.Run()
}
