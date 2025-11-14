package problem2536

import (
	"reflect"
	"testing"
)

func TestRangeAddQueries(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		queries [][]int
		want    [][]int
	}{
		{
			name: "Example 1 - Single query",
			n:    3,
			queries: [][]int{
				{1, 1, 2, 2},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 1},
				{0, 1, 1},
			},
		},
		{
			name: "Example 2 - Multiple overlapping queries",
			n:    3,
			queries: [][]int{
				{0, 0, 1, 1},
				{1, 0, 2, 1},
			},
			want: [][]int{
				{1, 1, 0},
				{2, 2, 0},
				{1, 1, 0},
			},
		},

		{
			name:    "Single cell matrix with single query",
			n:       1,
			queries: [][]int{{0, 0, 0, 0}},
			want:    [][]int{{1}},
		},
		{
			name:    "Single cell matrix with no queries",
			n:       1,
			queries: [][]int{},
			want:    [][]int{{0}},
		},
		{
			name:    "No queries - empty result",
			n:       3,
			queries: [][]int{},
			want: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "Full matrix query",
			n:    2,
			queries: [][]int{
				{0, 0, 1, 1},
			},
			want: [][]int{
				{1, 1},
				{1, 1},
			},
		},
		{
			name: "Single row query",
			n:    3,
			queries: [][]int{
				{1, 0, 1, 2},
			},
			want: [][]int{
				{0, 0, 0},
				{1, 1, 1},
				{0, 0, 0},
			},
		},
		{
			name: "Single column query",
			n:    3,
			queries: [][]int{
				{0, 1, 2, 1},
			},
			want: [][]int{
				{0, 1, 0},
				{0, 1, 0},
				{0, 1, 0},
			},
		},
		
		{
			name: "Query at top-left corner",
			n:    3,
			queries: [][]int{
				{0, 0, 0, 0},
			},
			want: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "Query at bottom-right corner",
			n:    3,
			queries: [][]int{
				{2, 2, 2, 2},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 1},
			},
		},
		{
			name: "Multiple queries - same cell",
			n:    2,
			queries: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			want: [][]int{
				{3, 0},
				{0, 0},
			},
		},

		{
			name: "Nested queries",
			n:    4,
			queries: [][]int{
				{0, 0, 3, 3}, 
				{1, 1, 2, 2}, 
			},
			want: [][]int{
				{1, 1, 1, 1},
				{1, 2, 2, 1},
				{1, 2, 2, 1},
				{1, 1, 1, 1},
			},
		},
		{
			name: "Diagonal pattern",
			n:    3,
			queries: [][]int{
				{0, 0, 0, 0},
				{1, 1, 1, 1},
				{2, 2, 2, 2},
			},
			want: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
		{
			name: "Cross pattern",
			n:    5,
			queries: [][]int{
				{2, 0, 2, 4},
				{0, 2, 4, 2}, 
			},
			want: [][]int{
				{0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0},
				{1, 1, 2, 1, 1},
				{0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0},
			},
		},

		{
			name: "Larger matrix with corner queries",
			n:    5,
			queries: [][]int{
				{0, 0, 1, 1}, 
				{3, 3, 4, 4}, 
			},
			want: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 1, 1},
			},
		},
		{
			name: "Multiple non-overlapping queries",
			n:    4,
			queries: [][]int{
				{0, 0, 0, 1},
				{0, 2, 0, 3},
				{2, 0, 2, 1},
				{2, 2, 2, 3},
			},
			want: [][]int{
				{1, 1, 1, 1},
				{0, 0, 0, 0},
				{1, 1, 1, 1},
				{0, 0, 0, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RangeAddQueries(tt.n, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RangeAddQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
