package pangram

// IsPangram determines if the given string contains every letter
func IsPangram(input string) bool {
	var seen int32

	for _, c := range input {
		if i := getCharIndex(c); i >= 0 && i <= 25 {
			seen |= 1 << i // sets 1 in the ith position
		}
	}

	// check if last 26 bits are all 1
	return seen == 0x3ffffff
}

// getCharIndex returns the position of the character in the alphabet, e.g. a = 0
func getCharIndex(c int32) int32 {
	return c&0xdf - 'A'
}
