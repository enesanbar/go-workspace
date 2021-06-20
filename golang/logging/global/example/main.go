package main

import "github.com/enesanbar/workspace/golang/logging/global"

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
