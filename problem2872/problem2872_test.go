package problem2872

import "testing"

func TestMaxKDivisibleComponents(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		edges  [][]int
		values []int
		k      int
		want   int
	}{
		{
			name:   "single_node_divisible",
			n:      1,
			edges:  [][]int{},
			values: []int{6},
			k:      3,
			want:   1,
		},
		{
			name:   "single_node_not_divisible",
			n:      1,
			edges:  [][]int{},
			values: []int{5},
			k:      3,
			want:   0,
		},
		{
			name: "simple_chain_all_divisible",
			n:    3,
			edges: [][]int{
				{0, 1},
				{1, 2},
			},
			values: []int{3, 3, 3},
			k:      3,
			want:   3,
		},
		{
			name: "simple_chain_partial",
			n:    3,
			edges: [][]int{
				{0, 1},
				{1, 2},
			},
			values: []int{1, 2, 3},
			k:      3,
			want:   2,
		},
		{
			name: "star_tree",
			n:    4,
			edges: [][]int{
				{0, 1},
				{0, 2},
				{0, 3},
			},
			values: []int{3, 1, 2, 3},
			k:      3,
			want:   2,
		},
		{
			name: "mixed_tree",
			n:    5,
			edges: [][]int{
				{0, 1},
				{1, 2},
				{1, 3},
				{3, 4},
			},
			values: []int{2, 1, 3, 2, 1},
			k:      3,
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxKDivisibleComponents(tt.n, tt.edges, tt.values, tt.k)
			if got != tt.want {
				t.Errorf(
					"MaxKDivisibleComponents(%d, %v, %v, %d) = %d; want %d",
					tt.n, tt.edges, tt.values, tt.k, got, tt.want,
				)
			}
		})
	}
}
