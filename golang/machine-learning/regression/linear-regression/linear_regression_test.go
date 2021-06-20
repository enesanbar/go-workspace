package linear_regression

import (
	"testing"

	"github.com/enesanbar/workspace/golang/machine-learning/data"
)

func TestDataSummary(t *testing.T) {
	DataSummary(data.GetAdvertisingCSVData())
}

func ExamplePlotHistogram() {
	PlotHistogram(data.GetAdvertisingCSVData())
	// Output:
}

func ExamplePlotScatter() {
	PlotScatter(data.GetAdvertisingCSVData())
	// Output:
}

func ExampleTrainModel() {
	// The mean sales value was 14.02 and the standard deviation was 5.21.
	// Thus, our MAE is less than the standard deviations of our sales values
	// and is about 20% of the mean value, and our model has some predictive power.
	TrainModel(data.GetAdvertisingCSVData())
	// Output:
	//Regression Formula:
	//Predicted = 7.0688 + TV*0.0489
	//
	//MAE = 3.01
}

func TestPlotRegressionLine(t *testing.T) {
	PlotRegressionLine(data.GetAdvertisingCSVData())
}

func ExampleTrainModelMultipleIndependentVariable() {
	// Our new multiple regression model has improved our MAE!
	// Now we are definitely in pretty good shape to predict Sales based on our advertising spends
	TrainModelMultipleIndependentVariable(data.GetAdvertisingCSVData())
	// Output:
	// Regression Formula:
	// Predicted = 2.9318 + TV*0.0473 + Radio*0.1794
	// MAE = 1.26
}
