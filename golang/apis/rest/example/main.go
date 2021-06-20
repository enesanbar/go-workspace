package main

import "github.com/enesanbar/workspace/golang/apis/rest"

func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
