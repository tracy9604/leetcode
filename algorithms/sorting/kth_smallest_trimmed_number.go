package sorting

import (
	"math"
	"sort"
	"strconv"
)

type NumsMap struct {
	originalIdx int
	currNum     int
	originalNum string
}

func FindKthSmallestTrimmedNumber(nums []string, queries [][]int) []int {
	ans := make([]int, 0)

	for _, query := range queries {
		placeVal := int(math.Pow(10, float64(query[1])))
		tmpNums := make([]NumsMap, len(nums))
		for i, item := range nums {
			numItem, _ := strconv.Atoi(item)
			tmpNums[i] = NumsMap{
				originalIdx: i,
				currNum:     numItem % placeVal,
				originalNum: item,
			}
		}
		sort.Slice(tmpNums, func(i, j int) bool {
			return tmpNums[i].currNum < tmpNums[j].currNum
		})
		ans = append(ans, tmpNums[query[0]-1].originalIdx)
	}

	return ans
}
