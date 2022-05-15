package tiketdotcom

import "testing"

func TestPohon(t *testing.T) {
	type args struct {
		N []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "unitest 1",
			args: args{N: []int{1, 3, 1, 2}},
			want: 0,
		},
		{
			name: "unitest 2",
			args: args{N: []int{5, 3, 1, 2}},
			want: -1,
		},
		{
			name: "unitest 3",
			args: args{N: []int{1, 3}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pohon(tt.args.N); got != tt.want {
				t.Errorf("Pohon() = %v, want %v", got, tt.want)
			}
		})
	}
}
