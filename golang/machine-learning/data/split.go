package data

import (
	"io"

	"github.com/kniren/gota/dataframe"
)

func SplitTrainingAndTest(b io.Reader) (dataframe.DataFrame, dataframe.DataFrame) {
	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	advertDF := dataframe.ReadCSV(b)

	// Calculate the number of elements in each set.
	trainingNum := (4 * advertDF.Nrow()) / 5
	testNum := advertDF.Nrow() / 5
	if trainingNum+testNum < advertDF.Nrow() {
		trainingNum++
	}

	// Create the subset indices.
	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	// Enumerate the training indices.
	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}

	// Enumerate the test indices.
	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}
	return advertDF.Subset(trainingIdx), advertDF.Subset(testIdx)
}
