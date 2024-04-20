package sorting

import (
	"math"
	"sort"
)

func FindMinimumAbsoluteDifference(arr []int) [][]int {
	sort.Ints(arr)

	minDiff := math.MaxInt
	minDiffMap := make(map[int][][]int)
	ans := make([][]int, 0)
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff < minDiff {
			minDiff = diff
		}
		if _, found := minDiffMap[diff]; !found {
			minDiffMap[diff] = make([][]int, 0)
		}

		minDiffMap[diff] = append(minDiffMap[diff], []int{arr[i], arr[i+1]})
	}

	if val, found := minDiffMap[minDiff]; found {
		ans = val
	}

	return ans
}
