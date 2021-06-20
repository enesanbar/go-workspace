package main

import (
	"fmt"

	"github.com/enesanbar/workspace/golang/logging/panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}
