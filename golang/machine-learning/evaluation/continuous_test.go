package evaluation

import (
	"github.com/enesanbar/workspace/golang/machine-learning/data"
)

func ExampleEvaluateContinuousData() {
	// the MAE is 2.55 and the mean of our observed values is 14.0,
	// so our MAE is about 20% of our mean value. Not very good, depending on the context.

	// R-squared also gives us a general idea about the deviations of our predictions
	// R-squared measures the proportion of the variance in the observed values
	// that we capture in the predicted values.
	// Remember that R-squared is a percentage and higher percentages are better.
	// Here, we are capturing about 37% of the variance in the variable
	// that we are trying to predict. Not very good.
	EvaluateContinuousData(data.GetContinuousCSVData())
	// Output: MAE = 2.55
	//MSE = 10.51
	//R^2 = 0.37
}
