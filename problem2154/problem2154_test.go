package problem2154

import "testing"

func TestFindFinalValue(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		original int
		want     int
	}{
		{
			name:     "basic example",
			nums:     []int{5, 3, 6, 1, 12},
			original: 3,
			want:     24,
		},
		{
			name:     "no doubling possible",
			nums:     []int{2, 7, 9},
			original: 4,
			want:     4,
		},
		{
			name:     "original already final",
			nums:     []int{1, 2, 3},
			original: 10,
			want:     10,
		},
		{
			name:     "multiple recursions",
			nums:     []int{4, 8, 16, 32},
			original: 4,
			want:     64, 
		},
		{
			name:     "multiple recursions corrected", 
			nums:     []int{4, 8, 16, 32},
			original: 4,
			want:     64,
		},
		{
			name:     "empty nums",
			nums:     []int{},
			original: 7,
			want:     7,
		},
		{
			name:     "nums containing duplicates",
			nums:     []int{2, 2, 4, 8},
			original: 2,
			want:     16,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FindFinalValue(tc.nums, tc.original)
			if got != tc.want {
				t.Errorf("FindFinalValue(%v, %d) = %d; want %d",tc.nums, tc.original, got, tc.want)
			}
		})
	}
}
