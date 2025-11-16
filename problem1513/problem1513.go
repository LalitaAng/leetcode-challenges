package problem1513

// Given a binary string s, return the number of substrings with all characters 1's. Since the answer may be too large, return it modulo 10^9 + 7.

func NumSub(s string) int {
    result := int64(0)
    onesCount := int64(0)  
    
    for _, c := range s {
        if c == '1' {
            onesCount++
            result += onesCount
        } else {
            onesCount = 0
        }
    }
    return int(result % 1_000_000_007)
}
