// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

type Remark string

// Hey should have a comment documenting it.
func Hey(remark string) string {
	r := Remark(strings.TrimSpace(remark))

	switch {
	case r.isQuestion() && r.isYelling():
		return "Calm down, I know what I'm doing!"
	case r.isQuestion():
		return "Sure."
	case r.isYelling():
		return "Whoa, chill out!"
	case r.isSilent():
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}

func (r Remark) isSilent() bool {
	return r == ""
}

func (r Remark) isQuestion() bool {
	return strings.HasSuffix(string(r), "?")
}

func (r Remark) isYelling() bool {
	yelling := false

	for _, char := range r {
		if !unicode.IsLetter(char) {
			continue
		}

		if unicode.IsUpper(char) {
			yelling = true
			continue
		}

		yelling = false
		break
	}

	return yelling
}
