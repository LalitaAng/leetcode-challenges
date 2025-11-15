package problem3234

import "testing"

func TestNumberOfSubstrings(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "example 1 - simple case",
			s:        "00011",
			expected: 5,
		},
		{
			name:     "example 2 - more ones",
			s:        "101101",
			expected: 16,
		},
		{
			name:     "single character - zero",
			s:        "0",
			expected: 0,
		},
		{
			name:     "single character - one",
			s:        "1",
			expected: 1,
		},
		{
			name:     "all zeros",
			s:        "0000",
			expected: 0,
		},
		{
			name:     "all ones",
			s:        "1111",
			expected: 10, 
		},
		{
			name:     "two characters - 01",
			s:        "01",
			expected: 2, 
		},
		{
			name:     "two characters - 10",
			s:        "10",
			expected: 2, 
		},
		{
			name:     "two characters - 00",
			s:        "00",
			expected: 0,
		},
		{
			name:     "two characters - 11",
			s:        "11",
			expected: 3, 
		},
		{
			name:     "alternating pattern",
			s:        "010101",
			expected: 10,
		},
		{
			name:     "mostly ones with few zeros",
			s:        "11101111",
			expected: 35,
		},
		{
			name:     "mostly zeros with few ones",
			s:        "00010001",
			expected: 5,
		},
		{
			name:     "zeros at start",
			s:        "00111",
			expected: 9,
		},
		{
			name:     "zeros at end",
			s:        "11100",
			expected: 9,
		},
		{
			name:     "zeros in middle",
			s:        "110011",
			expected: 11,
		},
		{
			name:     "long string with balanced distribution",
			s:        "1100110011",
			expected: 19,
		},
		{
			name:     "single zero in middle of ones",
			s:        "1110111",
			expected: 27,
		},
		{
			name:     "multiple isolated zeros",
			s:        "10101010",
			expected: 14,
		},
		{
			name:     "medium length - pattern",
			s:        "111000111000111",
			expected: 30,
		},
		{
			name:     "three ones",
			s:        "111",
			expected: 6, 
		},
		{
			name:     "one zero two ones",
			s:        "011",
			expected: 5, 
		},
		{
			name:     "two zeros one one",
			s:        "001",
			expected: 2, 
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NumberOfSubstrings(tt.s)
			if result != tt.expected {
				t.Errorf("NumberOfSubstrings(%q) = %d, expected %d", tt.s, result, tt.expected)
			}
		})
	}
}
