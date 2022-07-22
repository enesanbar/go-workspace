package main

import (
	"fmt"

	"github.com/enesanbar/workspace/golang/error-handling/basicerrors"
)

func main() {
	basicerrors.BasicErrors()

	err := basicerrors.SomeFunc()
	fmt.Println("custom error: ", err)
}
