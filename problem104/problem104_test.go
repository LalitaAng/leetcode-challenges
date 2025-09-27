package problem104

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "single node",
			root:     &TreeNode{Val: 1},
			expected: 1,
		},
		{
			name: "balanced tree",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: 3,
		},
		{
			name: "skewed tree (left)",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
			},
			expected: 3,
		},
		{
			name: "skewed tree (right)",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:  2,
					Right: &TreeNode{
						Val:  3,
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
			},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxDepth(tt.root)
			if result != tt.expected {
				t.Errorf("MaxDepth(%v) = %d; expected %d", tt.name, result, tt.expected)
			}
		})
	}
}
