package problem217

// Given an integer array nums, return true if any value appears at least twice in the array, 
// and return false if every element is distinct.

func ContainsDuplicate(nums []int) bool {
    seen := make(map[int]struct{})
    for _, num := range nums {
        if _, ok := seen[num]; ok {
            return true
        }
        seen[num] = struct{}{}
    }
    return false
}
