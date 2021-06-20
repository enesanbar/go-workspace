package main

import (
	mongodb "github.com/enesanbar/workspace/golang/distributed/orchestrate"
)

func main() {
	if err := mongodb.Exec("mongodb://mongodb:27017"); err != nil {
		panic(err)
	}
}
