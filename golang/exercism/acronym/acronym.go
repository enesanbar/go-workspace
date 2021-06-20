// Package acronym contains a method to abbreviate a sentence
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate return the abbreviation of the given string.
func Abbreviate(s string) string {
	splited := regexp.MustCompile("[-,_ ]+").Split(s, -1)

	var acronym string
	for _, word := range splited {
		acronym += string(word[0])
	}

	return strings.ToUpper(acronym)
}
