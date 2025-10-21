package problem105

/* Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and 
inorder is the inorder traversal of the same tree, construct and return the binary tree. */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
    inorderMap := make(map[int]int)
    for i, val := range inorder {
        inorderMap[val] = i
    }

    var i int
    
    var helper func(j, k int) *TreeNode
    helper = func(j, k int) *TreeNode {
        if j > k {
            return nil
        }
        
        nodeVal := preorder[i]
        i++
        node := &TreeNode{Val: nodeVal}
        idx := inorderMap[nodeVal]
        node.Left = helper(j, idx-1)
        node.Right = helper(idx+1, k)
        return node
    }
    
    return helper(0, len(inorder)-1)
}
