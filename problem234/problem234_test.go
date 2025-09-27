package problem234
import "testing"

func buildList(nums []int) *ListNode {
    if len(nums) == 0 {
        return nil
    }
    head := &ListNode{Val: nums[0]}
    curr := head
    for _, v := range nums[1:] {
        curr.Next = &ListNode{Val: v}
        curr = curr.Next
    }
    return head
}

func TestIsPalindrome(t *testing.T) {
    tests := []struct {
        name     string
        input    []int
        expected bool
    }{
        {"empty list", []int{}, true},
        {"single node", []int{1}, true},
        {"two nodes palindrome", []int{1, 1}, true},
        {"two nodes non-palindrome", []int{1, 2}, false},
        {"odd length palindrome", []int{1, 2, 3, 2, 1}, true},
        {"odd length non-palindrome", []int{1, 2, 3, 4, 5}, false},
        {"even length palindrome", []int{1, 2, 2, 1}, true},
        {"even length non-palindrome", []int{1, 2, 3, 4}, false},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            head := buildList(tc.input)
            got := IsPalindrome(head)
            if got != tc.expected {
                t.Errorf("IsPalindrome(%v) = %v, want %v", tc.input, got, tc.expected)
            }
        })
    }
}
