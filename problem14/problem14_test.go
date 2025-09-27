package problem14

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
    tests := []struct {
        strs []string
        want string
    }{
        {[]string{"flower", "flow", "flight"}, "fl"},
        {[]string{"dog", "racecar", "car"}, ""},
        {[]string{"interspecies", "interstellar", "interstate"}, "inters"},
        {[]string{"throne", "throne"}, "throne"},
        {[]string{"", ""}, ""},
        {[]string{"a"}, "a"},
        {[]string{"ab", "a"}, "a"},
        {[]string{"abc", "abc", "abc"}, "abc"},
    }

    for _, tt := range tests {
        got := LongestCommonPrefix(tt.strs)
        if got != tt.want {
            t.Errorf("LongestCommonPrefix(%v) = %q; want %q", tt.strs, got, tt.want)
        }
    }
}
