package problem76

/* Given two strings s and t of lengths m and n respectively, 
return the minimum window substring of s such that every character in t (including duplicates) is included in the window. 
If there is no such substring, return the empty string "".
The testcases will be generated such that the answer is unique. */

func MinWindow(s string, t string) string {
    if len(t) == 0 || len(s) == 0 {
        return ""
    }

    tFreq := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        tFreq[t[i]]++
    }

    windowFreq := make(map[byte]int)
    have, need := 0, len(tFreq)
    res := []int{-1, -1}
    resLen := 1 << 31 
    l := 0

    for r := 0; r < len(s); r++ {
        c := s[r]
        windowFreq[c]++

        if count, ok := tFreq[c]; ok && windowFreq[c] == count {
            have++
        }

        for have == need {
            if (r - l + 1) < resLen {
                res = []int{l, r}
                resLen = r - l + 1
            }

            windowFreq[s[l]]--
            if count, ok := tFreq[s[l]]; ok && windowFreq[s[l]] < count {
                have--
            }
            l++
        }
    }

    if resLen == 1<<31 {
        return ""
    }
    return s[res[0] : res[1]+1]
}

// improved version
/*
func MinWindow(s string, t string) string {
    if len(t) == 0 || len(s) == 0 {
        return ""
    }

    var tFreq, windowFreq [128]int
    for i := 0; i < len(t); i++ {
        tFreq[t[i]]++
    }

    have, need := 0, 0
    for _, v := range tFreq {
        if v > 0 {
            need++
        }
    }

    res := []int{-1, -1}
    resLen := 1 << 31
    l := 0

    for r := 0; r < len(s); r++ {
        c := s[r]
        windowFreq[c]++

        if tFreq[c] > 0 && windowFreq[c] == tFreq[c] {
            have++
        }

        for have == need {
            if (r - l + 1) < resLen {
                res = []int{l, r}
                resLen = r - l + 1
            }

            leftChar := s[l]
            windowFreq[leftChar]--
            if tFreq[leftChar] > 0 && windowFreq[leftChar] < tFreq[leftChar] {
                have--
            }
            l++
        }
    }

    if resLen == 1<<31 {
        return ""
    }
    return s[res[0] : res[1]+1]
}
*/
