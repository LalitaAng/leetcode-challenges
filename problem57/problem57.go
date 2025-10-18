package problem57

/* You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval 
and intervals is sorted in ascending order by starti. You are also given an interval newInterval = [start, end] that represents 
the start and end of another interval. Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals 
still does not have any overlapping intervals (merge overlapping intervals if necessary).
Return intervals after the insertion.
Note that you don't need to modify intervals in-place. You can make a new array and return it. */

func Insert(intervals [][]int, newInterval []int) [][]int {
    res := [][]int{}

    for _, interval := range intervals {
        if interval[1] < newInterval[0] {
            res = append(res, interval)
        } else if interval[0] > newInterval[1] {
            res = append(res, newInterval)
            newInterval = interval 
        } else {
            newInterval[0] = min(newInterval[0], interval[0])
            newInterval[1] = max(newInterval[1], interval[1])
        }
    }

    res = append(res, newInterval)
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
