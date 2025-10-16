package problem152

/* Given an integer array nums, find a subarray that has the largest product, and return the product.
The test cases are generated so that the answer will fit in a 32-bit integer. */

func MaxProduct(nums []int) int {
    res, currMax, currMin := nums[0], nums[0], nums[0]

    for i := 1; i < len(nums); i++ {
        val := nums[i]
        if val < 0 {
            currMax, currMin = currMin, currMax
        }

        currMax = max(val, currMax*val)
        currMin = min(val, currMin*val)
        res = max(res, currMax)
    }

    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
