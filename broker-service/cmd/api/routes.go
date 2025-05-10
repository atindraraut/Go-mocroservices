package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config)routes() http.Handler{
	mux:=chi.NewRouter();
	//specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored for preflighted requests
	}))
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Get("/v1/broker", app.Broker)
	return mux
}