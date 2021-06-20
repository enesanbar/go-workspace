package names_scores

import (
	"unicode"
)

func CalculateNameScore(names []string) int {
	var totalScore int

	for idx, name := range names {
		var score int
		for _, c := range name {
			//fmt.Printf("(%T) %c %d\n", c, c, c)
			score += int(unicode.ToUpper(c)) - 'A' + 1
		}
		totalScore += score * (idx + 1)
	}

	return totalScore
}
