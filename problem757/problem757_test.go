package problem757

import "testing"

func TestIntersectionSizeTwo(t *testing.T) {
    tests := []struct {
        name      string
        intervals [][]int
        expected  int
    }{
        {
            name:      "Example 1",
            intervals: [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}},
            expected:  3,
        },
        {
            name:      "Example 2",
            intervals: [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}},
            expected:  5,
        },
        {
            name:      "Single interval",
            intervals: [][]int{{1, 5}},
            expected:  2,
        },
        {
            name:      "Non-overlapping",
            intervals: [][]int{{1, 2}, {4, 5}, {7, 8}},
            expected:  6,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := IntersectionSizeTwo(tt.intervals)
            if got != tt.expected {
                t.Errorf("got %d, expected %d", got, tt.expected)
            }
        })
    }
}
