package problem230

import "testing"

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		k        int
		expected int
	}{
		{
			name: "Example 1",
			root: &TreeNode{3,
				&TreeNode{1, nil, &TreeNode{2, nil, nil}},
				&TreeNode{4, nil, nil},
			},
			k:        1,
			expected: 1,
		},
		{
			name: "Example 2",
			root: &TreeNode{5,
				&TreeNode{3,
					&TreeNode{2, &TreeNode{1, nil, nil}, nil},
					&TreeNode{4, nil, nil},
				},
				&TreeNode{6, nil, nil},
			},
			k:        3,
			expected: 3,
		},
		{
			name: "Single node tree",
			root: &TreeNode{1, nil, nil},
			k:    1,
			expected: 1,
		},
		{
			name: "Right skewed tree",
			root: &TreeNode{1, nil, &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil, nil}}}},
			k:        4,
			expected: 4,
		},
		{
			name: "Left skewed tree",
			root: &TreeNode{4, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, nil}, nil},
			k:        2,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KthSmallest(tt.root, tt.k)
			if result != tt.expected {
				t.Errorf("KthSmallest() = %v; want %v", result, tt.expected)
			}
		})
	}
}
