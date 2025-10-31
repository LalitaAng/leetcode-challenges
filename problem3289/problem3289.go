package problem3289

/* In the town of Digitville, there was a list of numbers called nums containing integers from 0 to n - 1. 
Each number was supposed to appear exactly once in the list, however, two mischievous numbers sneaked 
in an additional time, making the list longer than usual.
As the town detective, your task is to find these two sneaky numbers. 
Return an array of size two containing the two numbers (in any order), so peace can return to Digitville. */

func GetSneakyNumbers(nums []int) []int {
    n := len(nums) - 2
    seen := make([]bool, n)
    res := make([]int, 0, 2)
    for _, x := range nums {
        if seen[x] {
            res = append(res, x)
            if len(res) == 2 {
                break
            }
        } else {
            seen[x] = true
        }
    }
    return res
}
