package etl

import "strings"

// Transform convert a legacy scrabble point system into new system.
func Transform(scores map[int][]string) map[string]int {
	result := make(map[string]int, 26)
	for score, letters := range scores {
		for _, letter := range letters {
			result[strings.ToLower(letter)] = score
		}
	}

	return result
}
