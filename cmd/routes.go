package main

import (
	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func (app *Application) routes() http.Handler {

	mux := chi.NewRouter()
	prom := chiprometheus.NewMiddleware("Key-value-service")
	mux.Use(prom)
	mux.Use(middleware.DefaultLogger)
	mux.Handle("/metrics", promhttp.Handler())
	mux.Get("/api/get/{key}", app.get)
	mux.Post("/api/set", app.set)
	mux.Get("/api/search", app.search)
	return mux
}
