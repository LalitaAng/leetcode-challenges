package problem1611

import "testing"

func TestMinimumOneBitOperations(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero", 0, 0},
		{"one", 1, 1},
		{"two", 2, 3},
		{"three", 3, 2},
		{"four", 4, 7},
		{"five", 5, 6},
		{"six", 6, 4},
		{"seven", 7, 5},
		{"eight", 8, 15},
		{"nine", 9, 14},
		{"ten", 10, 12},
		{"eleven", 11, 13},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinimumOneBitOperations(tt.n)
			if got != tt.want {
				t.Errorf("MinimumOneBitOperations(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
