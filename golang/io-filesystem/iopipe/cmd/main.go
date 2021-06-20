package main

import (
	"fmt"

	"github.com/enesanbar/workspace/golang/io-filesystem/interfaces/iopipe"
)

func main() {
	fmt.Print("stdout on PipeExample = ")
	if err := iopipe.PipeExample(); err != nil {
		panic(err)
	}
}
