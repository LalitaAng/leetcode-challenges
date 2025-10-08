package problem424

import "testing"

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		s      string
		k      int
		expect int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
		{"AAAA", 2, 4},
		{"ABCDE", 1, 2},
		{"BAAAB", 2, 5},
		{"ABCABC", 2, 4},
		{"A", 0, 1},
		{"", 2, 0},
	}

	for _, tc := range tests {
		got := CharacterReplacement(tc.s, tc.k)
		if got != tc.expect {
			t.Errorf("CharacterReplacement(%q, %d) = %d; want %d",
				tc.s, tc.k, got, tc.expect)
		}
	}
}
