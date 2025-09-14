package problem49

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

import "slices"

func GroupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]int)

	for idx, word := range strs {
        sortedRune := []rune(word)
        slices.Sort(sortedRune)
        sortedWord := string(sortedRune)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], idx)
	}

	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
        r := []string{}
        for _, idx := range group {
            r = append(r, strs[idx])
        }
        result = append(result, r)
	}

	return result
}
