package validation

import (
	"testing"

	"github.com/enesanbar/workspace/golang/machine-learning/data"
)

func TestValidate(t *testing.T) {
	Validate(data.GetDiabetesCSVData())
}
