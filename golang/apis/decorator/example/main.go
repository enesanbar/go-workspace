package main

import "github.com/enesanbar/workspace/golang/apis/decorator"

func main() {
	if err := decorator.Exec(); err != nil {
		panic(err)
	}
}
