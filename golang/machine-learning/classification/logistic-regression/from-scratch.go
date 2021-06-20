package logistic_regression

import (
	"fmt"
	"image/color"
	"io"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/enesanbar/workspace/golang/machine-learning/data"
	"github.com/gonum/matrix/mat64"
	"github.com/kniren/gota/dataframe"
	"github.com/kniren/gota/series"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func PlotLogisticFunction() {
	// Create a new plot.
	p := plot.New()
	p.Title.Text = "Logistic Function"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "f(x)"

	// Create the plotter function.
	logisticPlotter := plotter.NewFunction(func(x float64) float64 { return logistic(x) })
	logisticPlotter.Color = color.RGBA{B: 255, A: 255}

	// Add the plotter function to the plot.
	p.Add(logisticPlotter)

	// Set the axis ranges.
	// Unlike other data sets, functions don't set the axis ranges automatically
	// since functions don't necessarily have a finite range of x and y values.
	p.X.Min = -10
	p.X.Max = 10
	p.Y.Min = -0.1
	p.Y.Max = 1.1

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "./images/logistic.png"); err != nil {
		log.Fatal(err)
	}
}

// logistic implements the logistic function,
// which is used in logistic regression.
func logistic(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func cleanLoanData(b io.Reader) dataframe.DataFrame {
	csvDataframe := dataframe.ReadCSV(b)

	cleanedDataframe := csvDataframe.Copy()
	// Sequentially move the rows writing out the parsed values.
	for i, record := range csvDataframe.Records() {
		// Skip the header row.
		if i == 0 {
			cleanedDataframe.Set(series.Ints(i), dataframe.LoadRecords([][]string{{"FICO_score", "class"}}))
			continue
		}

		// Initialize a slice to hold our parsed values.
		outRecord := make([]string, 2)

		// Parse and normalize the FICO score.
		score, err := strconv.ParseFloat(strings.Split(record[0], "-")[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		outRecord[0] = strconv.FormatFloat((score-640.0)/(830.0-640.0), 'f', 4, 64)

		// Parse the Interest rate class.
		rate, err := strconv.ParseFloat(strings.TrimSuffix(record[1], "%"), 64)
		if err != nil {
			log.Fatal(err)
		}

		if rate <= 12.0 {
			outRecord[1] = "1.0"
			cleanedDataframe.Set(series.Ints(i), dataframe.LoadRecords([][]string{outRecord}))
			continue
		}

		outRecord[1] = "0.0"
		cleanedDataframe.Set(series.Ints(i), dataframe.LoadRecords([][]string{outRecord}))
	}

	return cleanedDataframe
}

func PlotHistogramWithCleanedData(b io.Reader) {
	// Create a cleaned dataframe from the CSV file.
	loanDF := dataframe.ReadCSV(b)

	// Use the Describe method to calculate summary statistics
	// for all of the columns in one shot.
	loanSummary := loanDF.Describe()

	// Output the summary statistics to stdout.
	fmt.Println(loanSummary)

	// Create a histogram for each of the columns in the dataset.
	for _, colName := range loanDF.Names() {

		// Create a plotter.Values value and fill it with the
		// values from the respective column of the dataframe.
		plotVals := make(plotter.Values, loanDF.Nrow())
		for i, floatVal := range loanDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		// Make a plot and set its title.
		p := plot.New()
		p.Title.Text = fmt.Sprintf("Histogram of %s", colName)

		// Create a histogram of our values.
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}

		// Normalize the histogram.
		h.Normalize(1)

		// Add the histogram to the plot.
		p.Add(h)

		// Save the plot to a PNG file.
		if err := p.Save(4*vg.Inch, 4*vg.Inch, "./images/"+colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}
}

func TrainModel(b io.Reader) {
	training, _ := data.SplitTrainingAndTest(b)

	// featureData and labels will hold all the float values that
	// will eventually be used in our training.
	featureData := make([]float64, 2*(len(training.Records())-1))
	labels := make([]float64, len(training.Records())-1)

	// featureIndex will track the current index of the features matrix values.
	var featureIndex int

	// Sequentially move the rows into the slices of floats.
	for idx, record := range training.Records() {
		// Skip the header row.
		if idx == 0 {
			continue
		}

		// Add the FICO score feature.
		featureVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		featureData[featureIndex] = featureVal

		// Add an intercept.
		featureData[featureIndex+1] = 1.0

		// Increment our feature row.
		featureIndex += 2

		// Add the class label.
		labelVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		labels[idx-1] = labelVal
	}

	// Form a matrix from the features.
	features := mat64.NewDense(len(training.Records())-1, 2, featureData)

	// Train the logistic regression model.
	weights := logisticRegression(features, labels, 1000, 0.3)

	// Output the Logistic Regression model formula to stdout.
	formula := "p = 1 / ( 1 + exp(- m1 * FICO.score - m2) )"
	fmt.Printf("\n%s\n\nm1 = %0.2f\nm2 = %0.2f\n\n", formula, weights[0], weights[1])
}

// logisticRegression fits a logistic regression model for the given data.
func logisticRegression(features *mat64.Dense, labels []float64, numSteps int, learningRate float64) []float64 {

	// Initialize random weights.
	_, numWeights := features.Dims()
	weights := make([]float64, numWeights)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for idx := range weights {
		weights[idx] = r.Float64()
	}

	// Iteratively optimize the weights.
	for i := 0; i < numSteps; i++ {

		// Initialize a variable to accumulate error for this iteration.
		var sumError float64

		// Make predictions for each label and accumulate error.
		for idx, label := range labels {

			// Get the features corresponding to this label.
			featureRow := mat64.Row(nil, idx, features)

			// Calculate the error for this iteration's weights.
			pred := logistic(featureRow[0]*weights[0] + featureRow[1]*weights[1])
			predError := label - pred
			sumError += math.Pow(predError, 2)

			// Update the feature weights.
			for j := 0; j < len(featureRow); j++ {
				weights[j] += learningRate * predError * pred * (1 - pred) * featureRow[j]
			}
		}
	}

	return weights
}

func TestTrainedData(b io.Reader) {
	_, testData := data.SplitTrainingAndTest(b)

	// observed and predicted will hold the parsed observed and predicted values
	// form the labeled data file.
	var observed []float64
	var predicted []float64

	for idx, record := range testData.Records() {
		// Skip the header.
		if idx == 0 {
			continue
		}

		// Read in the observed value.
		observedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", idx)
			continue
		}

		// Make the corresponding prediction.
		score, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", idx)
			continue
		}

		predictedVal := predict(score)

		// Append the record to our slice, if it has the expected type.
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
	}

	// This variable will hold our count of true positive and
	// true negative values.
	var truePosNeg int

	// Accumulate the true positive/negative count.
	for idx, oVal := range observed {
		if oVal == predicted[idx] {
			truePosNeg++
		}
	}

	// Calculate the accuracy (subset accuracy).
	accuracy := float64(truePosNeg) / float64(len(observed))

	// Output the Accuracy value to standard out.
	fmt.Printf("\nAccuracy = %0.2f\n\n", accuracy)
}

// predict makes a prediction based on our
// trained logistic regression model.
func predict(score float64) float64 {
	// Calculate the predicted probability.
	p := 1 / (1 + math.Exp(-13.65*score+4.89))

	// Output the corresponding class.
	if p >= 0.5 {
		return 1.0
	}

	return 0.0
}
