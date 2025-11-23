package problem1262

// Given an integer array nums, return the maximum possible sum of elements of the array such that it is divisible by three.

import "math"

func MaxSumDivThree(nums []int) int {
    maxSum := []int{0, math.MinInt, math.MinInt}
    
    for _, num := range nums {
        temp := make([]int, 3)
        copy(temp, maxSum)
        
        remainder := num % 3
        for r := 0; r < 3; r++ {
            newRemainder := (r + remainder) % 3
            maxSum[newRemainder] = max(maxSum[newRemainder], temp[r] + num)
        }
    }
    
    return maxSum[0]
}
