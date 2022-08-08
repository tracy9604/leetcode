package sorting

import "math"

func FindSquareOfSortedArray(nums []int) []int {
	n := len(nums)
	j := n - 1
	for p := n - 1; p >= 0; p-- {
		if math.Abs(float64(nums[0])) > math.Abs(float64(nums[j])) {
			nums[0], nums[j] = nums[j], nums[0]
		}
		nums[j] *= nums[j]
		j--
	}
	return nums
}
