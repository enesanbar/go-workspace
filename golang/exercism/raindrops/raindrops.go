// Package raindrops implements Convert function
package raindrops

import (
	"bytes"
	"strconv"
)

var translations = []struct {
	num  int
	word string
}{
	{3, `Pling`},
	{5, `Plang`},
	{7, `Plong`},
}

// Convert returns a Pling/Plang/Plong depending on the passed number
func Convert(n int) string {
	var b bytes.Buffer

	for _, trans := range translations {
		if n%trans.num == 0 {
			b.WriteString(trans.word)
		}
	}

	if b.String() == "" {
		b.WriteString(strconv.Itoa(n))
	}

	return b.String()
}
