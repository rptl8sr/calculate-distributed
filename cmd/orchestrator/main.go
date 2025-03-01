package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"calculate-distributed/internal/api"
	"calculate-distributed/internal/server"
)

func main() {
	si := server.New()

	r := chi.NewMux()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := api.HandlerFromMux(si, r)

	s := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}

	fmt.Println("Starting server on port 8080")
	if err := s.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
