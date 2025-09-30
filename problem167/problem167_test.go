package problem167

import (
    "reflect"
    "testing"
)

func TestTwoSum(t *testing.T) {
    tests := []struct {
        numbers []int
        target  int
        want    []int
    }{
        {[]int{2, 7, 11, 15}, 9, []int{1, 2}},
        {[]int{2, 3, 4}, 6, []int{1, 3}},
        {[]int{-1, 0}, -1, []int{1, 2}},
        {[]int{1, 2, 3, 4, 4, 9, 56, 90}, 8, []int{4, 5}},
    }

    for _, tt := range tests {
        got := TwoSum(tt.numbers, tt.target)
        if !reflect.DeepEqual(got, tt.want) {
            t.Errorf("TwoSum(%v, %d) = %v; want %v", tt.numbers, tt.target, got, tt.want)
        }
    }
}
