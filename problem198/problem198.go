package problem198

/* You are a professional robber planning to rob houses along a street. 
Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them 
is that adjacent houses have security systems connected and it will automatically contact the police 
if two adjacent houses were broken into on the same night.
Given an integer array nums representing the amount of money of each house, return the maximum amount of money 
you can rob tonight without alerting the police.*/

func Rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    
    prev2 := nums[0]                   
    prev1 := max(nums[0], nums[1])    
    
    for i := 2; i < len(nums); i++ {
        current := max(nums[i] + prev2, prev1)
        prev2 = prev1
        prev1 = current
    }
    
    return prev1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
