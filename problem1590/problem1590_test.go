package problem1590

import "testing"

func TestMinSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		p    int
		want int
	}{
		{
			name: "already_divisible",
			nums: []int{3, 1, 2},
			p:    3,
			want: 0,
		},
		{
			name: "single_element_remove",
			nums: []int{1, 2, 3},
			p:    5,
			want: 1,
		},
		{
			name: "remove_middle_subarray",
			nums: []int{3, 1, 4, 2},
			p:    6,
			want: 1,
		},
		{
			name: "remove_prefix",
			nums: []int{6, 3, 5, 2},
			p:    9,
			want: 2,
		},
		{
			name: "remove_suffix",
			nums: []int{1, 2, 3, 4},
			p:    3,
			want: 1,
		},
		{
			name: "whole_array_needed_not_allowed",
			nums: []int{1, 2},
			p:    4,
			want: -1,
		},
		{
			name: "single_element_array_not_divisible",
			nums: []int{5},
			p:    3,
			want: -1,
		},
		{
			name: "single_element_array_divisible",
			nums: []int{6},
			p:    3,
			want: 0,
		},
		{
			name: "large_values",
			nums: []int{1000000000, 1000000000, 1000000000},
			p:    3,
			want: 0,
		},
		{
			name: "complex_case",
			nums: []int{8, 32, 31, 18, 34, 20, 21, 13, 1, 27, 23, 22, 11, 15, 30, 4, 2},
			p:    148,
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinSubarray(tt.nums, tt.p)
			if got != tt.want {
				t.Fatalf("MinSubarray(%v, %d) = %d; want %d",
					tt.nums, tt.p, got, tt.want)
			}
		})
	}
}
