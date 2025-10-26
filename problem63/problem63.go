package problem63

/* You are given an m x n integer array grid. There is a robot initially located at the top-left corner 
(i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). 
The robot can only move either down or right at any point in time.
An obstacle and space are marked as 1 or 0 respectively in grid. 
A path that the robot takes cannot include any square that is an obstacle.
Return the number of possible unique paths that the robot can take to reach the bottom-right corner.
The testcases are generated so that the answer will be less than or equal to 2 * 109. */

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	previous := make([]int, n)
	current := make([]int, n)
	previous[0] = 1

	for i := 0; i < m; i++ {
		for j := range current {
			current[j] = 0
		}

		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				current[j] = 0
				continue
			}

			if i == 0 && j == 0 {
				current[j] = 1
				continue
			}

			if j > 0 {
				current[j] += current[j-1]
			}
			if i > 0 {
				current[j] += previous[j]
			}
		}

		previous, current = current, previous
	}

	return previous[n-1]
}
