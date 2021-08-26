package garasidotid

import (
	"testing"
)

func TestIsPrimeByFunc(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
		{
			name: "positif",
			args: args{929},
			want: 1,
		},
		{
			name: "negatif",
			args: args{1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrimeByFunc(tt.args.n); got != tt.want {
				t.Errorf("IsPrimeByFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrimeBySemiNative(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
		{
			name: "positif",
			args: args{929},
			want: 1,
		},
		{
			name: "negatif 1",
			args: args{1},
			want: 2,
		},
		{
			name: "negatif 2",
			args: args{4},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrimeBySemiNative(tt.args.n); got != tt.want {
				t.Errorf("IsPrimeBySemiNative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrimeByNative(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
		{
			name: "positif",
			args: args{929},
			want: 1,
		},
		{
			name: "negatif 1",
			args: args{1},
			want: 2,
		},
		{
			name: "negatif 2",
			args: args{4},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrimeByNative(tt.args.n); got != tt.want {
				t.Errorf("IsPrimeByNative1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsPrimeByFunc(b *testing.B) {
	var num int64 = 928
	for i := 0; i < b.N; i++ {
		IsPrimeByFunc(num)
	}
}

func BenchmarkIsPrimeBySemiNative(b *testing.B) {
	var num int64 = 928
	for i := 0; i < b.N; i++ {
		IsPrimeBySemiNative(num)
	}
}

func BenchmarkIsPrimeByNative(b *testing.B) {
	var num int64 = 928
	for i := 0; i < b.N; i++ {
		IsPrimeByNative(num)
	}
}
