package problem3381

import "testing"

func TestMaxSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int64
	}{
		{
			name: "simple_positive",
			nums: []int{1, 2, 3, 4, 5},
			k:    2,
			want: 14,
		},
		{
			name: "exact_k_length",
			nums: []int{1, 2, 3},
			k:    3,
			want: 6,
		},
		{
			name: "mixed_positive_negative",
			nums: []int{2, -1, 2, -1, 2},
			k:    2,
			want: 2,
		},
		{
			name: "all_negative",
			nums: []int{-5, -2, -3, -4},
			k:    2,
			want: -5, 
		},
		{
			name: "k_equals_1",
			nums: []int{-1, 5, -2, 3},
			k:    1,
			want: 6, 
		},
		{
			name: "large_values",
			nums: []int{100000, -100000, 100000},
			k:    2,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxSubarraySum(tt.nums, tt.k)
			if got != tt.want {
				t.Fatalf("MaxSubarraySum(%v, %d) = %d; want %d",tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
