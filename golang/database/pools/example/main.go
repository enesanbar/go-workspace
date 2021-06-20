package main

import "github.com/enesanbar/workspace/golang/database/pools"

func main() {
	if err := pools.ExecWithTimeout(); err != nil {
		panic(err)
	}
}
