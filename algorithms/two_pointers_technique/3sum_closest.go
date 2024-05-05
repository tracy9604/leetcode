package two_pointers_technique

import (
	"math"
	"sort"
)

func Find3SumClosest(nums []int, target int) int {
	sort.Ints(nums)

	n := len(nums)
	diff := math.MaxInt
	var ans int

	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return target
			}
			if int(math.Abs(float64(sum-target))) < diff {
				diff = int(math.Abs(float64(sum - target)))
				ans = sum
			}

			if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return ans
}
