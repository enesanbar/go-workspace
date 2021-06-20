package main

import (
	"github.com/enesanbar/workspace/golang/database/storageinterface"
)

func main() {
	if err := storageinterface.Exec(); err != nil {
		panic(err)
	}
}
