package problem2257

/* You are given two integers m and n representing a 0-indexed m x n grid. 
You are also given two 2D integer arrays guards and walls where guards[i] = [rowi, coli] and 
walls[j] = [rowj, colj] represent the positions of the ith guard and jth wall respectively.
A guard can see every cell in the four cardinal directions (north, east, south, or west) starting from their position 
unless obstructed by a wall or another guard. A cell is guarded if there is at least one guard that can see it.
Return the number of unoccupied cells that are not guarded. */

func CountUnguarded(m int, n int, guards [][]int, walls [][]int) int {
    grid := make([][]int, m)
    for i := range m {
        grid[i] = make([]int, n)
    }
    for _, w := range walls {
        grid[w[0]][w[1]] = 2
    }
    for _, g := range guards {
        grid[g[0]][g[1]] = 1
    }

    dirs := [4][2]int{{-1,0},{1,0},{0,-1},{0,1}}
    for _, g := range guards {
        r, c := g[0], g[1]
        for _, d := range dirs {
            nr, nc := r + d[0], c + d[1]
            for nr >= 0 && nr < m && nc >= 0 && nc < n {
                if grid[nr][nc] == 2 || grid[nr][nc] == 1 {
                    break
                }
                if grid[nr][nc] == 0 {
                    grid[nr][nc] = 3
                }
                nr += d[0]
                nc += d[1]
            }
        }
    }

    ans := 0
    for i := range m {
        for j := range n {
            if grid[i][j] == 0 {
                ans++
            }
        }
    }
    return ans
}
