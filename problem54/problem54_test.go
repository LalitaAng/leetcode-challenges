package problem54

import (
    "reflect"
    "testing"
)

func TestSpiralOrder(t *testing.T) {
    tests := []struct {
        name     string
        matrix   [][]int
        expected []int
    }{
        {
            name: "3x4 matrix",
            matrix: [][]int{
                {1, 2, 3, 4},
                {5, 6, 7, 8},
                {9, 10, 11, 12},
            },
            expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
        },
        {
            name: "1x4 matrix",
            matrix: [][]int{
                {1, 2, 3, 4},
            },
            expected: []int{1, 2, 3, 4},
        },
        {
            name: "4x1 matrix",
            matrix: [][]int{
                {1},
                {2},
                {3},
                {4},
            },
            expected: []int{1, 2, 3, 4},
        },
        {
            name: "2x2 matrix",
            matrix: [][]int{
                {1, 2},
                {3, 4},
            },
            expected: []int{1, 2, 4, 3},
        },
        {
            name: "3x3 matrix",
            matrix: [][]int{
                {1, 2, 3},
                {4, 5, 6},
                {7, 8, 9},
            },
            expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
        },
        {
            name: "empty matrix",
            matrix: [][]int{},
            expected: []int{},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := SpiralOrder(tt.matrix)
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("SpiralOrder(%v) = %v; want %v", tt.matrix, result, tt.expected)
            }
        })
    }
}
