// Package isogram provides utility method IsIsogram
package isogram

import (
	"unicode"
)

// IsIsogram determis if a given string is an isogram
func IsIsogram(word string) bool {
	seen := make(map[rune]bool)

	for _, c := range word {
		if !unicode.IsLetter(c) {
			continue
		}

		if seen[unicode.ToLower(c)] {
			return false
		}

		seen[unicode.ToLower(c)] = true
	}

	return true
}
