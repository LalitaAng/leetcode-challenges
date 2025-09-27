package problem3

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"abcabcbb", 3},  
		{"bbbbb", 1},     
		{"pwwkew", 3},    
		{"", 0},         
		{"au", 2},        
		{"dvdf", 3},      
		{"abba", 2},       
		{"tmmzuxt", 5},   
	}

	for _, tt := range tests {
		got := LengthOfLongestSubstring(tt.input)
		if got != tt.want {
			t.Errorf("LengthOfLongestSubstring(%q) = %d; want %d", tt.input, got, tt.want)
		}
	}
}
