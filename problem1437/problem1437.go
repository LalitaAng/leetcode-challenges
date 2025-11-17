package problem1437

// Given an binary array nums and an integer k, return true if all 1's are at least k places away from each other, otherwise return false.

func KLengthApart(nums []int, k int) bool {
    prevOneIdx := -1
    for i, val := range nums {
        if val == 1 {
            if prevOneIdx != -1 && i-prevOneIdx-1 < k {
                return false
            }
            prevOneIdx = i
        }
    }
    return true
}
