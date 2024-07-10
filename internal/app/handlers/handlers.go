package handlers

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/neogan74/url-shortner/internal/app/services"
)

func NewRouter(service *services.Shortner) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	h := NewHander(service)

	r.Route("/", func(r chi.Router) {
		r.Get("/{id}", h.Expand)
		r.Post("/", h.Shorten)
	})
	return r
}

type Handler struct {
	Mux     *chi.Mux
	service *services.Shortner
}

func NewHander(service *services.Shortner) *Handler {
	return &Handler{
		Mux:     chi.NewMux(),
		service: service,
	}
}

func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	url, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if string(url) == "" {
		http.Error(w, "url required", http.StatusBadRequest)
		return
	}

	su, err := h.service.Shorten(string(url))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(su))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Expand(w http.ResponseWriter, r *http.Request) {
	uID := chi.URLParam(r, "id")

	fullUrl, err := h.service.Expand(uID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if fullUrl == "" {
		http.Error(w, "can't find full URL", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, fullUrl, http.StatusTemporaryRedirect)
}
