package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pjol/opencred-go-server/routes"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/callback", routes.Callback)
	http.ListenAndServe(":9000", r)
}
