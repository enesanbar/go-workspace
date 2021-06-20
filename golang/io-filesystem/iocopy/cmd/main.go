package main

import (
	"bytes"
	"fmt"

	"github.com/enesanbar/workspace/golang/io-filesystem/interfaces/iocopy"
)

func main() {
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Print("stdout on Copy = ")
	if err := iocopy.Copy(in, out); err != nil {
		panic(err)
	}

	fmt.Println("out bytes buffer =", out.String())
}
