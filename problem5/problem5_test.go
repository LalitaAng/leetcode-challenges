package problem5

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected []string 
	}{
		{"babad", []string{"bab", "aba"}}, 
		{"cbbd", []string{"bb"}},
		{"a", []string{"a"}},
		{"ac", []string{"a", "c"}},
		{"racecar", []string{"racecar"}},
		{"abacdfgdcaba", []string{"aba", "aca"}}, 
	}

	for _, tt := range tests {
		result := LongestPalindrome(tt.input)
		found := false
		for _, exp := range tt.expected {
			if result == exp {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("For input %q, expected one of %v, but got %q",
				tt.input, tt.expected, result)
		}
	}
}
