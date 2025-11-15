package problem3234

/* You are given a binary string s. Return the number of substrings with dominant ones.
A string has dominant ones if the number of ones in the string is greater than or equal to the square of the number of zeros in the string. */

func NumberOfSubstrings(s string) int {
    n := len(s)
    
    prevZeroPos := make([]int, n+1)
    prevZeroPos[0] = -1
    
    for i := 0; i < n; i++ {
        if i == 0 || s[i-1] == '0' {
            prevZeroPos[i+1] = i
        } else {
            prevZeroPos[i+1] = prevZeroPos[i]
        }
    }
    
    result := 0
    
    for i := 1; i <= n; i++ {
        zerosCount := 0
        if s[i-1] == '0' {
            zerosCount = 1
        }
        
        j := i
        for j > 0 && zerosCount*zerosCount <= n {
            onesCount := (i - prevZeroPos[j]) - zerosCount
            
            if zerosCount*zerosCount <= onesCount {
                validSubstrings := j - prevZeroPos[j]
                maxValidLength := onesCount - zerosCount*zerosCount + 1
                
                if maxValidLength < validSubstrings {
                    validSubstrings = maxValidLength
                }
                result += validSubstrings
            }
            
            j = prevZeroPos[j]
            zerosCount++
        }
    }
    
    return result
}
