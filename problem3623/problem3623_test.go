package problem3623

import "testing"

func TestCountTrapezoids(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name:   "empty input",
			points: [][]int{},
			want:   0,
		},
		{
			name:   "single point",
			points: [][]int{{0, 0}},
			want:   0,
		},
		{
			name:   "two points same height",
			points: [][]int{{0, 1}, {2, 1}},
			want:   0,
		},
		{
			name:   "two heights but only one pair each",
			points: [][]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
			want:   1,
		},
		{
			name:   "three heights one pair each",
			points: [][]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {0, 2}, {1, 2}},
			want:   3, 
		},
		{
			name: "uneven points per height",
			points: [][]int{
				{0, 0}, {1, 0}, {2, 0}, 
				{0, 1}, {1, 1},        
			},
			want: 3, 
		},
		{
			name: "multiple heights with varying counts",
			points: [][]int{
				{0, 0}, {1, 0}, {2, 0},      
				{0, 1}, {1, 1},              
				{0, 2}, {1, 2}, {2, 2}, {3, 2},
			},
			want: 27,
		},
		{
			name: "negative coordinates",
			points: [][]int{
				{0, -1}, {1, -1},
				{0, -2}, {1, -2},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountTrapezoids(tt.points)
			if got != tt.want {
				t.Errorf("CountTrapezoids(%v) = %d; want %d", tt.points, got, tt.want)
			}
		})
	}
}
