package problem73

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name: "basic case",
			input: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			name: "multiple zeros",
			input: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
		{
			name: "single element zero",
			input: [][]int{
				{0},
			},
			expected: [][]int{
				{0},
			},
		},
		{
			name: "single element nonzero",
			input: [][]int{
				{1},
			},
			expected: [][]int{
				{1},
			},
		},
		{
			name: "entire first row and column zero",
			input: [][]int{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 4, 5},
				{0, 7, 8},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetZeroes(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("FAILED %s\nGot:      %v\nExpected: %v", tt.name, tt.input, tt.expected)
			}
		})
	}
}
