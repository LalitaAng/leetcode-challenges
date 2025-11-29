package problem3512

/* You are given an integer array nums and an integer k. You can perform the following operation any number of times:
Select an index i and replace nums[i] with nums[i] - 1.
Return the minimum number of operations required to make the sum of the array divisible by k.*/

func MinOperations(nums []int, k int) int {
    var sum int64

    for _, num := range nums {
        sum += int64(num)
    }

    operations := int(sum % int64(k))
    return operations
}
