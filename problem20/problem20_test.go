package problem20

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
		{"(", false},
		{")", false},
		{"((()))", true},
		{"((())", false},
	}
	
	for _, tt := range tests {
		got := IsValid(tt.s)
		if got != tt.want {
			t.Errorf("IsValid(%q) = %v; want %v", tt.s, got, tt.want)
		}
	}
}