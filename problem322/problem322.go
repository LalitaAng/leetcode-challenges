package problem322

/* You are given an integer array coins representing coins of different denominations and 
an integer amount representing a total amount of money.
Return the fewest number of coins that you need to make up that amount. 
If that amount of money cannot be made up by any combination of the coins, return -1.
You may assume that you have an infinite number of each kind of coin.*/

func CoinChange(coins []int, amount int) int {
  dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 { 
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}
