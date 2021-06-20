// Package proverb provides utility functions to create a rhyme
package proverb

import "fmt"

const (
	stanza = "For want of a %s the %s was lost."
	last   = "And all for the want of a %s."
)

// Proverb returns the relevant proverb given a list of inputs
func Proverb(rhyme []string) []string {
	length := len(rhyme)

	if length == 0 {
		return []string{}
	}

	var proverbial []string

	for i := 0; i < length-1; i++ {
		proverbial = append(proverbial, fmt.Sprintf(stanza, rhyme[i], rhyme[i+1]))
	}

	proverbial = append(proverbial, fmt.Sprintf(last, rhyme[0]))

	return proverbial
}
