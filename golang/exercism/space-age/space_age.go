package space

import (
	"fmt"
	"strconv"
)

// Planet is a custom string type
type Planet string

const earthOrbit float64 = 31557600

// Age calculates the given number in the given planet year
func Age(seconds float64, planet Planet) float64 {
	planetOrbits := map[Planet]float64{
		"Earth":   1.0,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	rounded := fmt.Sprintf("%.2f", seconds/(earthOrbit*planetOrbits[planet]))
	result, _ := strconv.ParseFloat(rounded, 64)
	return result
}
