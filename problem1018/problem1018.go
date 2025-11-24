package problem1018

/* You are given a binary array nums (0-indexed).
We define xi as the number whose binary representation is the subarray nums[0..i] 
(from most-significant-bit to least-significant-bit).
For example, if nums = [1,0,1], then x0 = 1, x1 = 2, and x2 = 5.
Return an array of booleans answer where answer[i] is true if xi is divisible by 5. */

func PrefixesDivBy5(nums []int) []bool {
    result := make([]bool, len(nums))
    remainder := 0
    
    for i, bit := range nums {
        remainder = (remainder << 1 | bit) % 5
        result[i] = remainder == 0
    }
    return result
}
