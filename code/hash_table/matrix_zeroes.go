package hash_table

import (
	"fmt"
	"math"
)

func SetZeroes(matrix [][]int) {
	rsMap := map[string][]int{
		"row": make([]int, 0),
		"col": make([]int, 0),
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				// set matrix[i]
				rsMap["row"] = append(rsMap["row"], i)
				rsMap["col"] = append(rsMap["col"], j)
			}
		}
	}

	for _, row := range rsMap["row"] {
		for j := 0; j < len(matrix[row]); j++ {
			matrix[row][j] = 0
		}
	}

	for _, col := range rsMap["col"] {
		for i := 0; i < len(matrix); i++ {
			matrix[i][col] = 0
		}
	}
}

func SetZeroesV2(matrix [][]int) {
	rows := len(matrix)
	cols := 0
	for i := 0; i < len(matrix); i++ {
		cols = len(matrix[i])
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				idn := i - 1
				for idn >= 0 {
					if matrix[idn][j] != 0 {
						matrix[idn][j] = math.MaxInt
					}
					idn--
				}

				idn = i + 1
				for idn < rows {
					if matrix[idn][j] != 0 {
						matrix[idn][j] = math.MaxInt
					}
					idn++
				}

				idn = j - 1
				for idn >= 0 {
					if matrix[i][idn] != 0 {
						matrix[i][idn] = math.MaxInt
					}
					idn--
				}

				idn = j + 1
				for idn < cols {
					if matrix[i][idn] != 0 {
						matrix[i][idn] = math.MaxInt
					}
					idn++
				}
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == math.MaxInt {
				matrix[i][j] = 0
			}
		}
	}
	fmt.Println(matrix)
}
