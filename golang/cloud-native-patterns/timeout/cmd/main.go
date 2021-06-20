package main

import (
	"cloud-native-patterns/timeout"
	"context"
	"fmt"
	"time"
)

func Slow(s string) (string, error) {
	time.Sleep(time.Second * 2)
	return "", nil
}

func main() {
	ctx := context.Background()
	ctxt, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	to := timeout.Timeout(Slow)
	res, err := to(ctxt, "some input")

	fmt.Println(res, err)
}
