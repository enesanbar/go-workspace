package luhn

import (
	"strconv"
	"strings"
)

// Valid determines whether or not the input is valid per the Luhn formula.
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")

	length := len(input)
	if length <= 1 {
		return false
	}

	var count int
	shouldDouble := length%2 == 0
	for _, r := range input {
		value, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}

		if shouldDouble {
			value *= 2
			if value > 9 {
				value -= 9
			}
		}

		shouldDouble = !shouldDouble
		count += value
	}
	return count%10 == 0
}
