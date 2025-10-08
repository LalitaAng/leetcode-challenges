package problem424

/* You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. 
You can perform this operation at most k times.
Return the length of the longest substring containing the same letter you can get after performing the above operations.*/

func CharacterReplacement(s string, k int) int {
    freqMap := make(map[byte]int) 
    longest := 0                  
    left := 0                     
    maxFreq := 0                  

    for right := 0; right < len(s); right++ {
        freqMap[s[right]]++
        if freqMap[s[right]] > maxFreq {
            maxFreq = freqMap[s[right]]
        }

        for (right-left+1)-maxFreq > k {
            freqMap[s[left]]--
            left++
        }

        if right-left+1 > longest {
            longest = right - left + 1
        }
    }

    return longest
}
