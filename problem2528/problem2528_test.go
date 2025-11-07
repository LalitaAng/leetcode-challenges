package problem2528

import "testing"

func TestMaxPower(t *testing.T) {
	tests := []struct {
		name      string
		stations  []int
		r, k      int
		want      int64
	}{
		{
			name:     "basic example 1",
			stations: []int{1, 2, 4, 5, 0},
			r:        1,
			k:        2,
			want:     5,
		},
		{
			name:     "basic example 2",
			stations: []int{4, 4, 4, 4},
			r:        0,
			k:        10,
			want:     6, // fixed
		},
		{
			name:     "no additional stations (k = 0)",
			stations: []int{1, 1, 1, 1},
			r:        1,
			k:        0,
			want:     2,
		},
		{
			name:     "large range covers all stations",
			stations: []int{2, 3, 5},
			r:        2,
			k:        1,
			want:     11, // fixed
		},
		{
			name:     "single station only",
			stations: []int{5},
			r:        0,
			k:        5,
			want:     10,
		},
		{
			name:     "increasing pattern with range 1",
			stations: []int{0, 1, 2, 3, 4, 5},
			r:        1,
			k:        3,
			want:     4, // fixed
		},
		{
			name:     "all zeros, big k",
			stations: []int{0, 0, 0, 0, 0},
			r:        1,
			k:        10,
			want:     5, // fixed
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxPower(tt.stations, tt.r, tt.k)
			if got != tt.want {
				t.Errorf("MaxPower(%v, %d, %d) = %d, want %d",
					tt.stations, tt.r, tt.k, got, tt.want)
			}
		})
	}
}
