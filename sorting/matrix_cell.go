package sorting

import (
	"math"
	"sort"
)

func GetMatrixCells(rows int, cols int, rCenter int, cCenter int) [][]int {
	tmp := make([][]int, 0)
	result := [][]int{{rCenter, cCenter}}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i != rCenter || j != cCenter {
				tmp = append(tmp, []int{i, j})
			}
		}
	}

	sort.SliceStable(tmp, func(i, j int) bool {
		return math.Abs(float64(tmp[i][0])-float64(rCenter))+math.Abs(float64(tmp[i][1])-float64(cCenter)) < math.Abs(float64(tmp[j][0])-float64(rCenter))+math.Abs(float64(tmp[j][1])-float64(cCenter))
	})

	result = append(result, tmp...)

	return result
}
