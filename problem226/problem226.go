package problem226

// Given the root of a binary tree, invert the tree, and return its root.

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
    if root != nil {
         root.Left,root.Right = InvertTree(root.Right), InvertTree(root.Left)
    }
    return root
}
