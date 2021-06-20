package main

import (
	"fmt"
	"os"
	"which"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}
	file := arguments[1]
	fullPath, err := which.Which(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("executable found at", fullPath)
}
