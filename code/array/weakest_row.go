package array

import "sort"

func FindWeakestRows(mat [][]int, k int) []int {
	score := make([]int, 0)
	rows := len(mat)

	for i := 0; i < len(mat); i++ {
		j := 0
		for ; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				break
			}
		}

		score = append(score, j*rows+i)
	}

	sort.Ints(score)

	// get rowIndex
	for i := 0; i < len(score); i++ {
		score[i] = score[i] % rows
	}

	return score[:k]
}
