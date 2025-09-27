package problem141

import "testing"

func TestHasCycle(t *testing.T) {
	if HasCycle(nil) {
		t.Errorf("Expected false for empty list, got true")
	}

	node1 := &ListNode{Val: 1}
	if HasCycle(node1) {
		t.Errorf("Expected false for single node without cycle, got true")
	}

	node1.Next = node1
	if !HasCycle(node1) {
		t.Errorf("Expected true for single node with cycle, got false")
	}
	node1.Next = nil 

	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3
	if HasCycle(node1) {
		t.Errorf("Expected false for list without cycle, got true")
	}

	node3.Next = node2
	if !HasCycle(node1) {
		t.Errorf("Expected true for list with cycle, got false")
	}
}
