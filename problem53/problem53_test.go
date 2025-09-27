package problem53

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "all positive numbers",
			nums: []int{1, 2, 3, 4, 5},
			want: 15,
		},
		{
			name: "all negative numbers",
			nums: []int{-1, -2, -3, -4},
			want: -1,
		},
		{
			name: "mixed positive and negative numbers",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6, 
		},
		{
			name: "single element",
			nums: []int{7},
			want: 7,
		},
		{
			name: "large negative followed by positive",
			nums: []int{-10, -2, 3, 5, -1},
			want: 8, 
		},
		{
			name: "prefix negative numbers",
			nums: []int{-2, -1, 3},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxSubArray(tt.nums)
			if got != tt.want {
				t.Errorf("MaxSubArray(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
