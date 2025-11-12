package problem2654

/* You are given a 0-indexed array nums consisiting of positive integers. You can do the following operation on the array any number of times:
Select an index i such that 0 <= i < n - 1 and replace either of nums[i] or nums[i+1] with their gcd value.
Return the minimum number of operations to make all elements of nums equal to 1. If it is impossible, return -1.
The gcd of two integers is the greatest common divisor of the two integers. */

func MinOperations(nums []int) int {
	n := len(nums)

	countOnes := 0
	for _, num := range nums {
		if num == 1 {
			countOnes++
		}
	}
	if countOnes > 0 {
		return n - countOnes
	}

	overallGcd := 0
	for _, num := range nums {
		overallGcd = gcd(overallGcd, num)
	}
	if overallGcd > 1 {
		return -1
	}

	const inf = int(1e9)
	minSegmentLength := inf

	for i := 0; i < n; i++ {
		segmentGcd := 0
		for j := i; j < n; j++ {
			segmentGcd = gcd(segmentGcd, nums[j])
			if segmentGcd == 1 {
				length := j - i + 1
				if length < minSegmentLength {
					minSegmentLength = length
				}
				break
			}
		}
	}

	return (minSegmentLength - 1) + (n - 1)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}
