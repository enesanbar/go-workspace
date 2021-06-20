package main

import (
	"fmt"

	"github.com/enesanbar/workspace/golang/logging/structured"
)

func main() {
	fmt.Println("Logrus:")
	structured.Logrus()

	fmt.Println()
	fmt.Println("Apex:")
	structured.Apex()
}
