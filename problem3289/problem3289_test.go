package problem3289

import (
	"reflect"
	"testing"
)

func TestGetSneakyNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "basic case",
			nums: []int{0, 1, 2, 1, 3, 2},
			want: []int{1, 2},
		},
		{
			name: "duplicates at start",
			nums: []int{2, 2, 1, 0, 1, 3},
			want: []int{2, 1},
		},
		{
			name: "duplicates at end",
			nums: []int{0, 1, 2, 3, 2, 1},
			want: []int{2, 1},
		},
		{
			name: "includes n as value",
			nums: []int{0, 5, 1, 4, 3, 2, 5, 3},
			want: []int{5, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetSneakyNumbers(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSneakyNumbers(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
