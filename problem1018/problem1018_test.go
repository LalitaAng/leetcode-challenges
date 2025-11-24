package problem1018

import "testing"

func TestPrefixesDivBy5(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []bool
	}{
		{
			name:     "example 1 - basic case",
			nums:     []int{0, 1, 1},
			expected: []bool{true, false, false},
		},
		{
			name:     "example 2 - multiple divisible",
			nums:     []int{1, 1, 1},
			expected: []bool{false, false, false},
		},
		{
			name:     "example 3 - mixed results",
			nums:     []int{0, 1, 1, 1, 1, 1},
			expected: []bool{true, false, false, false, true, false},
		},
		{
			name:     "single zero",
			nums:     []int{0},
			expected: []bool{true},
		},
		{
			name:     "single one",
			nums:     []int{1},
			expected: []bool{false},
		},
		{
			name:     "all zeros",
			nums:     []int{0, 0, 0, 0},
			expected: []bool{true, true, true, true},
		},
		{
			name:     "all ones - binary: 1=1, 11=3, 111=7, 1111=15",
			nums:     []int{1, 1, 1, 1},
			expected: []bool{false, false, false, true}, 
		},
		{
			name:     "alternating pattern - 1=1, 10=2, 101=5, 1010=10, 10101=21, 101010=42",
			nums:     []int{1, 0, 1, 0, 1, 0},
			expected: []bool{false, false, true, true, false, false},
		},
		{
			name:     "divisible by 5 at specific positions - 1=1, 10=2, 101=5, 1010=10, 10101=21",
			nums:     []int{1, 0, 1, 0, 1},
			expected: []bool{false, false, true, true, false},
		},
		{
			name:     "binary 10 (decimal 2)",
			nums:     []int{1, 0},
			expected: []bool{false, false},
		},
		{
			name:     "binary 101 (decimal 5)",
			nums:     []int{1, 0, 1},
			expected: []bool{false, false, true},
		},
		{
			name:     "binary 1010 (decimal 10)",
			nums:     []int{1, 0, 1, 0},
			expected: []bool{false, false, true, true},
		},
		{
			name:     "binary 1111 (decimal 15)",
			nums:     []int{1, 1, 1, 1},
			expected: []bool{false, false, false, true},
		},
		{
			name:     "pattern leading to multiples of 5 - 1=1, 11=3, 110=6, 1100=12, 11001=25",
			nums:     []int{1, 1, 0, 0, 1},
			expected: []bool{false, false, false, false, true},
		},
		{
			name:     "starts with 0 then builds to 5",
			nums:     []int{0, 1, 0, 1},
			expected: []bool{true, false, false, true}, 
		},
		{
			name:     "multiple of 5 early - 0=0, 00=0, 000=0, 0000=0, 00001=1",
			nums:     []int{0, 0, 0, 0, 1},
			expected: []bool{true, true, true, true, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PrefixesDivBy5(tt.nums)

			if len(result) != len(tt.expected) {
				t.Errorf("Length mismatch: got %d, want %d", len(result), len(tt.expected))
				return
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					decimalValue := calculateDecimal(tt.nums[:i+1])
					t.Errorf("Index %d: got %v, want %v (binary prefix: %v, decimal: %d, %d mod 5 = %d)", i, result[i], tt.expected[i], tt.nums[:i+1], decimalValue, decimalValue, decimalValue%5)
				}
			}
		})
	}
}

func calculateDecimal(binary []int) int {
	result := 0
	for _, bit := range binary {
		result = result*2 + bit
	}
	return result
}
