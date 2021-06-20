package concealed_square

import (
	"fmt"
	"strconv"
)

func IsFormCorrect(input int) bool {
	squared := input * input
	str := strconv.Itoa(squared)
	fmt.Println(stringToBin(str))
	return false
}

func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}
