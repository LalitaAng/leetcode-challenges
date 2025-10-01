package problem142

import "testing"

func buildList(values []int, pos int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	nodes := make([]*ListNode, len(values))
	for i, v := range values {
		nodes[i] = &ListNode{Val: v}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}

	if pos >= 0 {
		nodes[len(values)-1].Next = nodes[pos]
	}

	return nodes[0]
}

func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		pos      int 
		expected int 
	}{
		{"no cycle", []int{3, 2, 0, -4}, -1, -1},
		{"cycle at node 2", []int{3, 2, 0, -4}, 1, 2},
		{"cycle at head", []int{1, 2}, 0, 1},
		{"single node no cycle", []int{1}, -1, -1},
		{"single node cycle", []int{1}, 0, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.values, tt.pos)
			result := DetectCycle(head)

			if tt.expected == -1 {
				if result != nil {
					t.Errorf("expected nil, got %v", result.Val)
				}
			} else {
				if result == nil {
					t.Errorf("expected node with value %d, got nil", tt.expected)
				} else if result.Val != tt.expected {
					t.Errorf("expected node with value %d, got %d", tt.expected, result.Val)
				}
			}
		})
	}
}
