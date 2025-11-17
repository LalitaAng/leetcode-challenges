package problem1437

import "testing"

func TestKLengthApart(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        k    int
        want bool
    }{
        {
            name: "example_1",
            nums: []int{1, 0, 0, 0, 1, 0, 0, 1},
            k:    2,
            want: true,
        },
        {
            name: "example_2",
            nums: []int{1, 0, 0, 1, 0, 1},
            k:    2,
            want: false,
        },
        {
            name: "no_ones",
            nums: []int{0, 0, 0, 0},
            k:    3,
            want: true,
        },
        {
            name: "single_one",
            nums: []int{0, 0, 1, 0, 0},
            k:    10,
            want: true,
        },
        {
            name: "ones_exactly_k_apart",
            nums: []int{1, 0, 0, 1},
            k:    2,
            want: true,
        },
        {
            name: "ones_less_than_k_apart",
            nums: []int{1, 0, 1},
            k:    2,
            want: false,
        },
        {
            name: "consecutive_ones_invalid",
            nums: []int{1, 1, 1},
            k:    1,
            want: false,
        },
        {
            name: "long_valid",
            nums: []int{1, 0, 0, 0, 0, 0, 1},
            k:    5,
            want: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := KLengthApart(tt.nums, tt.k)
            if got != tt.want {
                t.Errorf("KLengthApart(%v, %d) = %v; want %v",
                    tt.nums, tt.k, got, tt.want)
            }
        })
    }
}
