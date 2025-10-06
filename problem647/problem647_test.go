package problem647

import "testing"

func TestCountSubstrings(t *testing.T) {
    tests := []struct {
        input string
        want  int
    }{
        {"a", 1},          
        {"aa", 3},        
        {"aaa", 6},        
        {"abc", 3},       
        {"aba", 4},        
        {"abba", 6},       
        {"abac", 5},       
        {"racecar", 10},   
    }

    for _, tt := range tests {
        got := CountSubstrings(tt.input)
        if got != tt.want {
            t.Errorf("CountSubstrings(%q) = %d; want %d", tt.input, got, tt.want)
        }
    }
}
