package common

import "testing"

func TestMinMax(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name    string
		args    args
		wantMin int
		wantMax int
	}{
		// TODO: Add test cases.
		{
			name:    "test 1",
			args:    args{[]int{3, 2, 1, 4, 5}},
			wantMin: 10,
			wantMax: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax := MinMax(tt.args.data)
			if gotMin != tt.wantMin {
				t.Errorf("MinMax() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if gotMax != tt.wantMax {
				t.Errorf("MinMax() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
