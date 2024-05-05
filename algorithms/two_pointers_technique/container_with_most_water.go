package two_pointers_technique

import "math"

func FindContainerWithMostWater(height []int) int {
	rs := 0
	left := 0
	right := len(height) - 1

	for left < right {
		minHeight := int(math.Min(float64(height[left]), float64(height[right])))
		tmp := (right - left) * minHeight
		if rs < tmp {
			rs = tmp
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return rs
}
