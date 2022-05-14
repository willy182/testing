package alami

import "testing"

func TestNumericReverse(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "unitest 1",
			args: args{n: 689},
			want: 986,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumericReverse(tt.args.n); got != tt.want {
				t.Errorf("NumericReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
