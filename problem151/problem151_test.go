package problem151

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Basic case",
			input:    "the sky is blue",
			expected: "blue is sky the",
		},
		{
			name:     "Extra spaces between words",
			input:    "  hello   world  ",
			expected: "world hello",
		},
		{
			name:     "Single word",
			input:    "word",
			expected: "word",
		},
		{
			name:     "Only spaces",
			input:    "     ",
			expected: "",
		},
		{
			name:     "Trailing and leading spaces",
			input:    "  a good   example  ",
			expected: "example good a",
		},
		{
			name:     "Multiple short words",
			input:    "Go is fun",
			expected: "fun is Go",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseWords(tt.input)
			if result != tt.expected {
				t.Errorf("ReverseWords(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
