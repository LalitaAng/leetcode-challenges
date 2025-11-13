package problem3228

import "testing"

func TestMaxOperations(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "all_zeros",
			s:    "0000",
			want: 0,
		},
		{
			name: "all_ones",
			s:    "1111",
			want: 0,
		},
		{
			name: "single_one_in_middle",
			s:    "010",
			want: 1,
		},
		{
			name: "alternating_ones_and_zeros",
			s:    "1010",
			want: 3,
		},
		{
			name: "complex_case_1",
			s:    "1100101",
			want: 5,
		},
		{
			name: "complex_case_2",
			s:    "1001001",
			want: 3,
		},
		{
			name: "one_zero_between_ones",
			s:    "101",
			want: 1,
		},
		{
			name: "leading_ones_followed_by_zeros",
			s:    "111000",
			want: 3,
		},
		{
			name: "empty_string",
			s:    "",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxOperations(tt.s)
			if got != tt.want {
				t.Errorf("MaxOperations(%q) = %d; want %d", tt.s, got, tt.want)
			}
		})
	}
}
