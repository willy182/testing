package somethinc

import (
	"reflect"
	"testing"
)

func TestFractions(t *testing.T) {
	type args struct {
		n    int
		data []int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{n: 284000, data: []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}},
			want: []byte(`[{"fraction":100000,"qty":2},{"fraction":50000,"qty":1},{"fraction":20000,"qty":1},{"fraction":10000,"qty":1},{"fraction":2000,"qty":2}]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fractions(tt.args.n, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fractions() = %v, want %v", got, tt.want)
			}
		})
	}
}
