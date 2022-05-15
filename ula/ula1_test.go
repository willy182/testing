package ula

import "testing"

func TestGetNumberLetter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "unitest 1",
			args: args{s: "occurrences"},
			want: "o1c3u1r2e2n1s1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNumberLetter(tt.args.s); got != tt.want {
				t.Errorf("GetNumberLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
