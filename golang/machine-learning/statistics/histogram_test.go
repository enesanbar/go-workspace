package statistics

import (
	"testing"

	"github.com/enesanbar/workspace/golang/machine-learning/data"
)

func TestVisualizeHistogram(t *testing.T) {
	VisualizeHistogram(data.GetIrisCsv())
}
