package problem2654

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "already_all_ones",
			nums: []int{1, 1, 1},
			want: 0,
		},
		{
			name: "single_one_present",
			nums: []int{1, 2, 3},
			want: 2, 
		},
		{
			name: "no_one_but_possible",
			nums: []int{2, 4, 3, 9},
			want: 4, 
		},
		{
			name: "no_one_and_impossible",
			nums: []int{2, 4, 6, 8},
			want: -1,
		},
		{
			name: "small_example",
			nums: []int{3, 6, 9, 10},
			want: 4,
		},
		{
			name: "single_element_one",
			nums: []int{1},
			want: 0,
		},
		{
			name: "single_element_non_one",
			nums: []int{2},
			want: -1,
		},
		{
			name: "mixed_case",
			nums: []int{4, 3, 2, 1},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinOperations(tt.nums)
			if got != tt.want {
				t.Errorf("MinOperations(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
