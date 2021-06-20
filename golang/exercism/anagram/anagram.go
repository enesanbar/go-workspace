package anagram

import (
	"sort"
	"strings"
)

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

// Detect returns anagrams of given input from the given candidates
func Detect(input string, candidates []string) (result []string) {
	lowerInput := strings.ToLower(input)

	r1 := ByRune(lowerInput)
	sort.Sort(r1)

	for _, c := range candidates {
		lowerCandidate := strings.ToLower(c)
		if len(lowerCandidate) != len(lowerInput) && lowerCandidate == lowerInput {
			continue
		}

		r2 := ByRune(lowerCandidate)
		sort.Sort(r2)

		if string(r1) == string(r2) {
			result = append(result, c)
		}
	}
	return
}
