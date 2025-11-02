package problem2257

import "testing"

func TestCountUnguarded(t *testing.T) {
    tests := []struct {
        name   string
        m, n   int
        guards [][]int
        walls  [][]int
        want   int
    }{
        {
            name:   "no guards no walls",
            m:      3,
            n:      3,
            guards: [][]int{},
            walls:  [][]int{},
            want:   9,
        },
        {
            name:   "single guard center",
            m:      3,
            n:      3,
            guards: [][]int{{1, 1}},
            walls:  [][]int{},
            want:   4, 
        },
        {
            name:   "walls block vision",
            m:      4,
            n:      6,
            guards: [][]int{{1, 2}, {3, 1}},
            walls:  [][]int{{2, 3}},
            want:   7,
        },
        {
            name:   "guard in corner",
            m:      2,
            n:      2,
            guards: [][]int{{0, 0}},
            walls:  [][]int{},
            want:   1, 
        },
        {
            name:   "guard blocked by wall immediately",
            m:      3,
            n:      3,
            guards: [][]int{{1, 1}},
            walls:  [][]int{{1, 2}},
            want:   4,
        },
        {
            name:   "multiple guards block each other",
            m:      3,
            n:      5,
            guards: [][]int{{1, 1}, {1, 3}},
            walls:  [][]int{},
            want:   6,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := CountUnguarded(tt.m, tt.n, tt.guards, tt.walls)
            if got != tt.want {
                t.Errorf("CountUnguarded() = %d, want %d", got, tt.want)
            }
        })
    }
}
