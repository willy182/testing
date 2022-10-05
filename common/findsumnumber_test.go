package common

import "testing"

func TestFindSumNumber(t *testing.T) {
	type args struct {
		findSum int
		data    []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test: found",
			args: args{findSum: 8, data: []int{5, 2, 3, 1, 10}},
			want: true,
		},
		{
			name: "test: not found",
			args: args{findSum: 9, data: []int{5, 2, 3, 1, 10}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSumNumber(tt.args.findSum, tt.args.data); got != tt.want {
				t.Errorf("FindSumNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
