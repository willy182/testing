package alami

import "testing"

func TestFactorial(t *testing.T) {
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
			args: args{n: 5},
			want: 120,
		},
		{
			name: "unitest 2",
			args: args{n: 7},
			want: 5040,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.args.n); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
