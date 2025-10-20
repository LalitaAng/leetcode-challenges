package problem151

/* Given an input string s, reverse the order of the words.
A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
Return a string of the words in reverse order concatenated by a single space.
Note that s may contain leading or trailing spaces or multiple spaces between two words. 
The returned string should only have a single space separating the words. Do not include any extra spaces. */

import "strings"

func ReverseWords(s string) string {
    words := strings.Fields(s)
    for left, right := 0, len(words)-1; left < right; left, right = left+1, right-1 {
        words[left], words[right] = words[right], words[left]
    }
    return strings.Join(words, " ")
}

