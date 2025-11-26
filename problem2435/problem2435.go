package problem2435

/* You are given a 0-indexed m x n integer matrix grid and an integer k. You are currently at 
position (0, 0) and you want to reach position (m - 1, n - 1) moving only down or right.
Return the number of paths where the sum of the elements on the path is divisible by k. 
Since the answer may be very large, return it modulo 109 + 7. */

func NumberOfPaths(grid [][]int, k int) int {
    const MOD = 1_000_000_007
    m, n := len(grid), len(grid[0])
    
    dp := make([][]int, n)
    for c := range n {
        dp[c] = make([]int, k)
    }
    
    dp[0][grid[0][0]%k] = 1
    
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if r == 0 && c == 0 {
                continue 
            }
            
            currentValue := grid[r][c] % k
            newCounts := make([]int, k)
            
            if c > 0 {
                for remainder := 0; remainder < k; remainder++ {
                    newRemainder := (remainder + currentValue) % k
                    newCounts[newRemainder] = (newCounts[newRemainder] + dp[c-1][remainder]) % MOD
                }
            }
            
            if r > 0 {
                for remainder := 0; remainder < k; remainder++ {
                    newRemainder := (remainder + currentValue) % k
                    newCounts[newRemainder] = (newCounts[newRemainder] + dp[c][remainder]) % MOD
                }
            }
            
            dp[c] = newCounts
        }
    }
    
    return dp[n-1][0]
}
