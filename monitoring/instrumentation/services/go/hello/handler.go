package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	time.Sleep(2 * time.Second)
	vars := mux.Vars(request)
	name := vars["name"]
	greeting := fmt.Sprintf("Hello %s", name)
	writer.Write([]byte(greeting))
}
