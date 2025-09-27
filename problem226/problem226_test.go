package problem226

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{
			name: "empty tree",
			root: nil,
			want: nil,
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			want: &TreeNode{Val: 1},
		},
		{
			name: "two level tree",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			want: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 2},
			},
		},
		{
			name: "three level tree",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 1},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InvertTree(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
