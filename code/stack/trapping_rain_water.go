package stack

import "math"

// Brute force
// Traverse the array from start to end:
// 		For every element:
// 			Traverse the array from start to that index and find the maximum height (a) and
// 			Traverse the array from the current index to the end, and find the maximum height (b).
// The amount of water that will be stored in this column is min(a,b) – array[i], add this value to the total amount of water stored
// Print the total amount of water stored.

func FindTrappingRainWater(height []int) int {
	ans := 0
	left, right := 0, 0
	for i := 0; i < len(height); i++ {
		left = height[i]
		for j := 0; j < i; j++ {
			left = int(math.Max(float64(left), float64(height[j])))
		}

		right = height[i]
		for j := i + 1; j < len(height); j++ {
			right = int(math.Max(float64(right), float64(height[j])))
		}
		ans += int(math.Min(float64(left), float64(right))) - height[i]
	}
	return ans
}

func FindTrappingRainWater2(height []int) int {
	if len(height) == 0 {
		return 0
	}
	ans := 0
	left, right := make([]int, len(height)), make([]int, len(height))
	max := height[0]
	for i := 0; i < len(height); i++ {
		max = int(math.Max(float64(max), float64(height[i])))
		left[i] = max
	}

	max = height[len(height)-1]
	for j := len(height) - 1; j >= 0; j-- {
		max = int(math.Max(float64(max), float64(height[j])))
		right[j] = max
	}

	for i := 0; i < len(height); i++ {
		ans += int(math.Min(float64(left[i]), float64(right[i]))) - height[i]
	}
	return ans
}

func FindTrappingRainWater3(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	ans := 0
	s := IntegerStack{}
	var width int
	for i := 0; i < len(height); i++ {
		for !s.isEmpty() && height[s.top()] < height[i] {
			val, _ := s.pop()
			if s.isEmpty() {
				break
			}
			width = i - s.top() - 1
			ans += (int(math.Min(float64(height[i]), float64(height[s.top()]))) - height[val]) * width
		}
		s.push(i)
	}

	return ans
}
