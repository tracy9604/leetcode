package contest12_11

import (
	"fmt"
	"math"
)

func FindPairs(nums []int) int {
	max := math.MinInt
	pairs := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			fmt.Println(nums[i])
			fmt.Println(nums[j])
			if math.Abs(float64(nums[i])-float64(nums[j])) <= math.Min(float64(nums[i]), float64(nums[j])) && (nums[i]^nums[j]) > max {
				max = nums[i] ^ nums[j]
				pairs = append(pairs, []int{nums[i], nums[j]})
			}
		}
	}
	fmt.Println(pairs)
	return max
}
