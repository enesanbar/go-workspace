package evaluation

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/gonum/stat"
	"gonum.org/v1/gonum/integrate"
)

func EvaluateAUC() {

	// Define our scores and classes.
	scores := []float64{0.1, 0.35, 0.4, 0.8}
	classes := []bool{true, false, true, false}

	// Calculate the true positive rates (recalls) and false positive rates.
	tpr, fpr := stat.ROC(0, scores, classes, nil)

	// Compute the Area Under Curve.
	auc := integrate.Trapezoidal(fpr, tpr)

	// Output the results to standard out.
	fmt.Printf("true  positive rate: %v\n", tpr)
	fmt.Printf("false positive rate: %v\n", fpr)
	fmt.Printf("auc: %v\n", auc)
}

func EvaluateCategoricalData(b io.Reader) {
	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(b)

	// observed and predicted will hold the parsed observed and predicted values
	observed, predicted := readColumns(reader)

	// This variable will hold our count of true positive and true negative values.
	var truePosNeg int

	// Accumulate the true positive/negative count.
	for idx, observedVal := range observed {
		if observedVal == predicted[idx] {
			truePosNeg++
		}
	}

	// Calculate the accuracy (subset accuracy).
	// The percentage of predictions that were right
	accuracy := float64(truePosNeg) / float64(len(observed))

	// Output the Accuracy value to standard out.
	fmt.Printf("\nAccuracy = %0.2f\n\n", accuracy)
}

func EvaluateCategoricalDataWithClasses(b io.Reader) {
	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(b)

	// observed and predicted will hold the parsed observed and predicted values
	observed, predicted := readColumns(reader)

	// classes contains the three possible classes in the labeled data.
	classes := []int{0, 1, 2}

	// Loop over each class.
	for _, class := range classes {

		// These variables will hold our count of true positives and
		// our count of false positives.
		var truePos int
		var falsePos int
		var falseNeg int

		// Accumulate the true positive and false positive counts.
		for idx, oVal := range observed {

			switch oVal {

			// If the observed value is the relevant class, we should check to
			// see if we predicted that class.
			case class:
				if predicted[idx] == class {
					truePos++
					continue
				}

				falseNeg++

			// If the observed value is a different class, we should check to see
			// if we predicted a false positive.
			default:
				if predicted[idx] == class {
					falsePos++
				}
			}
		}

		// Calculate the precision.
		precision := float64(truePos) / float64(truePos+falsePos)

		// Calculate the recall.
		recall := float64(truePos) / float64(truePos+falseNeg)

		// Output the precision value to standard out.
		fmt.Printf("Precision (class %d) = %0.2f\n", class, precision)
		fmt.Printf("Recall (class %d) = %0.2f\n", class, recall)
	}
}

func readColumns(reader *csv.Reader) ([]int, []int) {
	var observed []int
	var predicted []int

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
		observedVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.Atoi(record[1])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// Append the record to our slice, if it has the expected type.
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}
	return observed, predicted
}
