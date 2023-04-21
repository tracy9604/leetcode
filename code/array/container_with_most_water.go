package array

import "math"

func FindMaxArea(height []int) int {
	maxArea := 0
	i := 0
	j := len(height) - 1
	for j > i {
		maxArea = int(math.Max(float64(maxArea), float64(j-i)*math.Min(float64(height[i]), float64(height[j]))))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxArea
}
