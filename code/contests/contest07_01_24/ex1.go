package contest07_01_24

import "math"

func FindAreaOfMaxDiagonal(dimensions [][]int) int {
	max := 0
	var longestDiagonal float64

	for i := 0; i < len(dimensions); i++ {
		length := dimensions[i][0]
		width := dimensions[i][1]
		diagonal := math.Sqrt(float64(length*length + width*width))
		if diagonal > longestDiagonal {
			longestDiagonal = diagonal
			max = length * width
		} else if diagonal == longestDiagonal && max < length*width {
			max = length * width
		}

	}
	return max
}
