package problem105

import "testing"

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     *TreeNode
	}{
		{
			name:     "Example 1",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		{
			name:     "Single node",
			preorder: []int{1},
			inorder:  []int{1},
			want: &TreeNode{
				Val: 1,
			},
		},
		{
			name:     "Only left children",
			preorder: []int{3, 2, 1},
			inorder:  []int{1, 2, 3},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
				},
			},
		},
		{
			name:     "Only right children",
			preorder: []int{1, 2, 3},
			inorder:  []int{1, 2, 3},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildTree(tt.preorder, tt.inorder)
			if !isSameTree(got, tt.want) {
				t.Errorf("Test %v failed.\nExpected: %#v\nGot: %#v", tt.name, tt.want, got)
			}
		})
	}
}
