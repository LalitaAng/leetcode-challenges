package problem56

/* Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, 
and return an array of the non-overlapping intervals that cover all the intervals in the input. */

import "sort"

func Merge(intervals [][]int) [][]int {
    intervals = sortSliceByStart(intervals)

    finalSlice := [][]int{} 

    start, end := intervals[0][0], intervals[0][1]

    for idx := 1; idx < len(intervals); idx++ {
        currentStart := intervals[idx][0]
        currentEnd := intervals[idx][1]

        if currentStart <= end {
            if currentEnd > end {
                end = currentEnd
            }
        } else {
            finalSlice = append(finalSlice, []int{start, end})
            start = currentStart
            end = currentEnd
        }
    }

    finalSlice = append(finalSlice, []int{start, end})

    return finalSlice
}

func sortSliceByStart(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    return intervals
}
