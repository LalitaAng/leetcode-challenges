package problem70

/* You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top? */

func ClimbStairs(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }

    oneStepBefore := 2
    twoStepsBefore := 1
    totalWays := 0

    for i := 3; i <= n; i++ {
        totalWays = oneStepBefore + twoStepsBefore
        twoStepsBefore = oneStepBefore
        oneStepBefore = totalWays
    }

    return totalWays
}
