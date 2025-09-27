package problem383

/* Given two strings ransomNote and magazine, return true if ransomNote can be constructed 
by using the letters from magazine and false otherwise.
Each letter in magazine can only be used once in ransomNote. */

func CanConstruct(ransomNote string, magazine string) bool {
    counts := [26]int{}
    for _, ch := range magazine {
        counts[ch-'a']++
    }

    for _, ch := range ransomNote {
        counts[ch-'a']--
        if counts[ch-'a'] < 0 {
            return false
        }
    }
    return true
}
