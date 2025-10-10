package problem300

// Given an integer array nums, return the length of the longest strictly increasing subsequence.

func LengthOfLIS(nums []int) int {
	tails := []int{}
	for _, num := range nums {
		if len(tails) == 0 || tails[len(tails)-1] < num {
			tails = append(tails, num)
		} else {
			index := lowerBound(tails, num)
			tails[index] = num
		}
	}
	return len(tails)
}

func lowerBound(tails []int, target int) int {
	left, right := 0, len(tails)-1
	for left <= right {
		mid := (left + right) / 2
		if tails[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
