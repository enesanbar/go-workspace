package main

import "github.com/enesanbar/workspace/golang/io-filesystem/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
