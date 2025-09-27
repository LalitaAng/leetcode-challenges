package problem383

import "testing"

func TestCanConstruct(t *testing.T) {
    tests := []struct {
        ransomNote string
        magazine   string
        want       bool
    }{
        {"a", "b", false},          
        {"aa", "ab", false},     
        {"aa", "aab", true},       
        {"abc", "cba", true},       
        {"aabbcc", "abcabc", true}, 
        {"abcd", "abc", false},     
        {"", "anything", true},    
        {"longnote", "short", false},
    }

    for _, tt := range tests {
        got := CanConstruct(tt.ransomNote, tt.magazine)
        if got != tt.want {
            t.Errorf("CanConstruct(%q, %q) = %v; want %v",
                tt.ransomNote, tt.magazine, got, tt.want)
        }
    }
}
