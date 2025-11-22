package problem3190

import "testing"

func TestMinimumOperations(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1: mix of divisible and non-divisible by 3",
			nums:     []int{1, 2, 3, 4},
			expected: 3,
		},
		{
			name:     "example 2: all divisible by 3",
			nums:     []int{3, 6, 9},
			expected: 0,
		},
		{
			name:     "example 3: none divisible by 3",
			nums:     []int{1, 2, 4, 5},
			expected: 4,
		},
		{
			name:     "single element divisible by 3",
			nums:     []int{9},
			expected: 0,
		},
		{
			name:     "single element not divisible by 3",
			nums:     []int{7},
			expected: 1,
		},
		{
			name:     "array with zero",
			nums:     []int{0, 1, 2},
			expected: 2,
		},
		{
			name:     "all zeros (divisible by 3)",
			nums:     []int{0, 0, 0},
			expected: 0,
		},
		{
			name:     "minimum length array with operation needed",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "minimum length array with no operation needed",
			nums:     []int{3},
			expected: 0,
		},
		{
			name:     "large numbers divisible by 3",
			nums:     []int{300, 3000, 30000},
			expected: 0,
		},
		{
			name:     "large numbers not divisible by 3",
			nums:     []int{301, 3001, 30001},
			expected: 3,
		},
		{
			name:     "all remainders are 1 when divided by 3",
			nums:     []int{1, 4, 7, 10, 13},
			expected: 5,
		},
		{
			name:     "all remainders are 2 when divided by 3",
			nums:     []int{2, 5, 8, 11, 14},
			expected: 5,
		},
		{
			name:     "mixed remainders (0, 1, 2)",
			nums:     []int{3, 4, 5, 6, 7, 8, 9},
			expected: 4,
		},
		{
			name:     "consecutive numbers starting from 0",
			nums:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: 6,
		},
		{
			name:     "alternating divisible and non-divisible",
			nums:     []int{3, 1, 6, 2, 9, 4},
			expected: 3,
		},
		{
			name:     "multiples of 3 in ascending order",
			nums:     []int{3, 6, 9, 12, 15, 18, 21},
			expected: 0,
		},
		{
			name:     "powers of 3",
			nums:     []int{3, 9, 27, 81},
			expected: 0,
		},
		{
			name:     "numbers just before multiples of 3",
			nums:     []int{2, 5, 8, 11, 14},
			expected: 5,
		},
		{
			name:     "numbers just after multiples of 3",
			nums:     []int{4, 7, 10, 13, 16},
			expected: 5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := MinimumOperations(tc.nums)
			if result != tc.expected {
				t.Errorf("MinimumOperations(%v) = %d; expected %d",
					tc.nums, result, tc.expected)
			}
		})
	}
}
