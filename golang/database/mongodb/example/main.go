package main

import "github.com/enesanbar/workspace/golang/database/mongodb"

func main() {
	if err := mongodb.Exec("mongodb://localhost"); err != nil {
		panic(err)
	}
}
