package problem57

import (
    "reflect"
    "testing"
)

func TestInsert(t *testing.T) {
    tests := []struct {
        name       string
        intervals  [][]int
        newInterval []int
        want       [][]int
    }{
        {
            name:       "Insert in middle with merge",
            intervals:  [][]int{{1, 3}, {6, 9}},
            newInterval: []int{2, 5},
            want:       [][]int{{1, 5}, {6, 9}},
        },
        {
            name:       "Insert at beginning without merge",
            intervals:  [][]int{{3, 5}, {6, 7}, {8, 10}, {12, 16}},
            newInterval: []int{1, 2},
            want:       [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
        },
        {
            name:       "Insert with multiple merges",
            intervals:  [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
            newInterval: []int{4, 9},
            want:       [][]int{{1, 2}, {3, 10}, {12, 16}},
        },
        {
            name:       "Insert at end without merge",
            intervals:  [][]int{{1, 3}, {4, 6}},
            newInterval: []int{8, 10},
            want:       [][]int{{1, 3}, {4, 6}, {8, 10}},
        },
        {
            name:       "Insert overlapping last interval",
            intervals:  [][]int{{1, 3}, {6, 9}},
            newInterval: []int{8, 10},
            want:       [][]int{{1, 3}, {6, 10}},
        },
        {
            name:       "Insert into empty list",
            intervals:  [][]int{},
            newInterval: []int{5, 7},
            want:       [][]int{{5, 7}},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Insert(tt.intervals, tt.newInterval)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Insert() = %v; want %v", got, tt.want)
            }
        })
    }
}
