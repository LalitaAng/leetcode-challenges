package problem474

/* You are given an array of binary strings strs and two integers m and n.
Return the size of the largest subset of strs such that there are at most m 0's and n 1's in the subset.
A set x is a subset of a set y if all elements of x are also elements of y. */

import "strings"

func FindMaxForm(strs []string, m int, n int) int {
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    
    for _, str := range strs {
        zeroes := strings.Count(str, "0")
        ones := len(str) - zeroes
        
        for i := m; i >= zeroes; i-- {
            for j := n; j >= ones; j-- {
                dp[i][j] = max(dp[i][j], dp[i-zeroes][j-ones]+1)
            }
        }
    }
    
    return dp[m][n]
}