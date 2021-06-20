package logistic_regression

import (
	"testing"

	"github.com/enesanbar/workspace/golang/machine-learning/data"
)

func ExamplePlotLogisticFunction() {
	PlotLogisticFunction()
	// Output:
}

func TestPlotHistogramWithCleanedData(t *testing.T) {
	PlotHistogramWithCleanedData(data.GetCleanLoanCSVData())
}

func ExampleTrainModel() {
	TrainModel(data.GetCleanLoanCSVData())
	// Output:
	// p = 1 / ( 1 + exp(- m1 * FICO.score - m2) )
	//
	// m1 = 13.65
	// m2 = -4.89
}

func ExampleTestTrainedData() {
	TestTrainedData(data.GetCleanLoanCSVData())
	// Output:
	// Accuracy = 0.83
}
