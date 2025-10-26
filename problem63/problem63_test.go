package problem63

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "no obstacles 3x3",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			want: 6,
		},
		{
			name: "obstacle in the middle",
			grid: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			want: 2,
		},
		{
			name: "start blocked",
			grid: [][]int{
				{1, 0},
				{0, 0},
			},
			want: 0,
		},
		{
			name: "end blocked",
			grid: [][]int{
				{0, 0},
				{0, 1},
			},
			want: 0,
		},
		{
			name: "single cell no obstacle",
			grid: [][]int{
				{0},
			},
			want: 1,
		},
		{
			name: "single cell obstacle",
			grid: [][]int{
				{1},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniquePathsWithObstacles(tt.grid)
			if got != tt.want {
				t.Errorf("UniquePathsWithObstacles() = %d; want %d", got, tt.want)
			}
		})
	}
}
