package names_scores

import "testing"

func TestCalculateNameScore(t *testing.T) {
	type args struct {
		names []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "One name",
			args: args{
				names: []string{
					"COLIN",
				},
			},
			want: 53,
		},
		{
			name: "Two names",
			args: args{
				names: []string{
					"aaaazzzz",
					"enes",
				},
			},
			want: 194,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateNameScore(tt.args.names); got != tt.want {
				t.Errorf("CalculateNameScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
