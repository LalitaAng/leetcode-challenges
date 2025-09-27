package problem49

import (
	"reflect"
	"sort"
	"testing"
)

func sortGroups(groups [][]string) {
	for i := range groups {
		sort.Strings(groups[i])
	}
	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i]) == 0 || len(groups[j]) == 0 {
			return len(groups[i]) < len(groups[j])
		}
		return groups[i][0] < groups[j][0]
	})
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		input    []string
		expected [][]string
	}{
		{
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}},
		},
		{
			input:    []string{""},
			expected: [][]string{{""}},
		},
		{
			input:    []string{"a"},
			expected: [][]string{{"a"}},
		},
		{
			input:    []string{"ab", "ba", "abc", "cab", "bca"},
			expected: [][]string{{"ab", "ba"}, {"abc", "bca", "cab"}},
		},
	}

	for _, tt := range tests {
		result := GroupAnagrams(tt.input)

		sortGroups(result)
		sortGroups(tt.expected)

		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("groupAnagrams(%v) = %v; expected %v", tt.input, result, tt.expected)
		}
	}
}
