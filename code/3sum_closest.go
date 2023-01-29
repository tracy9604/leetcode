package code

import (
	"math"
	"sort"
)

func Find3SumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minClosest := math.MaxInt16
	sum := 0
	maxSum := 0
	for i := 0; i < len(nums); i++ {
		second := i + 1
		third := len(nums) - 1
		for second < third {
			sum = nums[i] + nums[second] + nums[third]
			if sum < target {
				second++
			} else if sum > target {
				third--
			} else {
				return target
			}

			if int(math.Abs(float64(sum)-float64(target))) < minClosest {
				minClosest = int(math.Abs(float64(sum) - float64(target)))
				maxSum = sum
			}
		}
	}

	return maxSum
}
