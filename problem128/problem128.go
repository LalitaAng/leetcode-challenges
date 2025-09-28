package problem128

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

func LongestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    set := make(map[int]bool)
    for _, num := range nums {
        set[num] = true
    }

    longest := 0
    for num := range set { 
        if !set[num-1] {
            current := num
            streak := 1

            for set[current+1] {
                current++
                streak++
            }

            if streak > longest {
                longest = streak
            }
        }
    }

    return longest
}

