package problem438

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},

		{"", "a", []int{}},            
		{"a", "", []int{}},            
		{"a", "ab", []int{}},         
		{"aaaaaaaaaa", "aa", []int{0, 1, 2, 3, 4, 5, 6, 7, 8}}, 
		{"abc", "def", []int{}},       
		{"baa", "aa", []int{1}},      
	}

	for _, tt := range tests {
		got := FindAnagrams(tt.s, tt.p)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FindAnagrams(%q, %q) = %v; want %v", tt.s, tt.p, got, tt.want)
		}
	}
}
