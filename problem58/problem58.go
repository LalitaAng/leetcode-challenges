package problem58

/* Given a string s consisting of words and spaces, return the length of the last word in the string.
A word is a maximal substring consisting of non-space characters only. */

import "strings"

func LengthOfLastWord(s string) int {
    s = strings.TrimSpace(s)
    words := strings.Split(s, " ")
    return len(words[len(words)-1])
}
