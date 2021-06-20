package main

import (
	"fmt"

	"github.com/enesanbar/workspace/golang/logging/std"
)

func main() {
	fmt.Println("basic logging and modification of logger:")
	std.Log()
	fmt.Println("logging 'handled' errors:")
	std.FinalDestination()
}
