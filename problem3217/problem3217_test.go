package problem3217

import (
	"reflect"
	"testing"
)

func listToSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range nums {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func TestModifiedList(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		list     []int
		expected []int
	}{
		{
			name:     "empty list",
			nums:     []int{1, 2, 3},
			list:     []int{},
			expected: []int{},
		},
		{
			name:     "no deletions",
			nums:     []int{},
			list:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "delete some values",
			nums:     []int{2, 4},
			list:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 3, 5},
		},
		{
			name:     "delete head",
			nums:     []int{1},
			list:     []int{1, 2, 3},
			expected: []int{2, 3},
		},
		{
			name:     "delete all",
			nums:     []int{1, 2, 3},
			list:     []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "duplicates in list",
			nums:     []int{2},
			list:     []int{1, 2, 2, 3, 2, 4},
			expected: []int{1, 3, 4},
		},
		{
			name:     "supports negative values",
			nums:     []int{-1, 3},
			list:     []int{-1, 2, 3, 4},
			expected: []int{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.list)
			result := ModifiedList(tt.nums, head)
			got := listToSlice(result)

			if got == nil {
    			got = []int{}
			}
			if tt.expected == nil {
    			tt.expected = []int{}
			}

			if !reflect.DeepEqual(got, tt.expected) {
    			t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}
