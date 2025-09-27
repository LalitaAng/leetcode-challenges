package problem100

import "testing"
func TestIsSameTree(t *testing.T) {
	tests := []struct {
		p, q   *TreeNode
		expect bool
	}{
		{nil, nil, true},

		{&TreeNode{Val: 1}, nil, false},

		{&TreeNode{Val: 1}, &TreeNode{Val: 1}, true},

		{
			&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			true,
		},

		{
			&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}},
			false,
		},

		{
			&TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			&TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			false,
		},
	}

	for i, tc := range tests {
		got := IsSameTree(tc.p, tc.q)
		if got != tc.expect {
			t.Errorf("test %d failed: expected %v, got %v", i, tc.expect, got)
		}
	}
}
