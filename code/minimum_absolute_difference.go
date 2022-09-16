package code

import (
	"math"
	"sort"
)

func FindMinimumAbsDifference(arr []int) [][]int {
	if len(arr) == 2 {
		return [][]int{{arr[0], arr[1]}}
	}

	result := make([][]int, 0)
	sort.Ints(arr)

	minAbs := int(math.Abs(float64(arr[1] - arr[0])))
	l := len(arr)
	for idx, num := range arr {
		if idx+1 < l {
			if arr[idx+1]-num == minAbs {
				result = append(result, []int{num, arr[idx+1]})
			} else if int(math.Abs(float64(arr[idx+1]-num))) < minAbs {
				result = [][]int{{num, arr[idx+1]}}
				minAbs = int(math.Abs(float64(arr[idx+1] - num)))
			}
		}
	}

	return result
}
