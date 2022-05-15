package ula

import "testing"

func TestCamelToSnake(t *testing.T) {
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
			args: args{s: "CamelCase"},
			want: "camel_case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelToSnake(tt.args.s); got != tt.want {
				t.Errorf("CamelToSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}
