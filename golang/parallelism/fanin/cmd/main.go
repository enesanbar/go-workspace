package main

import (
	"fmt"
	"github.com/enesanbar/workspace/golang/parallelism/fanin"
	"time"

)

func main() {
	sources := make([]<-chan int, 0)

	for i := 0; i < 3; i++ {
		ch := make(chan int)
		sources = append(sources, ch)

		go func() {
			defer close(ch)

			for i := 0; i < 5; i++ {
				ch <- i
				time.Sleep(time.Second)
			}
		}()
	}

	dest := fanin.Funnel(sources...)
	for d := range dest {
		fmt.Println(d)
	}

}


