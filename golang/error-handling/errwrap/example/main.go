package main

import (
	"fmt"

	"github.com/enesanbar/workspace/golang/error-handling/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	errwrap.Unwrap()
	fmt.Println()
	errwrap.StackTrace()
}
