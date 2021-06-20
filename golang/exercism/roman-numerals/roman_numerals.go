package romannumerals

import (
	"strings"

	"github.com/pkg/errors"
)

type arabicToRoman struct {
	arabic int
	roman  string
}

var lookup = []arabicToRoman{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(num int) (string, error) {
	if num <= 0 || num > 3000 {
		return "", errors.New("Number not allowed")
	}

	var output strings.Builder

	for _, n := range lookup {
		for num >= n.arabic {
			output.WriteString(n.roman)
			num -= n.arabic
		}
	}

	return output.String(), nil
}
