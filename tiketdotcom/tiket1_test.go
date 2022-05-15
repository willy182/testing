package tiketdotcom

import "testing"

func TestFractions(t *testing.T) {
	type args struct {
		x []int
		y []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "unitest 1",
			args: args{
				x: []int{1, 2, 3, 1, 2},
				y: []int{2, 4, 6, 5, 10},
			},
			want: 3,
		},
		{
			name: "unitest 2",
			args: args{
				x: []int{1, 2, 3, 1, 2, 9},
				y: []int{2, 4, 6, 5, 10},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fractions(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Fractions() = %v, want %v", got, tt.want)
			}
		})
	}
}
