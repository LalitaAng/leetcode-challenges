package problem1262

import "testing"

func TestMaxSumDivThree(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example_1",
			nums: []int{3, 6, 5, 1, 8},
			want: 18,
		},
		{
			name: "example_2",
			nums: []int{4},
			want: 0,
		},
		{
			name: "example_3",
			nums: []int{1, 2, 3, 4, 4},
			want: 12,
		},
		{
			name: "all_divisible_by_3",
			nums: []int{3, 9, 12},
			want: 24,
		},
		{
			name: "none_divisible_by_3",
			nums: []int{1, 1, 1, 1},
			want: 3, 
		},
		{
			name: "mixed_small",
			nums: []int{2, 2, 2},
			want: 6,
		},
		{
			name: "bigger_values",
			nums: []int{8, 1, 7, 6, 2},
			want: 24,
		},
		{
			name: "empty_array",
			nums: []int{},
			want: 0,
		},
		{
			name: "single_not_divisible",
			nums: []int{2},
			want: 0,
		},
		{
			name: "large_mix",
			nums: []int{5, 2, 2, 2, 8, 9, 12, 17, 1, 14},
			want: 72,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxSumDivThree(tt.nums)
			if got != tt.want {
				t.Errorf("MaxSumDivThree(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
