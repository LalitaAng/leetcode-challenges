package problem647

/* Given a string s, return the number of palindromic substrings in it.
A string is a palindrome when it reads the same backward as forward.
A substring is a contiguous sequence of characters within the string.*/

func CountSubstrings(s string) int {
    n := len(s)
    count := 0

    for i := range n {
        count += expandAroundCenter(s, i, i)   
        count += expandAroundCenter(s, i, i+1)
    }

    return count
}

func expandAroundCenter(s string, left, right int) int {
    n := len(s)
    count := 0

    for left >= 0 && right < n && s[left] == s[right] {
        count++
        left--
        right++
    }

    return count
}
