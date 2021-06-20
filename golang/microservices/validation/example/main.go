package main

import (
	"fmt"
	"net/http"

	"github.com/enesanbar/workspace/golang/microservices/validation"
)

// curl "http://localhost:3333/" -X POST -d '{"name":"test","age": 5}' -v
func main() {
	c := validation.New()
	http.HandleFunc("/", c.Process)
	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
