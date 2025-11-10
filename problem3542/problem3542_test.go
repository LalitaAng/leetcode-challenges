package problem3542

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "empty_array",
			nums: []int{},
			want: 0,
		},
		{
			name: "single_element_nonzero",
			nums: []int{5},
			want: 1,
		},
		{
			name: "all_zero_elements",
			nums: []int{0, 0, 0},
			want: 0,
		},
		{
			name: "increasing_sequence",
			nums: []int{1, 2, 3, 4},
			want: 4,
		},
		{
			name: "decreasing_sequence",
			nums: []int{4, 3, 2, 1},
			want: 4,
		},
		{
			name: "mixed_sequence",
			nums: []int{1, 3, 2, 4},
			want: 4,
		},
		{
			name: "sequence_with_zeros",
			nums: []int{0, 1, 0, 2, 0, 3},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinOperations(tt.nums)
			if got != tt.want {
				t.Errorf("MinOperations(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}
