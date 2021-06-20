package main

import (
	"log"
	"net/http"

	"github.com/enesanbar/workspace/golang/projects/websockets-chat/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("starting channel listener")
	go handlers.ListenToWSChannel()

	log.Println("starting web server on 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Println(err)
	}

}
