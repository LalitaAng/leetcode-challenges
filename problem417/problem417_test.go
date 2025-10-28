package problem417

import (
	"reflect"
	"sort"
	"testing"
)

func sort2D(a [][]int) {
	sort.Slice(a, func(i, j int) bool {
		if a[i][0] == a[j][0] {
			return a[i][1] < a[j][1]
		}
		return a[i][0] < a[j][0]
	})
}

func TestPacificAtlantic(t *testing.T) {
	tests := []struct {
		name     string
		heights  [][]int
		expected [][]int
	}{
		{
			name: "Example 1",
			heights: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			expected: [][]int{
				{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0},
			},
		},
		{
			name: "Single Cell",
			heights: [][]int{
				{10},
			},
			expected: [][]int{{0, 0}},
		},
		{
			name:     "Empty Grid",
			heights:  [][]int{},
			expected: nil,
		},
		{
			name: "Flat Grid",
			heights: [][]int{
				{1, 1},
				{1, 1},
			},
			expected: [][]int{
				{0, 0}, {0, 1}, {1, 0}, {1, 1},
			},
		},
		{
			name: "Increasing Diagonal",
			heights: [][]int{
				{1, 2},
				{2, 3},
			},
			expected: [][]int{
				{0, 1}, {1, 0}, {1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PacificAtlantic(tt.heights)
			sort2D(result)
			sort2D(tt.expected)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PacificAtlantic() = %v, want %v", result, tt.expected)
			}
		})
	}
}
