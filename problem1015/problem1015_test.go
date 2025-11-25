package problem1015

import "testing"

func TestSmallestRepunitDivByK(t *testing.T) {
	tests := []struct {
		name string
		k    int
		want int
	}{
		{
			name: "k_is_even",
			k:    2,
			want: -1,
		},
		{
			name: "k_ends_with_5",
			k:    5,
			want: -1,
		},
		{
			name: "k_1",
			k:    1,
			want: 1, 
		},
		{
			name: "k_3",
			k:    3,
			want: 3,
		},
		{
			name: "k_7",
			k:    7,
			want: 6, 
		},
		{
			name: "k_11",
			k:    11,
			want: 2, 
		},
		{
			name: "k_17",
			k:    17,
			want: 16,
		},
		{
			name: "k_large_prime",
			k:    19,
			want: 18,
		},
		{
			name: "k_41",
			k:    41,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SmallestRepunitDivByK(tt.k)
			if got != tt.want {
				t.Errorf("SmallestRepunitDivByK(%d) = %d; want %d", tt.k, got, tt.want)
			}
		})
	}
}
