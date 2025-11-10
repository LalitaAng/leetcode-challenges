package problem3542

/* You are given an array nums of size n, consisting of non-negative integers. Your task is to apply 
some (possibly zero) operations on the array so that all elements become 0.
In one operation, you can select a subarray [i, j] (where 0 <= i <= j < n) and set all occurrences of the 
minimum non-negative integer in that subarray to 0.
Return the minimum number of operations required to make all elements in the array 0. */

func MinOperations(nums []int) int {
	stack := []int{}
	count := 0

	for _, num := range nums {
		for len(stack) > 0 && stack[len(stack)-1] > num {
			stack = stack[:len(stack)-1]
		}
		if num == 0 {
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] < num {
			count++
			stack = append(stack, num)
		}
	}
	return count
}
