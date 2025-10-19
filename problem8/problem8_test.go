package problem8

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"+1", 1},
		{"", 0},
		{"   ", 0},
		{"+-12", 0},
		{"00000-42a1234", 0},
		{"2147483648", 2147483647},
		{"-2147483649", -2147483648},
		{"0032", 32},
		{"   +0 123", 0},
	}

	for _, tt := range tests {
		got := MyAtoi(tt.input)
		if got != tt.expected {
			t.Errorf("MyAtoi(%q) = %d; expected %d", tt.input, got, tt.expected)
		}
	}
}
