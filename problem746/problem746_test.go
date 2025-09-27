package problem746

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		name string
		cost []int
		want int
	}{
		{
			name: "example 1",
			cost: []int{10, 15, 20},
			want: 15,
		},
		{
			name: "example 2",
			cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want: 6,
		},
		{
			name: "two steps only",
			cost: []int{5, 10},
			want: 5,
		},
		{
			name: "increasing cost",
			cost: []int{1, 2, 3, 4, 5},
			want: 6, 
		},
		{
			name: "all ones",
			cost: []int{1, 1, 1, 1, 1, 1},
			want: 3,
		},
		{
			name: "large numbers",
			cost: []int{1000, 2000, 3000, 1000, 2000, 1000},
			want: 4000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinCostClimbingStairs(tt.cost)
			if got != tt.want {
				t.Errorf("MinCostClimbingStairs(%v) = %d, want %d", tt.cost, got, tt.want)
			}
		})
	}
}
