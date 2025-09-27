package problem104

/* Given the root of a binary tree, return its maximum depth. A binary tree's maximum depth 
is the number of nodes along the longest path from the root node down to the farthest leaf node. */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    queue := []*TreeNode{root}
    depth := 0
    
    for len(queue) > 0 {
        depth++
        levelSize := len(queue)
        
        for range levelSize {
            node := queue[0]
            queue = queue[1:]
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    
    return depth
}
