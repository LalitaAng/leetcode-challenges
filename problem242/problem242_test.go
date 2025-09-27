package problem242

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"", "", true},
		{"a", "a", true},
		{"a", "b", false},
		{"ab", "ba", true},
		{"abc", "cba", true},
		{"abcd", "abc", false},
	}

	for _, tt := range tests {
		result := IsAnagram(tt.s, tt.t)
		if result != tt.expected {
			t.Errorf("IsAnagram(%q, %q) = %v; want %v", tt.s, tt.t, result, tt.expected)
		}
	}
}
