package problem121

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "increasing prices",
			prices: []int{1, 2, 3, 4, 5},
			want:   4, 
		},
		{
			name:   "decreasing prices",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			name:   "single valley and peak",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5, 
		},
		{
			name:   "valley in the middle",
			prices: []int{2, 4, 1},
			want:   2, 
		},
		{
			name:   "same prices",
			prices: []int{5, 5, 5, 5},
			want:   0,
		},
		{
			name:   "only one price",
			prices: []int{10},
			want:   0,
		},
		{
			name:   "profit late in array",
			prices: []int{9, 8, 7, 1, 2},
			want:   1, 
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxProfit(tt.prices)
			if got != tt.want {
				t.Errorf("MaxProfit(%v) = %d; want %d", tt.prices, got, tt.want)
			}
		})
	}
}
