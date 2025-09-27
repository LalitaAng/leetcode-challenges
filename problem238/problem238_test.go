package problem238

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums   []int
		expect []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{2, 3, 4, 5}, []int{60, 40, 30, 24}},
		{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{[]int{0, 1, 2, 3}, []int{6, 0, 0, 0}},
		{[]int{1, 2, 0, 4}, []int{0, 0, 8, 0}},
		{[]int{5}, []int{1}},
		{[]int{}, []int{}},  
	}

	for _, tt := range tests {
		got := ProductExceptSelf(tt.nums)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("productExceptSelf(%v) = %v; want %v", tt.nums, got, tt.expect)
		}
	}
}
