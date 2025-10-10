package problem300

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4, // LIS: [2, 3, 7, 101]
		},
		{
			name: "All increasing",
			nums: []int{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "All decreasing",
			nums: []int{5, 4, 3, 2, 1},
			want: 1,
		},
		{
			name: "Mixed sequence",
			nums: []int{0, 1, 0, 3, 2, 3},
			want: 4, // LIS: [0, 1, 2, 3]
		},
		{
			name: "With duplicates",
			nums: []int{4, 10, 4, 3, 8, 9},
			want: 3, // LIS: [4, 8, 9]
		},
		{
			name: "Single element",
			nums: []int{7},
			want: 1,
		},
		{
			name: "Empty array",
			nums: []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LengthOfLIS(tt.nums)
			if got != tt.want {
				t.Errorf("LengthOfLIS(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}
