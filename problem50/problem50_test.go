package problem50

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	tests := []struct {
		x    float64
		n    int
		want float64
	}{
		{2.0, 10, 1024.0},
		{2.1, 3, 9.261},
		{2.0, -2, 0.25},
		{1.0, 1000, 1.0},
		{0.0, 5, 0.0},
		{5.0, 0, 1.0},
		{-2.0, 3, -8.0},
		{-2.0, 4, 16.0},
		{2.0, 31, 2147483648.0},
	}

	for _, tt := range tests {
		got := MyPow(tt.x, tt.n)
		if math.Abs(got-tt.want) > 1e-3 {
			t.Errorf("MyPow(%v, %v) = %v; want %v", tt.x, tt.n, got, tt.want)
		}
	}
}
