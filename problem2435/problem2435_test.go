package problem2435

import "testing"

func TestNumberOfPaths(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		k    int
		want int
	}{
		{
			name: "single_cell_divisible",
			grid: [][]int{{5}},
			k:    5,
			want: 1,
		},
		{
			name: "single_cell_not_divisible",
			grid: [][]int{{5}},
			k:    3,
			want: 0,
		},
		{
			name: "2x2_simple",
			grid: [][]int{
				{1, 2},
				{3, 4},
			},
			k:    1,
			want: 2, 
		},
		{
			name: "2x2_no_valid_path",
			grid: [][]int{
				{1, 2},
				{3, 4},
			},
			k:    3,
			want: 0,
		},
		{
			name: "leetcode_example_1",
			grid: [][]int{
				{5, 2, 4},
				{3, 0, 5},
				{0, 7, 2},
			},
			k:    3,
			want: 2,
		},
		{
			name: "all_zeros",
			grid: [][]int{
				{0, 0},
				{0, 0},
			},
			k:    5,
			want: 2,
		},
		{
			name: "single_row",
			grid: [][]int{
				{1, 2, 3, 4},
			},
			k:    5,
			want: 1,
		},
		{
			name: "single_column",
			grid: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
			k:    5,
			want: 1,
		},
		{
			name: "larger_grid_no_match",
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			k:    10,
			want: 0,
		},
		{
			name: "larger_grid_multiple_paths",
			grid: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			k:    3,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumberOfPaths(tt.grid, tt.k)
			if got != tt.want {
				t.Errorf("NumberOfPaths(%v, %d) = %d; want %d",
					tt.grid, tt.k, got, tt.want)
			}
		})
	}
}
