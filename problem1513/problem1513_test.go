package problem1513

import (
	"strings"
	"testing"
)

func TestNumSub(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Example 1: alternating pattern",
			input:    "0110111",
			expected: 9,
		},
		{
			name:     "Example 2: single one",
			input:    "101",
			expected: 2,
		},
		{
			name:     "Example 3: all ones",
			input:    "111111",
			expected: 21,
		},
		{
			name:     "Example 4: all zeros",
			input:    "000",
			expected: 0,
		},
		{
			name:     "Edge: single character - one",
			input:    "1",
			expected: 1,
		},
		{
			name:     "Edge: single character - zero",
			input:    "0",
			expected: 0,
		},
		{
			name:     "Edge: two ones",
			input:    "11",
			expected: 3, 
		},
		{
			name:     "Edge: one zero one",
			input:    "101",
			expected: 2, 
		},
		{
			name:     "Boundary: starts with zero",
			input:    "01111",
			expected: 10, 
		},
		{
			name:     "Boundary: ends with zero",
			input:    "11110",
			expected: 10,
		},
		{
			name:     "Boundary: zeros in middle",
			input:    "110011",
			expected: 6,
		},
		{
			name:     "Boundary: multiple zero groups",
			input:    "100100100",
			expected: 3,
		},
		{
			name:     "Pattern: consecutive groups equal size",
			input:    "110110110",
			expected: 9,
		},
		{
			name:     "Pattern: consecutive groups different sizes",
			input:    "1101110",
			expected: 9,
		},
		{
			name:     "Pattern: long sequence of ones",
			input:    "1111111111", 
			expected: 55,       
		},
		{
			name:     "Modulo: sequence requiring modulo operation",
			input:    strings.Repeat("1", 100000), 
			expected: 49965,                       
		},
		{
			name:     "Special: alternating 1 and 0",
			input:    "10101010",
			expected: 4, 
		},
		{
			name:     "Special: two separate groups",
			input:    "11100111",
			expected: 12, 
		},
		{
			name:     "Special: increasing groups",
			input:    "101100111",
			expected: 10,
		},
		{
			name:     "Comprehensive: mixed pattern 1",
			input:    "11011011101111",
			expected: 22,
		},
		{
			name:     "Comprehensive: long zeros sequence",
			input:    "1110000001111",
			expected: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NumSub(tt.input)
			if result != tt.expected {
				t.Errorf("NumSub(%q) = %d, expected %d", tt.input, result, tt.expected)
			}
		})
	}
}
