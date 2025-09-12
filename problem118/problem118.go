package problem118

// Given an integer numRows, return the first numRows of Pascal's triangle.

func Generate(numRows int) [][]int {
    triangle := [][]int{}
    for i := 0; i < numRows; i++ {
        row := make([]int, i+1)
        row[0], row[i] = 1, 1
        for j := 1; j < i; j++ {
            row[j] = triangle[i-1][j-1] + triangle[i-1][j]
        }
        triangle = append(triangle, row)
    }
    return triangle
}
