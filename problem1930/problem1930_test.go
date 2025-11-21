package problem1930

import "testing"

func TestCountPalindromicSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example_1",
			s:    "aabca",
			want: 3,
		},
		{
			name: "example_2",
			s:    "adc",
			want: 0,
		},
		{
			name: "example_3",
			s:    "bbcbaba",
			want: 4, 
		},
		{
			name: "all_same_char",
			s:    "aaaa",
			want: 1, 
		},
		{
			name: "no_middle_chars",
			s:    "abca",
			want: 2, 
		},
		{
			name: "longer_mix",
			s:    "abcbaabc",
			want: 8,
		},
		{
			name: "single_char",
			s:    "a",
			want: 0,
		},
		{
			name: "two_chars",
			s:    "ab",
			want: 0,
		},
		{
			name: "palindromes_with_multiple_middle_options",
			s:    "abcbad",
			want: 3, 
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
            got := CountPalindromicSubsequence(tc.s)
			if got != tc.want {
				t.Errorf("CountPalindromicSubsequence(%q) = %d; want %d", tc.s, got, tc.want)
			}
		})
	}
}
