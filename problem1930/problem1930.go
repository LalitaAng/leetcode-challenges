package problem1930

/* Given a string s, return the number of unique palindromes of length three that are a subsequence of s.
Note that even if there are multiple ways to obtain the same subsequence, it is still only counted once.
A palindrome is a string that reads the same forwards and backwards.
A subsequence of a string is a new string generated from the original string with some characters 
(can be none) deleted without changing the relative order of the remaining characters.
For example, "ace" is a subsequence of "abcde". */

func CountPalindromicSubsequence(s string) int {
    firstOccurrence := make([]int, 26)
    lastOccurrence := make([]int, 26)
    
    for i := range firstOccurrence {
        firstOccurrence[i] = -1
    }
    
    for i := 0; i < len(s); i++ {
        charIndex := s[i] - 'a'
        if firstOccurrence[charIndex] == -1 {
            firstOccurrence[charIndex] = i
        }
        lastOccurrence[charIndex] = i
    }
    
    totalUniquePalindromes := 0
    
    for char := 0; char < 26; char++ {
        if firstOccurrence[char] < lastOccurrence[char] {
            seenMiddleChars := [26]bool{}
            uniqueMiddleCount := 0
            
            for pos := firstOccurrence[char] + 1; pos < lastOccurrence[char]; pos++ {
                middleCharIndex := s[pos] - 'a'
                if !seenMiddleChars[middleCharIndex] {
                    seenMiddleChars[middleCharIndex] = true
                    uniqueMiddleCount++
                }
            }
            totalUniquePalindromes += uniqueMiddleCount
        }
    }
    
    return totalUniquePalindromes
}
