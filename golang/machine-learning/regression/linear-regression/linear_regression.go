package linear_regression

import (
	"fmt"
	"image/color"
	"io"
	"log"
	"math"
	"strconv"

	"github.com/enesanbar/workspace/golang/machine-learning/data"
	"github.com/kniren/gota/dataframe"
	"github.com/sajari/regression"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// DataSummary prints summary statistics of data
// 1) Profiling the data
func DataSummary(b io.Reader) {
	// Create a dataframe from io.Reader
	advertDF := dataframe.ReadCSV(b)

	// Use the Describe method to calculate summary statistics
	// for all of the columns in one shot.
	advertSummary := advertDF.Describe()

	// Output the summary statistics to stdout.
	fmt.Println(advertSummary)
}

// PlotHistogram plots histogram of data
// 1) Profiling the data
func PlotHistogram(b io.Reader) {
	// Create a dataframe from io.Reader
	advertDF := dataframe.ReadCSV(b)

	// Create a histogram for each of the columns in the dataset.
	for _, colName := range advertDF.Names() {

		// Create a plotter.Values value and fill it with the
		// values from the respective column of the dataframe.
		plotVals := make(plotter.Values, advertDF.Nrow())
		for i, floatVal := range advertDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		// Make a plot and set its title.
		p := plot.New()
		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

		// Create a histogram of our values drawn
		// from the standard normal.
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

func PlotScatter(b io.Reader) {
	// Create a dataframe from the CSV file.
	advertDF := dataframe.ReadCSV(b)

	// Extract the target column.
	yVals := advertDF.Col("Sales").Float()

	// Create a scatter plot for each of the features in the dataset.
	for _, colName := range advertDF.Names() {

		// pts will hold the values for plotting
		pts := make(plotter.XYs, advertDF.Nrow())

		// Fill pts with data.
		for i, floatVal := range advertDF.Col(colName).Float() {
			pts[i].X = floatVal
			pts[i].Y = yVals[i]
		}

		// Create the plot.
		p := plot.New()
		p.X.Label.Text = colName
		p.Y.Label.Text = "y"
		p.Add(plotter.NewGrid())

		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(err)
		}
		s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
		s.GlyphStyle.Radius = vg.Points(3)

		// Save the plot to a PNG file.
		p.Add(s)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, "./images/"+colName+"_scatter.png"); err != nil {
			log.Fatal(err)
		}
	}
}

func TrainModel(b io.Reader) {
	trainingData, testData := data.SplitTrainingAndTest(b)

	// In this case we are going to try and model our Sales (y)
	// by the TV feature plus an intercept.  As such, let's create
	// the struct needed to train a model using github.com/sajari/regression.
	var r regression.Regression
	r.SetObserved("Sales")
	r.SetVar(0, "TV")

	// Loop of records in the CSV, adding the training data to the regression value.
	for i, record := range trainingData.Records() {

		// Skip the header.
		if i == 0 {
			continue
		}

		// Parse the Sales regression measure, or "y".
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the TV value.
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Add these points to the regression value.
		r.Train(regression.DataPoint(yVal, []float64{tvVal}))
	}

	// Train/fit the regression model.
	r.Run()

	// Output the trained model parameters.
	fmt.Printf("\nRegression Formula:\n%v\n\n", r.Formula)

	// Evaluating the trained model
	// Loop over the test data predicting y and evaluating the prediction
	// with the mean absolute error.
	var mAE float64
	for i, record := range testData.Records() {

		// Skip the header.
		if i == 0 {
			continue
		}

		// Parse the observed sales.
		yObserved, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the tv value.
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Predict y with our trained model.
		yPredicted, err := r.Predict([]float64{tvVal})

		// Add the to the mean absolute error.
		mAE += math.Abs(yObserved-yPredicted) / float64(len(testData.Records()))
	}

	// Output the MAE to standard out.
	fmt.Printf("MAE = %0.2f\n\n", mAE)
}

func TrainModelMultipleIndependentVariable(b io.Reader) {
	trainingData, testData := data.SplitTrainingAndTest(b)

	// In this case we are going to try and model our Sales (y)
	// by the TV feature plus an intercept.  As such, let's create
	// the struct needed to train a model using github.com/sajari/regression.
	var r regression.Regression
	r.SetObserved("Sales")
	r.SetVar(0, "TV")
	r.SetVar(1, "Radio")

	// Loop of records in the CSV, adding the training data to the regression value.
	for i, record := range trainingData.Records() {

		// Skip the header.
		if i == 0 {
			continue
		}

		// Parse the Sales regression measure, or "y".
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the TV value.
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the Radio value.
		radioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Add these points to the regression value.
		r.Train(regression.DataPoint(yVal, []float64{tvVal, radioVal}))
	}

	// Train/fit the regression model.
	r.Run()

	// Output the trained model parameters.
	fmt.Printf("Regression Formula:\n%v\n", r.Formula)

	// Evaluating the trained model
	// Loop over the test data predicting y and evaluating the prediction
	// with the mean absolute error.
	var mAE float64
	for i, record := range testData.Records() {

		// Skip the header.
		if i == 0 {
			continue
		}

		// Parse the observed sales.
		yObserved, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the tv value.
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the tv value.
		radioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Predict y with our trained model.
		yPredicted, err := r.Predict([]float64{tvVal, radioVal})

		// Add the to the mean absolute error.
		mAE += math.Abs(yObserved-yPredicted) / float64(len(testData.Records()))
	}

	// Output the MAE to standard out.
	fmt.Printf("MAE = %0.2f\n", mAE)
}

func PlotRegressionLine(b io.Reader) {
	// Create a dataframe from the CSV file.
	advertDF := dataframe.ReadCSV(b)

	// Extract the target column.
	yVals := advertDF.Col("Sales").Float()

	// pts will hold the values for plotting.
	pts := make(plotter.XYs, advertDF.Nrow())

	// ptsPred will hold the predicted values for plotting.
	ptsPred := make(plotter.XYs, advertDF.Nrow())

	// Fill pts with data.
	for i, floatVal := range advertDF.Col("TV").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
		ptsPred[i].X = floatVal
		ptsPred[i].Y = predict(floatVal)
	}

	// Create the plot.
	p := plot.New()
	p.X.Label.Text = "TV"
	p.Y.Label.Text = "Sales"
	p.Add(plotter.NewGrid())

	// Add the scatter plot points for the observations.
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	// Add the line plot points for the predictions.
	l, err := plotter.NewLine(ptsPred)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	// Save the plot to a PNG file.
	p.Add(s, l)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "./images/regression_line.png"); err != nil {
		log.Fatal(err)
	}
}

// predict uses our trained regression model to made a prediction.
func predict(tv float64) float64 {
	return 7.07 + tv*0.05
}
