package statistics

import (
	"testing"

	"github.com/enesanbar/workspace/golang/machine-learning/data"
)

func TestVisualizeBoxPlot(t *testing.T) {
	VisualizeBoxPlot(data.GetIrisCsv())
}
