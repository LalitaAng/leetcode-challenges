package problem3512

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "sum divisible by k",
			nums: []int{3, 6, 9},
			k:    3,
			want: 0,
		},
		{
			name: "simple remainder",
			nums: []int{1, 2, 3},
			k:    4,
			want: 2, 
		},
		{
			name: "single element",
			nums: []int{5},
			k:    3,
			want: 2,
		},
		{
			name: "large numbers",
			nums: []int{1_000_000_000, 1_000_000_000},
			k:    3,
			want: 2,
		},
		{
			name: "k greater than sum",
			nums: []int{1, 2},
			k:    10,
			want: 3,
		},
		{
			name: "empty slice",
			nums: []int{},
			k:    5,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinOperations(tt.nums, tt.k)
			if got != tt.want {
				t.Fatalf(
					"MinOperations(%v, %d) = %d; want %d",
					tt.nums, tt.k, got, tt.want,
				)
			}
		})
	}
}
