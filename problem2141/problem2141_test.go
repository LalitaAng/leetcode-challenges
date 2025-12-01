package problem2141

import "testing"

func TestMaxRunTime(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		batteries []int
		want      int64
	}{
		{
			name:      "single_computer_single_battery",
			n:         1,
			batteries: []int{10},
			want:      10,
		},
		{
			name:      "single_computer_multiple_batteries",
			n:         1,
			batteries: []int{3, 3, 3},
			want:      9,
		},
		{
			name:      "equal_batteries_equal_computers",
			n:         2,
			batteries: []int{3, 3},
			want:      3,
		},
		{
			name:      "more_batteries_than_computers",
			n:         2,
			batteries: []int{1, 2, 3},
			want:      3,
		},
		{
			name:      "uneven_batteries",
			n:         3,
			batteries: []int{10, 10, 3, 3},
			want:      6,
		},
		{
			name:      "large_values",
			n:         3,
			batteries: []int{1000000000, 1000000000, 1000000000},
			want:      1000000000,
		},
		{
			name:      "cannot_run_at_all",
			n:         3,
			batteries: []int{1, 1},
			want:      0,
		},
		{
			name:      "classic_leetcode_example",
			n:         2,
			batteries: []int{3, 3, 3},
			want:      4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxRunTime(tt.n, tt.batteries)
			if got != tt.want {
				t.Errorf("MaxRunTime(%d, %v) = %d; want %d", tt.n, tt.batteries, got, tt.want)
			}
		})
	}
}
