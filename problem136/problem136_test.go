package problem136

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
		{[]int{7, 3, 5, 3, 7}, 5},
		{[]int{10, 10, 20, 30, 30}, 20},
	}

	for _, tt := range tests {
		result := SingleNumber(tt.nums)
		if result != tt.expected {
			t.Errorf("SingleNumber(%v) = %d; want %d", tt.nums, result, tt.expected)
		}
	}
}
