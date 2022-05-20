package main

import (
	"github.com/Nguyenthaiduc/go-web/pkg/config"
	"github.com/Nguyenthaiduc/go-web/pkg/handlers"
	"github.com/bmizerany/pat.com"
	"net/http"
)
func routes(app *config.AppConfig) http.Handler {
	mux := pat.New() //mux mean multiplexer?

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
