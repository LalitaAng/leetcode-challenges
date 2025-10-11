package problem200

func NumIslands(islandMap [][]byte) int {
    islandCount := 0

    for row := 0; row < len(islandMap); row++ {
        for col := 0; col < len(islandMap[row]); col++ {
            if islandMap[row][col] == '1' {
                islandCount++
                exploreIslandBFS(islandMap, row, col)
            }
        }
    }

    return islandCount
}

func exploreIslandBFS(islandMap [][]byte, startRow, startCol int) {
    queue := [][2]int{{startRow, startCol}}
    islandMap[startRow][startCol] = '2'

    directions := [4][2]int{
        {1, 0},  
        {-1, 0}, 
        {0, 1},  
        {0, -1}, 
    }

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        for _, dir := range directions {
            nextRow, nextCol := current[0]+dir[0], current[1]+dir[1]

            if nextRow >= 0 && nextRow < len(islandMap) &&
                nextCol >= 0 && nextCol < len(islandMap[nextRow]) &&
                islandMap[nextRow][nextCol] == '1' {

                islandMap[nextRow][nextCol] = '2'
                queue = append(queue, [2]int{nextRow, nextCol})
            }
        }
    }
}
