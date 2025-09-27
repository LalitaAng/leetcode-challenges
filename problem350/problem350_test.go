package problem350

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		nums1   []int
		nums2   []int
		want    []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
		{[]int{1, 2}, []int{3, 4}, []int{}},
		{[]int{1, 1, 1}, []int{1, 1}, []int{1, 1}},
		{[]int{}, []int{1, 2, 3}, []int{}},
		{[]int{1, 2, 3}, []int{}, []int{}},
	}

	for _, tt := range tests {
		got := Intersect(tt.nums1, tt.nums2)

		if !reflect.DeepEqual(sortInts(got), sortInts(tt.want)) {
			t.Errorf("Intersect(%v, %v) = %v; want %v",
				tt.nums1, tt.nums2, got, tt.want)
		}
	}
}

func sortInts(nums []int) []int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	for i := 0; i < len(sorted); i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[j] < sorted[i] {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}
	return sorted
}
