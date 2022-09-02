package common

import "testing"

func TestFibonaci1(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test limit 13",
			args: args{13},
			want: "0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144",
		},
		{
			name: "test limit 1",
			args: args{1},
			want: "0",
		},
		{
			name: "test limit 0",
			args: args{0},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonaci1(tt.args.limit); got != tt.want {
				t.Errorf("Fibonaci1() = %v, want %v", got, tt.want)
			}
		})
	}
}
