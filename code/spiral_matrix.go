package code

func spiralMatrix(matrix [][]int, rows, cols int) []int {
	k := 0
	l := 0
	result := make([]int, 0)

	for k < rows && l < cols {
		// print first row
		for i := l; i < cols; i++ {
			result = append(result, matrix[k][i])
		}
		k++
		// print last col
		for i := k; i < rows; i++ {
			result = append(result, matrix[i][cols-1])
		}
		cols--
		// print last row
		if k < rows {
			for i := cols - 1; i >= l; i-- {
				result = append(result, matrix[rows-1][i])
			}
			rows--
		}

		// print first col
		if l < cols {
			for i := rows - 1; i >= k; i-- {
				result = append(result, matrix[i][l])
			}
			l++
		}

	}
	return result

}
func FindSpiralMatrix(matrix [][]int) []int {
	return spiralMatrix(matrix, len(matrix), len(matrix[0]))
}
