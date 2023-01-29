package sorting

import "math"

// FindMaxProduct https://massivealgorithms.blogspot.com/search?q=Maximum+Product+of+Three+Numbers
func FindMaxProduct(nums []int) int {
	min1, min2 := math.MaxInt, math.MaxInt
	max1, max2, max3 := math.MinInt, math.MinInt, math.MinInt

	// result = max((min1 * min2 * max1), (max1 * max2 * max3))
	for _, num := range nums {
		// find min1, min2
		if num <= min1 {
			min2 = min1
			min1 = num
		} else if num <= min2 {
			min2 = num
		}

		// find max1, max2, max3
		if num >= max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num >= max2 {
			max3 = max2
			max2 = num
		} else if num >= max3 {
			max3 = num
		}
	}

	return int(math.Max(float64(min1*min2*max1), float64(max1*max2*max3)))
}
