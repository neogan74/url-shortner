package config

import "os"

type Config struct {
	Port string
	Host string
}

func New() Config {
	return Config{
		Port: getPort(),
		Host: getHost(),
	}
}

func getHost() string {
	h := os.Getenv("HOST")
	if h == "" {
		h = "http://localhost"
	}
	return h
}

func getPort() string {
	p := os.Getenv("PORT")
	if p == "" {
		p = "8080"
	}
	return p
}
