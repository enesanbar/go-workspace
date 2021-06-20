package concealed_square

import "testing"

func TestIsFormCorrect(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Input",
			args: args{input: 12341234121231},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFormCorrect(tt.args.input); got != tt.want {
				t.Errorf("IsFormCorrect() = %v, want %v", got, tt.want)
			}
		})
	}
}
