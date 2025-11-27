package problem3381

/* You are given an array of integers nums and an integer k.
Return the maximum sum of a subarray of nums, such that the size of the subarray is divisible by k. */

func MaxSubarraySum(nums []int, k int) int64 {
    minPrefixSum := make([]int64, k)
	prefixSum := int64(0)
	for i := range k - 1 {
		prefixSum += int64(nums[i])
		minPrefixSum[i] = prefixSum
	}

	result := prefixSum + int64(nums[k-1])
	for i := k - 1; i < len(nums); i++ {
		prefixSum += int64(nums[i])
		j := i % k
		result = max(result, prefixSum-minPrefixSum[j])
		minPrefixSum[j] = min(minPrefixSum[j], prefixSum)
	}
	return result
}
