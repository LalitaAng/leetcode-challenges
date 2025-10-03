package problem56

import (
    "reflect"
    "testing"
)

func TestMerge(t *testing.T) {
    tests := []struct {
        intervals [][]int
        want      [][]int
    }{
        {
            intervals: [][]int{{1,3},{2,6},{8,10},{15,18}},
            want:      [][]int{{1,6},{8,10},{15,18}},
        },
        {
            intervals: [][]int{{1,4},{4,5}},
            want:      [][]int{{1,5}},
        },
        {
            intervals: [][]int{{1,4},{0,2},{3,5}},
            want:      [][]int{{0,5}},
        },
        {
            intervals: [][]int{{1,10},{2,3},{4,5},{6,7},{8,9}},
            want:      [][]int{{1,10}},
        },
        {
            intervals: [][]int{{1,3}},
            want:      [][]int{{1,3}},
        },
    }

    for _, tt := range tests {
        got := Merge(tt.intervals)
        if !reflect.DeepEqual(got, tt.want) {
            t.Errorf("Merge(%v) = %v, want %v", tt.intervals, got, tt.want)
        }
    }
}
