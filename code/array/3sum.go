package array

import "sort"

func Find3Sum(nums []int) [][]int {
	sort.Ints(nums)
	rs := make([][]int, 0)
	duplicateMap := make(map[[3]int]int)
	for i := 0; i < len(nums); i++ {
		second := i + 1
		third := len(nums) - 1
		for second < third {
			sum := nums[i] + nums[second] + nums[third]
			if sum > 0 {
				third--
			} else if sum < 0 {
				second++
			} else {
				if _, found := duplicateMap[[3]int{nums[i], nums[second], nums[third]}]; !found {
					duplicateMap[[3]int{nums[i], nums[second], nums[third]}] = 1
					rs = append(rs, []int{nums[i], nums[second], nums[third]})
				}
				second++
				third--
			}
		}
	}

	return rs
}
