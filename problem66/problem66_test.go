package problem66

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{
			name:   "simple increment",
			digits: []int{1, 2, 3},
			want:   []int{1, 2, 4},
		},
		{
			name:   "carry over last digit",
			digits: []int{1, 2, 9},
			want:   []int{1, 3, 0},
		},
		{
			name:   "multiple carry",
			digits: []int{9, 9, 9},
			want:   []int{1, 0, 0, 0},
		},
		{
			name:   "single digit not 9",
			digits: []int{7},
			want:   []int{8},
		},
		{
			name:   "single digit 9",
			digits: []int{9},
			want:   []int{1, 0},
		},
		{
			name:   "leading zeros unaffected",
			digits: []int{0, 0, 1, 2, 3},
			want:   []int{0, 0, 1, 2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PlusOne(append([]int{}, tt.digits...)) 
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlusOne(%v) = %v; want %v", tt.digits, got, tt.want)
			}
		})
	}
}
