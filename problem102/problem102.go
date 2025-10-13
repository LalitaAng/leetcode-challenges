package problem102

// Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    var res [][]int
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        levelSize := len(queue)
        level := make([]int, 0, levelSize)

        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]

            level = append(level, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        res = append(res, level)
    }

    return res
}
