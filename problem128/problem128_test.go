package problem128

import "testing"

func TestLongestConsecutive(t *testing.T) {
    tests := []struct {
        nums     []int
        expected int
    }{
        {[]int{100, 4, 200, 1, 3, 2}, 4},         
        {[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9}, 
        {[]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, 7},
        {[]int{1, 2, 0, 1}, 3},                   
        {[]int{1, 2, 3, 4, 5}, 5},                
        {[]int{}, 0},                             
        {[]int{10}, 1},                           
        {[]int{1, 2, 0, 1, 2, 3, 3, 4}, 5},       
    }

    for _, tt := range tests {
        got := LongestConsecutive(tt.nums)
        if got != tt.expected {
            t.Errorf("LongestConsecutive(%v) = %d; want %d", tt.nums, got, tt.expected)
        }
    }
}
