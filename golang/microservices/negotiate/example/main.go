package main

import (
	"fmt"
	"net/http"

	"github.com/enesanbar/workspace/golang/microservices/negotiate"
)

func main() {
	http.HandleFunc("/", negotiate.Handler)
	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
