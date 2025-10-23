package problem73

// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's. You must do it in place.

func SetZeroes(matrix [][]int) {
    m := len(matrix)
    if m == 0 || len(matrix[0]) == 0 {
        return
    }
    n := len(matrix[0])

    firstRowZero, firstColZero := false, false

    for j := 0; j < n; j++ {
        if matrix[0][j] == 0 {
            firstRowZero = true
            break
        }
    }

    for i := 0; i < m; i++ {
        if matrix[i][0] == 0 {
            firstColZero = true
            break
        }
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }

    for i := 1; i < m; i++ {
        if matrix[i][0] == 0 {
            for j := 1; j < n; j++ {
                matrix[i][j] = 0
            }
        }
    }

    for j := 1; j < n; j++ {
        if matrix[0][j] == 0 {
            for i := 1; i < m; i++ {
                matrix[i][j] = 0
            }
        }
    }

    if firstRowZero {
        for j := 0; j < n; j++ {
            matrix[0][j] = 0
        }
    }

    if firstColZero {
        for i := 0; i < m; i++ {
            matrix[i][0] = 0
        }
    }
}
