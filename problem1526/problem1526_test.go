package problem1526

import "testing"

func TestMinNumberOperations(t *testing.T) {
    tests := []struct {
        name   string
        target []int
        want   int
    }{
        {
            name:   "example1",
            target: []int{1, 2, 3, 2, 1},
            want:   3,
        },
        {
            name:   "example2",
            target: []int{3, 1, 1, 2},
            want:   4,
        },
        {
            name:   "example3",
            target: []int{3, 1, 5, 4, 2},
            want:   7,
        },
        {
            name:   "single element",
            target: []int{5},
            want:   5,
        },
        {
            name:   "empty",
            target: []int{},
            want:   0,
        },
        {
            name:   "plateau",
            target: []int{2, 2, 2},
            want:   2,
        },
        {
            name:   "increasing only",
            target: []int{1, 2, 4, 8},
            want:   8,
        },
        {
            name:   "decreasing only",
            target: []int{5, 4, 3, 2},
            want:   5,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := MinNumberOperations(tt.target)
            if got != tt.want {
                t.Errorf("MinNumberOperations(%v) = %d; want %d", tt.target, got, tt.want)
            }
        })
    }
}
