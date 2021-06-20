package statistics

import (
	"testing"

	"github.com/enesanbar/workspace/golang/machine-learning/data"
)

func TestBasics(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Calculate basic statistical values of iris dataset",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Basics(data.GetIrisCsv())
			Spread(data.GetIrisCsv())
		})
	}
}
