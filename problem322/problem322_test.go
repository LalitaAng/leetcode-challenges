package problem322

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins   []int
		amount  int
		want    int
	}{
		{[]int{1, 2, 5}, 11, 3},   
		{[]int{2}, 3, -1},         
		{[]int{1}, 0, 0},         
		{[]int{1}, 2, 2},         
		{[]int{186, 419, 83, 408}, 6249, 20}, 
		{[]int{2, 5, 10}, 7, 2},  
		{[]int{3, 4}, 6, 2},       
	}

	for _, tt := range tests {
		got := CoinChange(tt.coins, tt.amount)
		if got != tt.want {
			t.Errorf("coinChange(%v, %d) = %d; want %d",
				tt.coins, tt.amount, got, tt.want)
		}
	}
}
