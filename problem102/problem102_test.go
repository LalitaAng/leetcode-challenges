package problem102

import (
    "reflect"
    "testing"
)

func TestLevelOrder(t *testing.T) {
    tests := []struct {
        name string
        root *TreeNode
        want [][]int
    }{
        {
            name: "Empty tree",
            root: nil,
            want: [][]int{},
        },
        {
            name: "Single node",
            root: &TreeNode{Val: 1},
            want: [][]int{{1}},
        },
        {
            name: "Two levels",
            root: &TreeNode{
                Val: 1,
                Left: &TreeNode{Val: 2},
                Right: &TreeNode{Val: 3},
            },
            want: [][]int{{1}, {2, 3}},
        },
        {
            name: "Three levels, unbalanced",
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
            want: [][]int{{3}, {9, 20}, {15, 7}},
        },
        {
            name: "Left-heavy tree",
            root: &TreeNode{
                Val: 1,
                Left: &TreeNode{
                    Val: 2,
                    Left: &TreeNode{
                        Val: 3,
                        Left: &TreeNode{Val: 4},
                    },
                },
            },
            want: [][]int{{1}, {2}, {3}, {4}},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := LevelOrder(tt.root)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("LevelOrder() = %v; want %v", got, tt.want)
            }
        })
    }
}
