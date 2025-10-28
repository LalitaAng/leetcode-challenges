package problem417

/* There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean. 
The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.
The island is partitioned into a grid of square cells. You are given an m x n integer matrix heights where heights[r][c] represents 
the height above sea level of the cell at coordinate (r, c).
The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west 
if the neighboring cell's height is less than or equal to the current cell's height. Water can flow from any cell adjacent to an ocean into the ocean.
Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes that rain water can flow from cell (ri, ci) to 
both the Pacific and Atlantic oceans. */

func PacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return nil
	}

	m, n := len(heights), len(heights[0])

	pac := make([][]bool, m)
	atl := make([][]bool, m)
	for i := 0; i < m; i++ {
		pac[i] = make([]bool, n)
		atl[i] = make([]bool, n)
	}

	pacq := [][]int{}
	atlq := [][]int{}

	for i := 0; i < n; i++ {
		pac[0][i] = true
		pacq = append(pacq, []int{0, i})
	}
	for i := 0; i < m; i++ {
		if !pac[i][0] {
			pac[i][0] = true
			pacq = append(pacq, []int{i, 0})
		}
	}

	for i := 0; i < n; i++ {
		atl[m-1][i] = true
		atlq = append(atlq, []int{m - 1, i})
	}
	for i := 0; i < m; i++ {
		if !atl[i][n-1] {
			atl[i][n-1] = true
			atlq = append(atlq, []int{i, n - 1})
		}
	}

	bfs(heights, pacq, pac)
	bfs(heights, atlq, atl)

	res := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pac[i][j] && atl[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}

func bfs(heights [][]int, queue [][]int, seen [][]bool) {
	m, n := len(heights), len(heights[0])
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	head := 0
	for head < len(queue) {
		curr := queue[head]
		head++
		for _, dir := range dirs {
			x, y := curr[0]+dir[0], curr[1]+dir[1]
			if x < 0 || x >= m || y < 0 || y >= n || seen[x][y] {
				continue
			}
			if heights[x][y] >= heights[curr[0]][curr[1]] {
				seen[x][y] = true
				queue = append(queue, []int{x, y})
			}
		}
	}
}
