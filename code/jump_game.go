package code

import "math"

func IsCanJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if max < i {
			return false
		}
		if max >= len(nums)-1 {
			return true
		}
		max = int(math.Max(float64(i+nums[i]), float64(max)))
	}

	return false
}
