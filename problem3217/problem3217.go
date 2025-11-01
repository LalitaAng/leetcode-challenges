package problem3217

/* You are given an array of integers nums and the head of a linked list. 
Return the head of the modified linked list after removing all nodes from the linked list that have a value that exists in nums. */

type ListNode struct {
    Val int
    Next *ListNode
}

func ModifiedList(nums []int, head *ListNode) *ListNode {
    toDelete := make(map[int]struct{}, len(nums))
    for _, v := range nums {
        toDelete[v] = struct{}{}
    }

    dummy := &ListNode{Next: head}
    curr := dummy

    for curr.Next != nil {
        if _, ok := toDelete[curr.Next.Val]; ok {
            curr.Next = curr.Next.Next
        } else {
            curr = curr.Next
        }
    }
    return dummy.Next
}
