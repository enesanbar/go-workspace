// Package reverse implements utility function to reverse a string
package reverse

import (
	"strings"
)

// Reverse returns reversed version of a string
func Reverse(input string) (result string) {
	var output strings.Builder

	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		output.WriteRune(runes[i])
	}

	return output.String()
}
