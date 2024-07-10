package services

import (
	"fmt"

	"github.com/neogan74/url-shortner/internal/app/config"
	"github.com/neogan74/url-shortner/internal/app/services/generator"
	"github.com/neogan74/url-shortner/internal/app/storage"
)

type Shortner struct {
	repository storage.Repository
	generator  generator.Generator
	config     config.Config
}

func New(repository storage.Repository, generator generator.Generator, config config.Config) *Shortner {
	return &Shortner{
		repository: repository,
		generator:  generator,
		config:     config,
	}
}

func (service *Shortner) Shorten(url string) (string, error) {
	urlID, err := service.generator.GenerateIDFromString(url)
	if err != nil {
		return "", err
	}
	err = service.repository.Save(url, urlID)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%s:%s/%s", service.config.Host, service.config.Port, urlID)
	return result, nil
}

func (service *Shortner) Expand(id string) (string, error) {
	origURL, err := service.repository.GetByID(id)
	if err != nil {
		return "", err
	}
	return origURL, nil
}
