package common

import "testing"

func TestPalindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test anna",
			args: args{"anna"},
			want: true,
		},
		{
			name: "test Mom",
			args: args{"Mom"},
			want: true,
		},
		{
			name: "test malih",
			args: args{"malih"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Palindrome(tt.args.str); got != tt.want {
				t.Errorf("Palindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
