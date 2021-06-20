package main

import (
	"fmt"
	"math"

	concealedSquare "github.com/enesanbar/workspace/golang/projecteuler/206-concealed-square"
)

const (
	UpperBound = 1929394959697989990
	LowerBound = 1020304050607080900
)

func main() {
	upper := int(math.Floor(math.Sqrt(UpperBound)))
	lower := int(math.Floor(math.Sqrt(LowerBound)))

	for i := lower; i <= upper; i++ {
		if concealedSquare.IsFormCorrect(i) {
			fmt.Println(i)
			return
		}
	}
}
