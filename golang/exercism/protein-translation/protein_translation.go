package protein

import (
	"errors"
)

var ErrStop = errors.New("stopping")
var ErrInvalidBase = errors.New("invalid base")

const Stop string = "STOP"

var codonMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": Stop,
	"UAG": Stop,
	"UGA": Stop,
}

func FromCodon(codon string) (string, error) {
	amino, ok := codonMap[codon]

	if !ok {
		return "", ErrInvalidBase
	}

	if amino == Stop {
		return "", ErrStop
	}

	return amino, nil
}

func FromRNA(rna string) ([]string, error) {
	var result []string
	for i := 0; i < len(rna); i += 3 {
		amino, err := FromCodon(rna[i : i+3])

		switch {
		case err == ErrInvalidBase:
			return result, err
		case err == ErrStop:
			return result, nil
		default:
			result = append(result, amino)
		}
	}
	return result, nil
}
