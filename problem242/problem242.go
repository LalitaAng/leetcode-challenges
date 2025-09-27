package problem242

/* Given two strings s and t, return true if t is an anagram of s, and false otherwise. 
s and t consist of lowercase English letters.
An anagram is a word or phrase formed by rearranging the letters of a different word or phrase, 
using all the original letters exactly once. */

func IsAnagram(s string, t string) bool {
    charArrays := [26]int{}
    for _, char := range s {
        pos := char - 'a'
        charArrays[pos]++
    }
    for _, char := range t {
        pos := char - 'a'
        charArrays[pos]--
    }
    for _, value := range charArrays {
        if value != 0 {
            return false
        }
    }
    return true
}
