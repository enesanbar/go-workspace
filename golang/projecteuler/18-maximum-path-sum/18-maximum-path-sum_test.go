package maximum_path_sum

import "testing"

func TestFindMaxPathInTriangle(t *testing.T) {
	type args struct {
		input [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Single line array",
			args: args{
				input: [][]int{
					{3},
				},
			},
			want: 3,
		},
		{
			name: "Two line array",
			args: args{
				input: [][]int{
					{3, 0},
					{3, 4},
				},
			},
			want: 7,
		},
		{
			name: "Two line array (transposed)",
			args: args{
				input: [][]int{
					{3, 0},
					{4, 3},
				},
			},
			want: 7,
		},
		{
			name: "Small array",
			args: args{
				input: [][]int{
					{3, 0, 0, 0},
					{7, 4, 0, 0},
					{2, 4, 6, 0},
					{8, 5, 9, 3},
				},
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMaxPathInTriangle(tt.args.input)
			if got != tt.want {
				t.Errorf("FindMaxPathInTriangle() got = %v, want %v", got, tt.want)
			}
		})
	}
}

var input = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindMaxPathInTriangle(input)
	}
}
