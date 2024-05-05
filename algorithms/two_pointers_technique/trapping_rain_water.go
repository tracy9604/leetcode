package two_pointers_technique

import "math"

func TrappingRainWater(height []int) int {
	ans := 0
	for i := 1; i < len(height)-1; i++ {
		leftHeight := height[i]
		for j := 0; j < i; j++ {
			if height[j] > leftHeight {
				leftHeight = height[j]
			}
		}

		rightHeight := height[i]
		for j := i + 1; j < len(height); j++ {
			if height[j] > rightHeight {
				rightHeight = height[j]
			}
		}

		ans += int(math.Min(float64(leftHeight), float64(rightHeight))) - height[i]
	}
	return ans
}

func TrappingRainWaterV2(height []int) int {
	n := len(height)
	left, right := 0, n-1
	lMax, rMax := 0, 0

	ans := 0

	for left <= right {
		if rMax <= lMax {
			ans += int(math.Max(0, float64(rMax-height[right])))
			rMax = int(math.Max(float64(rMax), float64(height[right])))
			right--
		} else {
			ans += int(math.Max(0, float64(lMax-height[left])))
			lMax = int(math.Max(float64(lMax), float64(height[left])))
			left++
		}
	}
	return ans
}
