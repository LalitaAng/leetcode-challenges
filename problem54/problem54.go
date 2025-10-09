package problem54

// Given an m x n matrix, return all elements of the matrix in spiral order.

func SpiralOrder(matrix [][]int) []int {
     if len(matrix) == 0 || len(matrix[0]) == 0 {
        return []int{}
    }

    top, bottom := 0, len(matrix)-1
    left, right := 0, len(matrix[0])-1
    var result []int

    for top <= bottom && left <= right {
        for i := left; i <= right; i++ {
            result = append(result, matrix[top][i])
        }
        top++

        for i := top; i <= bottom; i++ {
            result = append(result, matrix[i][right])
        }
        right--

        if top <= bottom {
            for i := right; i >= left; i-- {
                result = append(result, matrix[bottom][i])
            }
            bottom--
        }

        if left <= right {
            for i := bottom; i >= top; i-- {
                result = append(result, matrix[i][left])
            }
            left++
        }
    }

    return result
}
