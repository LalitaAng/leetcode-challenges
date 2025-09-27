package problem169

import "testing"

func TestMajorityElement(t *testing.T) {
    tests := []struct {
        input    []int
        expected int
    }{
        {[]int{3, 2, 3}, 3},
        {[]int{2, 2, 1, 1, 1, 2, 2}, 2},
        {[]int{4, 4, 4, 2, 4, 2, 4}, 4},
        {[]int{6}, 6}, 
    }

    for _, tt := range tests {
        result := MajorityElement(tt.input)
        if result != tt.expected {
            t.Errorf("majorityElement(%v) = %d; want %d", tt.input, result, tt.expected)
        }
    }
}
