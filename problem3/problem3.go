package problem3

// Given a string s, find the length of the longest substring without duplicate characters.

func LengthOfLongestSubstring(s string) int {
    seen := make(map[byte]int) 
    left := 0
    maxLen := 0

    for right := 0; right < len(s); right++ {
        char := s[right]

        if idx, found := seen[char]; found && idx >= left {
            left = idx + 1
        }

        seen[char] = right

        if right-left+1 > maxLen {
            maxLen = right - left + 1
        }
    }

    return maxLen
}

/* optimized version
func LengthOfLongestSubstring(s string) int {
    n := len(s)
    maxLength := 0
    lastIndex := make([]int, 128)
    
    start := 0
    for end := 0; end < n; end++ {
        currentChar := s[end]
        if lastIndex[currentChar] > start {
            start = lastIndex[currentChar]
        }
        if end-start+1 > maxLength {
            maxLength = end - start + 1
        }
        lastIndex[currentChar] = end + 1
    }
    
    return maxLength
}
*/
