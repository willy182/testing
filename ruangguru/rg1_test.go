package ruangguru

import "testing"

func TestPangkat(t *testing.T) {
	var (
		result1, result2 float64
	)

	result1 = 4
	result2 = 1

	type args struct {
		val     int
		pangkat int
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
	}{
		// TODO: Add test cases.
		{
			name:       "Case 1: Pangkat greater than 0",
			args:       args{val: 2, pangkat: 2},
			wantResult: result1,
		},
		{
			name:       "Case 2: Pangkat smaller than or equivalen 0",
			args:       args{val: 2, pangkat: 0},
			wantResult: result2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Pangkat(tt.args.val, tt.args.pangkat); gotResult != tt.wantResult {
				t.Errorf("Pangkat() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
