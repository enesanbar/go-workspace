package main

import (
	"github.com/enesanbar/workspace/golang/io-filesystem/encoding/base64"
)

func main() {
	if err := base64.Example(); err != nil {
		panic(err)
	}

	if err := base64.Base64ExampleEncoder(); err != nil {
		panic(err)
	}

	if err := base64.GobExample(); err != nil {
		panic(err)
	}
}
