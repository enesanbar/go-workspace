package evaluation

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"

	"github.com/gonum/stat"
)

// EvaluateContinuousData prints MSE, MAE, R-Squared
func EvaluateContinuousData(b io.Reader) {
	// Create a new CSV reader reading from io.Reader.
	reader := csv.NewReader(b)

	// observed and predicted will hold the parsed observed and
	// predicted values from the continuous data file.
	var observed []float64
	var predicted []float64

	// line will track row numbers for logging.
	line := 1

	// Read in the records looking for unexpected types in the columns.
	for {
		// Read in a row. Check if we are at the end of the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Skip the header.
		if line == 1 {
			line++
			continue
		}

		// Read in the observed and predicted values.
		observedVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// Append the record to our slice, if it has the expected type.
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	// Both MSE and MAE give us a good overall picture of how good our predictions are
	// Calculate the mean absolute error and mean squared error.
	var mAE float64
	var mSE float64
	for idx, observedVal := range observed {
		mAE += math.Abs(observedVal-predicted[idx]) / float64(len(observed))
		mSE += math.Pow(observedVal-predicted[idx], 2) / float64(len(observed))
	}

	// Output the MAE and MSE value to standard out.
	fmt.Printf("MAE = %0.2f\n", mAE)
	fmt.Printf("MSE = %0.2f\n", mSE)

	// Calculate the R^2 value.
	rSquared := stat.RSquaredFrom(observed, predicted, nil)

	// Output the R^2 value to standard out.
	fmt.Printf("R^2 = %0.2f\n\n", rSquared)
}
