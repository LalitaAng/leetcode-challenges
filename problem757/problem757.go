package problem757

/* You are given a 2D integer array intervals where intervals[i] = [starti, endi] represents 
all the integers from starti to endi inclusively.
A containing set is an array nums where each interval from intervals has at least two integers in nums.
For example, if intervals = [[1,3], [3,7], [8,9]], then [1,2,4,7,8,9] and [2,3,4,8,9] are containing sets.
Return the minimum possible size of a containing set. */

import "slices"

func IntersectionSizeTwo(intervals [][]int) int {
    slices.SortFunc(intervals, func(i1, i2 []int) int {
        if i1[1] != i2[1] {
            return i1[1] - i2[1]
        }
        return i2[0] - i1[0] 
    })
    
    result := 0
    p1, p2 := -1, -1
    
    for _, interval := range intervals {
        start, end := interval[0], interval[1]
        if p2 < start {
            result += 2
            p1 = end - 1
            p2 = end
        } else if p1 < start {
            result += 1
            p1 = p2
            p2 = end
        }
    }
    
    return result
}
