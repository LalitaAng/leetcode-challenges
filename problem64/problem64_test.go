package problem64

import "testing"

func TestMinPathSum(t *testing.T) {
    tests := []struct {
        name string
        grid [][]int
        want int
    }{
        {
            name: "example_1",
            grid: [][]int{
                {1, 3, 1},
                {1, 5, 1},
                {4, 2, 1},
            },
            want: 7,
        },
        {
            name: "example_2",
            grid: [][]int{
                {1, 2, 3},
                {4, 5, 6},
            },
            want: 12,
        },
        {
            name: "single_cell",
            grid: [][]int{
                {5},
            },
            want: 5,
        },
        {
            name: "single_row",
            grid: [][]int{
                {1, 2, 3, 4},
            },
            want: 10,
        },
        {
            name: "single_column",
            grid: [][]int{
                {1},
                {2},
                {3},
            },
            want: 6,
        },
        {
            name: "larger_grid",
            grid: [][]int{
                {7, 1, 3, 5},
                {2, 0, 2, 1},
                {3, 1, 1, 0},
                {4, 2, 1, 2},
            },
            want: 12, 
        },
        {
            name: "with_zero_cost_paths",
            grid: [][]int{
                {0, 0},
                {0, 0},
            },
            want: 0,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := MinPathSum(tt.grid)
            if got != tt.want {
                t.Errorf("MinPathSum() = %d, want %d", got, tt.want)
            }
        })
    }
}
