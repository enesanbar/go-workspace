package srp

import (
	"fmt"
	"io"
)

// CalculatorBad calculates the test coverage for a directory and it's sub-directories
type CalculatorBad struct {
	// coverage data populated by `Calculate()` method
	data map[string]float64
}

// Calculate will calculate the coverage
func (c *CalculatorBad) Calculate(_ string) error {
	// run `go test -cover ./[path]/...` and store the results
	return nil
}

// Output will print the coverage data to the supplied writer
func (c CalculatorBad) Output(writer io.Writer) {
	for path, result := range c.data {
		fmt.Fprintf(writer, "%s -> %.1f\n", path, result)
	}
}

// OutputCSV will print the coverage data to the supplied writer
func (c CalculatorBad) OutputCSV(writer io.Writer) {
	for path, result := range c.data {
		fmt.Fprintf(writer, "%s,%.1f\n", path, result)
	}
}
