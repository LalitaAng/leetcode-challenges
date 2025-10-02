package problem5

// Given a string s, return the longest palindromic substring in s.

func LongestPalindrome(s string) string {
    n := len(s)
    maxLen := 1
    start := 0
    
    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
    }

    for i := 0; i < n; i++ {
        dp[i][i] = true
    }

    for i := 0; i < n-1; i++ {
        if s[i] == s[i+1] {
            dp[i][i+1] = true
            start = i
            maxLen = 2
        }
    }

    for length := 3; length <= n; length++ {
        for i := 0; i+length-1 < n; i++ {
            j := i + length - 1
            if s[i] == s[j] && dp[i+1][j-1] {
                dp[i][j] = true
                if length > maxLen {
                    start = i
                    maxLen = length
                }
            }
        }
    }

    return s[start : start+maxLen]
}
