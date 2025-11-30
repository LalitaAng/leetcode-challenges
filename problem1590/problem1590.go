package problem1590

/* Given an array of positive integers nums, remove the smallest subarray (possibly empty) such that 
the sum of the remaining elements is divisible by p. It is not allowed to remove the whole array.
Return the length of the smallest subarray that you need to remove, or -1 if it's impossible.
A subarray is defined as a contiguous block of elements in the array. */

func MinSubarray(nums []int, p int) int {
    arraySum := int64(0)
	for _, num := range nums {
		arraySum += int64(num)
	}

	subarrayRem := int(arraySum % int64(p))
	if subarrayRem == 0 {
		return 0
	}

	maxPrefix := map[int]int{0: -1}
	prefixRem := 0
	result := len(nums)

	for i, num := range nums {
		prefixRem = (prefixRem + num) % p
		complement := prefixRem - subarrayRem
		if complement < 0 {
			complement += p
		}

		if idx, ok := maxPrefix[complement]; ok {
			result = min(result, i-idx)
		}
		maxPrefix[prefixRem] = i
	}

	if result == len(nums) {
		return -1
	}
	return result
}
