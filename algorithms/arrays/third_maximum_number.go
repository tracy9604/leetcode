package arrays

import "math"

func FindThirdMaximumNumber(nums []int) int {
	if len(nums) < 3 {
		max := nums[0]
		for i := 1; i < len(nums); i++ {
			if max < nums[i] {
				max = nums[i]
			}
		}
		return max
	}

	dup := make(map[int]bool)
	max1, max2, max3 := math.MinInt, math.MinInt, math.MinInt
	for i := 0; i < len(nums); i++ {
		if _, found := dup[nums[i]]; !found {
			dup[nums[i]] = true
			if nums[i] >= max1 {
				max3 = max2
				max2 = max1
				max1 = nums[i]
			} else if nums[i] >= max2 {
				max3 = max2
				max2 = nums[i]
			} else if nums[i] >= max3 {
				max3 = nums[i]
			}
		}

	}
	if max3 != math.MinInt {
		return max3
	}
	return max1
}
