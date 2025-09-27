package problem217

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{input: []int{1, 2, 3, 1}, expected: true},
		{input: []int{1, 2, 3, 4}, expected: false},
		{input: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, expected: true},
		{input: []int{}, expected: false},
		{input: []int{0}, expected: false},
		{input: []int{100, 200, 300, 400, 500}, expected: false},
	}

	for _, tt := range tests {
		result := ContainsDuplicate(tt.input)
		if result != tt.expected {
			t.Errorf("containsDuplicate(%v) = %v; expected %v", tt.input, result, tt.expected)
		}
	}
}
