package main

import (
	"net/http"

	"github.com/enesanbar/workspace/golang/projects/websockets-chat/internal/handlers"

	"github.com/bmizerany/pat"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WebsocketEndpoint))

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
