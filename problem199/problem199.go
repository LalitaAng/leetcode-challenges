package problem199

/* Given the root of a binary tree, imagine yourself standing on the right side of it, 
 return the values of the nodes you can see ordered from top to bottom. */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 
func RightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    ans := []int{}
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        n := len(queue)
        for i := range n {
            node := queue[0]
            queue = queue[1:]
            if i == n-1 {
                ans = append(ans, node.Val)
            }
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    
    return ans
}
