package problem746

/* You are given an integer array cost where cost[i] is the cost of ith step on a staircase. 
Once you pay the cost, you can either climb one or two steps.
You can either start from the step with index 0, or the step with index 1.
Return the minimum cost to reach the top of the floor. */

func MinCostClimbingStairs(cost []int) int {
    n := len(cost)
    first, second := cost[0], cost[1]
    for i := 2; i < n; i++ {
        curr := cost[i] + min(first, second)
        first, second = second, curr
    }
    return min(first, second)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
