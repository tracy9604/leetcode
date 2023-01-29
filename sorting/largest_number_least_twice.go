package sorting

import "math"

func FindLargestNumberLeastTwice(nums []int) int {
	max := math.MinInt
	secondMax := math.MinInt
	idxMax := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			secondMax = max
			max = nums[i]
			idxMax = i
		} else if nums[i] > secondMax {
			secondMax = nums[i]
		}
	}

	if secondMax*2 <= max {
		return idxMax
	}

	return -1
}
