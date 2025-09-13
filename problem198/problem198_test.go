package problem198

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "empty houses",
			nums: []int{},
			want: 0,
		},
		{
			name: "one house",
			nums: []int{5},
			want: 5,
		},
		{
			name: "two houses pick max",
			nums: []int{2, 3},
			want: 3,
		},
		{
			name: "three houses skip middle",
			nums: []int{2, 7, 9},
			want: 11,
		},
		{
			name: "four houses alternating",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "longer sequence",
			nums: []int{2, 7, 9, 3, 1},
			want: 12, 
		},
		{
			name: "all same values",
			nums: []int{5, 5, 5, 5},
			want: 10, 
		},
		{
			name: "large increasing values",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 25, 
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Rob(tt.nums)
			if got != tt.want {
				t.Errorf("Rob(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
