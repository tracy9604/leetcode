package array

import "math"

func FindMaxConsecutive(nums []int) int {
	count := 0
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count = 0
		} else {
			count++
			result = int(math.Max(float64(result), float64(count)))
		}
	}
	return result
}
