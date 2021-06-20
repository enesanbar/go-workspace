package statistics

import "testing"

func TestChiSquare(t *testing.T) {
	type args struct {
		observed []float64
		expected []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Flipping a coin",
			args: args{
				observed: []float64{48, 52},
				expected: []float64{50, 50},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chi := ChiSquare(tt.args.observed, tt.args.expected)
			t.Logf("Chi Square: %.2f", chi)
		})
	}
}

func TestPValue(t *testing.T) {
	PValue()
}
