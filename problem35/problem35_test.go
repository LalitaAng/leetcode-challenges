package problem35

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},  
		{[]int{1, 3, 5, 6}, 2, 1},  
		{[]int{1, 3, 5, 6}, 7, 4},  
		{[]int{1, 3, 5, 6}, 0, 0},  
		{[]int{1}, 0, 0},           
		{[]int{1}, 2, 1},           
	}

	for _, tt := range tests {
		got := SearchInsert(tt.nums, tt.target)
		if got != tt.want {
			t.Errorf("searchInsert(%v, %d) = %d; want %d",
				tt.nums, tt.target, got, tt.want)
		}
	}
}
