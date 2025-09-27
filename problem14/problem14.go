package problem14

/* Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "". */

func LongestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    prefix := strs[0]
    for i := 1; i < len(strs); i++ {
        j := 0
        for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
            j++
        }
        prefix = prefix[:j] 
        if prefix == "" {
            return ""
        }
    }
    return prefix
}
