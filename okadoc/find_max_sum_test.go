package okadoc

import (
	"testing"
)

func BenchmarkFindMaxSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindMaxSum([]int{9, 2, 8, 1, 5})
	}
}

func BenchmarkFindMaxSumSortRevers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindMaxSumSortRevers([]int{9, 2, 8, 1, 5})
	}
}

func BenchmarkFindMaxSumSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindMaxSumSort([]int{9, 2, 8, 1, 5})
	}
}

func TestFindMaxSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test case find and sum",
			args: args{numbers: []int{5, 9, 7, 11}},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxSum(tt.args.numbers); got != tt.want {
				t.Errorf("FindMaxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMaxSumSortRevers(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test case find and sum",
			args: args{numbers: []int{5, 9, 7, 11}},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxSumSortRevers(tt.args.numbers); got != tt.want {
				t.Errorf("FindMaxSumSortRevers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMaxSumSort(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test case find and sum",
			args: args{numbers: []int{5, 9, 7, 11}},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxSumSort(tt.args.numbers); got != tt.want {
				t.Errorf("FindMaxSumSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
