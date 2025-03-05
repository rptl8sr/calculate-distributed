package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	timeout = time.Second * 60
)

func New() *chi.Mux {
	r := chi.NewMux()
	r.Use(
		middleware.Timeout(timeout),
		middleware.Logger,
		middleware.Recoverer,
	)

	return r
}
