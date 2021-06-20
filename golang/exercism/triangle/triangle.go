// Package triangle implements utility methods for triangles
package triangle

import "math"

// Kind is our custom int type
// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	NaT Kind = iota // Not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides returns the Kind of a triangle.
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if !isTriangle(a, b, c) {
		return NaT
	}

	if a == b && a == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}

	return k
}

func isTriangle(a, b, c float64) bool {
	switch {
	case a+b <= c, a+c <= b, b+c <= a:
		return false
	case a <= 0, b <= 0, c <= 0:
		return false
	case math.IsNaN(a + b + c):
		return false
	default:
		return true
	}
}
