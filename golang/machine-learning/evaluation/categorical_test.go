package evaluation

import (
	"github.com/enesanbar/workspace/golang/machine-learning/data"
)

func ExampleEvaluateCategoricalData() {
	// Accuracy: 97%! That's pretty good. That means we were right 97% of the time
	csvData := data.GetLabeledCategoricalCSVData()
	EvaluateCategoricalData(csvData)
	// Output: Accuracy = 0.97
}

func ExampleEvaluateCategoricalDataWithClasses() {
	csvData := data.GetLabeledCategoricalCSVData()
	EvaluateCategoricalDataWithClasses(csvData)
	// Output:
	//Precision (class 0) = 1.00
	//Recall (class 0) = 1.00
	//Precision (class 1) = 0.96
	//Recall (class 1) = 0.94
	//Precision (class 2) = 0.94
	//Recall (class 2) = 0.96
}

func ExampleEvaluateAUC() {
	EvaluateAUC()
	// Output:
	//true  positive rate: [0 0.5 0.5 1 1]
	//false positive rate: [0 0 0.5 0.5 1]
	//auc: 0.75
}
