package strand

var complements = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

// ToRNA returns RNA complement of a given a DNA strand
func ToRNA(dna string) string {
	var result string

	for _, nuc := range dna {
		result += complements[nuc]
	}

	return result
}
