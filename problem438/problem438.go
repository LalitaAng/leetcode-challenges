package problem438

// Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.

func FindAnagrams(s string, p string) []int {
   if len(p) == 0 || len(p) > len(s) {
        return []int{}
    }

    var need, window [26]uint8
    for i := 0; i < len(p); i++ {
        need[p[i]-'a']++
        window[s[i]-'a']++
    }

    var result []int
    if window == need {
        result = append(result, 0)
    }

    for i := len(p); i < len(s); i++ {
        window[s[i-len(p)]-'a']--
        window[s[i]-'a']++
        if window == need {
            result = append(result, i+1-len(p))
        }
    }

	if len(result) == 0 {
        return []int{}
    }

    return result
}
