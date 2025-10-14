package problem230

// Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func KthSmallest(root *TreeNode, k int) int {
    count := 0
    result := 0

    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil || count >= k {
            return
        }

        inorder(node.Left)

        count++
        if count == k {
            result = node.Val
            return
        }

        inorder(node.Right)
    }

    inorder(root)
    return result
}
