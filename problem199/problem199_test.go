package problem199

import (
	"reflect"
	"testing"
)

func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		if v != -1 {
			nodes[i] = &TreeNode{Val: v}
		}
	}
	for i := 0; i < len(vals); i++ {
		if nodes[i] != nil {
			leftIdx, rightIdx := 2*i+1, 2*i+2
			if leftIdx < len(vals) {
				nodes[i].Left = nodes[leftIdx]
			}
			if rightIdx < len(vals) {
				nodes[i].Right = nodes[rightIdx]
			}
		}
	}
	return nodes[0]
}

func TestRightSideView(t *testing.T) {
	tests := []struct {
		name string
		input []int
		want []int
		build func() *TreeNode
	}{
		{
			name:  "empty tree",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "single node",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "simple tree",
			input: []int{1,2,3,-1,5,-1,4},
			want: []int{1,3,4},
		},
		{
    		name: "left skewed tree",
    		build: func() *TreeNode {
        		return &TreeNode{Val: 1,
            		Left: &TreeNode{Val: 2,
                		Left: &TreeNode{Val: 3,
                    		Left: &TreeNode{Val: 4},
                			},
            		},
        		}
    		},
    		want: []int{1,2,3,4},
		},
		{
    		name: "right skewed tree",
    		build: func() *TreeNode {
        		return &TreeNode{Val: 1,
            		Right: &TreeNode{Val: 2,
                		Right: &TreeNode{Val: 3,
                    		Right: &TreeNode{Val: 4},
                		},
            		},
        		}
    		},
    		want: []int{1,2,3,4},
		},
		{
    		name: "right skewed tree",
    		build: func() *TreeNode {
        		return &TreeNode{Val: 1,
            		Right: &TreeNode{Val: 2,
                		Right: &TreeNode{Val: 3,
                    		Right: &TreeNode{Val: 4},
                		},
            		},
        		}
    		},
    		want: []int{1,2,3,4},
		},	
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var root *TreeNode
		if tt.build != nil {
    		root = tt.build()
		} else {
    		root = buildTree(tt.input)
		}
			got := RightSideView(root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("RightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
