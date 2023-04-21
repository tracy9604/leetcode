package array

import (
	"sort"
)

func Find4Sum(nums []int, target int) [][]int {
	sort.Ints(nums)
	mapDuplicate := make(map[[4]int]int)
	rs := make([][]int, 0)

	if len(nums) < 4 {
		return rs
	}

	for i := 0; i < len(nums); i++ {
		for second := i + 1; second < len(nums)-2; second++ {
			third := second + 1
			fourth := len(nums) - 1
			for third < fourth {
				sum := nums[i] + nums[second] + nums[third] + nums[fourth]
				if sum < target {
					third++
				} else if sum > target {
					fourth--
				} else {
					// check duplicate
					if _, found := mapDuplicate[[4]int{nums[i], nums[second], nums[third], nums[fourth]}]; !found {
						rs = append(rs, []int{nums[i], nums[second], nums[third], nums[fourth]})
						mapDuplicate[[4]int{nums[i], nums[second], nums[third], nums[fourth]}]++
					}
					third++
					fourth--
				}
			}
		}
	}

	return rs
}
