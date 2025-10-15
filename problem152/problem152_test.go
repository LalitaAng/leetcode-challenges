package problem152

import "testing"

func TestMaxProduct(t *testing.T) {
    tests := []struct {
        nums []int
        want int
    }{
        {[]int{2, 3, -2, 4}, 6},
        {[]int{-2, 0, -1}, 0},
        {[]int{-2, 3, -4}, 24},
        {[]int{0, 2}, 2},
        {[]int{-1, -2, -3, 0}, 6},
    }

    for _, tt := range tests {
        got := MaxProduct(tt.nums)
        if got != tt.want {
            t.Errorf("MaxProduct(%v) = %d; want %d", tt.nums, got, tt.want)
        }
    }
}
