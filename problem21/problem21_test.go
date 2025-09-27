package problem21

import ( 
	"reflect"
	"testing"
)

func listFromSlice(nums []int) *ListNode {
    dummy := &ListNode{}
    curr := dummy
    for _, v := range nums {
        curr.Next = &ListNode{Val: v}
        curr = curr.Next
    }
    return dummy.Next
}

func sliceFromList(head *ListNode) []int {
    var result []int
    for head != nil {
        result = append(result, head.Val)
        head = head.Next
    }
    if result == nil {
        return []int{}
    }
    return result
}


func TestMergeTwoLists(t *testing.T) {
    tests := []struct {
        list1 []int
        list2 []int
        want  []int
    }{
        {[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
        {[]int{}, []int{}, []int{}},
        {[]int{}, []int{0}, []int{0}},
        {[]int{2}, []int{1}, []int{1, 2}},
        {[]int{5, 6, 7}, []int{1, 2, 3}, []int{1, 2, 3, 5, 6, 7}},
    }

    for _, tt := range tests {
        got := MergeTwoLists(listFromSlice(tt.list1), listFromSlice(tt.list2))
        if !reflect.DeepEqual(sliceFromList(got), tt.want) {
            t.Errorf("MergeTwoLists(%v, %v) = %v; want %v",
                tt.list1, tt.list2, sliceFromList(got), tt.want)
        }
    }
}
