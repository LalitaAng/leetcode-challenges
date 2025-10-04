package problem15

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "No triplets",
			nums: []int{1, 2, 3, 4, 5},
			want: [][]int{},
		},
		{
			name: "All zeros",
			nums: []int{0, 0, 0, 0},
			want: [][]int{
				{0, 0, 0},
			},
		},
		{
			name: "Negative and positive balance",
			nums: []int{-2, 0, 1, 1, 2},
			want: [][]int{
				{-2, 0, 2},
				{-2, 1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ThreeSum(tt.nums)

			if got == nil {
				got = [][]int{}
			}
			if tt.want == nil {
				tt.want = [][]int{}
			}

			for _, triplet := range got {
				sort.Ints(triplet)
			}
			for _, triplet := range tt.want {
				sort.Ints(triplet)
			}

			sort.Slice(got, func(i, j int) bool {
				for k := 0; k < len(got[i]); k++ {
					if got[i][k] != got[j][k] {
						return got[i][k] < got[j][k]
					}
				}
				return false
			})
			sort.Slice(tt.want, func(i, j int) bool {
				for k := 0; k < len(tt.want[i]); k++ {
					if tt.want[i][k] != tt.want[j][k] {
						return tt.want[i][k] < tt.want[j][k]
					}
				}
				return false
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
