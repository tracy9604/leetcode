package array

import "math"

func FindDistanceValueTwoArrays(arr1 []int, arr2 []int, d int) int {
	count := 0
	for i := 0; i < len(arr1); i++ {
		isValidNumber := true
		for j := 0; j < len(arr2); j++ {
			if math.Abs(float64(arr1[i])-float64(arr2[j])) <= float64(d) {
				isValidNumber = false
				break
			}
		}
		if isValidNumber {
			count++
		}
	}
	return count
}
