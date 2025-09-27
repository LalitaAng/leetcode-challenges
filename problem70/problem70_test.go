package problem70

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "n = 1",
			n:    1,
			want: 1,
		},
		{
			name: "n = 2",
			n:    2,
			want: 2,
		},
		{
			name: "n = 3",
			n:    3,
			want: 3, 
		},
		{
			name: "n = 4",
			n:    4,
			want: 5, 
		},
		{
			name: "n = 5",
			n:    5,
			want: 8,
		},
		{
			name: "n = 10",
			n:    10,
			want: 89,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClimbStairs(tt.n)
			if got != tt.want {
				t.Errorf("ClimbStairs(%d) = %d; want %d", tt.n, got, tt.want)
			}
		})
	}
}
