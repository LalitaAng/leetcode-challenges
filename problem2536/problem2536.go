package problem2536

/* You are given a positive integer n, indicating that we initially have an n x n 0-indexed integer matrix mat filled with zeroes.
You are also given a 2D integer array query. For each query[i] = [row1i, col1i, row2i, col2i], you should do the following operation:
Add 1 to every element in the submatrix with the top left corner (row1i, col1i) and the bottom right corner (row2i, col2i). 
That is, add 1 to mat[x][y] for all row1i <= x <= row2i and col1i <= y <= col2i.
Return the matrix mat after performing every query. */

func RangeAddQueries(n int, queries [][]int) [][]int {
    matrix := make([][]int, n)
    for i := range matrix {
        matrix[i] = make([]int, n)
    }

    for _, query := range queries {
        rowStart, colStart := query[0], query[1]
        rowEnd, colEnd := query[2], query[3]
        
        matrix[rowStart][colStart] += 1
        
        if colEnd+1 < n {
            matrix[rowStart][colEnd+1] -= 1
        }
        if rowEnd+1 < n {
            matrix[rowEnd+1][colStart] -= 1
        }
        if rowEnd+1 < n && colEnd+1 < n {
            matrix[rowEnd+1][colEnd+1] += 1
        }
    }
    
    for row := 1; row < n; row++ {
        for col := 0; col < n; col++ {
            matrix[row][col] += matrix[row-1][col]
        }
    }
    
    for row := 0; row < n; row++ {
        for col := 1; col < n; col++ {
            matrix[row][col] += matrix[row][col-1]
        }
    }
    
    return matrix
}
