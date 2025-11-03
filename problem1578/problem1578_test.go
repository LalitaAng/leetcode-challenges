package problem1578

import "testing"

func TestMinCost(t *testing.T) {
	tests := []struct {
		name       string
		colors     string
		neededTime []int
		want       int
	}{
		{
			name:       "example1",
			colors:     "abaac",
			neededTime: []int{1, 2, 3, 4, 5},
			want:       3, 
		},
		{
			name:       "example2",
			colors:     "abc",
			neededTime: []int{1, 2, 3},
			want:       0, 
		},
		{
			name:       "example3",
			colors:     "aabaa",
			neededTime: []int{1, 2, 3, 4, 1},
			want:       2, 
		},
		{
			name:       "all same",
			colors:     "aaaa",
			neededTime: []int{4, 2, 8, 1},
			want:       7, 
		},
		{
			name:       "alternating colors",
			colors:     "ababab",
			neededTime: []int{1, 10, 2, 10, 3, 10},
			want:       0,
		},
		{
			name:       "two same adjacent",
			colors:     "aa",
			neededTime: []int{5, 1},
			want:       1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinCost(tt.colors, tt.neededTime)
			if got != tt.want {
				t.Errorf("MinCost(%q, %v) = %d; want %d", tt.colors, tt.neededTime, got, tt.want)
			}
		})
	}
}
