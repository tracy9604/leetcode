package contest_21_11

func ShiftMatrix(mat [][]int, k int) bool {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			idx := j + k
			if j+k >= len(mat[i]) {
				idx = j + k - len(mat[i])
			}
			mat[i][j], mat[i][idx] = mat[i][idx], mat[i][j]
		}
	}
	return false
}
