package problem76

import "testing"

func TestMinWindow(t *testing.T) {
    tests := []struct {
        s, t   string
        expect string
    }{
        {"ADOBECODEBANC", "ABC", "BANC"},
        {"a", "a", "a"},
        {"a", "aa", ""},
        {"ab", "b", "b"},
        {"aa", "aa", "aa"},
        {"ADOBECODEBANC", "AABC", "ADOBECODEBA"}, 
    }

    for _, tt := range tests {
        got := MinWindow(tt.s, tt.t)
        if got != tt.expect {
            t.Errorf("MinWindow(%q, %q) = %q; want %q", tt.s, tt.t, got, tt.expect)
        }
    }
}
