package problem11

/* You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of 
the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
Notice that you may not slant the container. */

func MaxArea(height []int) int {
    i, j := 0, len(height)-1
    res := 0

    for i < j {
        h := min(height[i], height[j])
        area := (j - i) * h
        res = max(res, area)

        if height[i] < height[j] {
            i++
        } else {
            j--
        }
    }

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
