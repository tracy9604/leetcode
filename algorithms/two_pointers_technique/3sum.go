package two_pointers_technique

import "sort"

func Find3Sum(nums []int) [][]int {
	sort.Ints(nums)

	n := len(nums)
	sum := 0
	res := make([][]int, 0)
	dupMaps := make(map[[3]int]bool)

	for i := 0; i < n; i++ {
		left := i + 1
		right := n - 1

		subSum := sum - nums[i]
		for left < right {
			if nums[left]+nums[right] == subSum {
				key := [3]int{nums[i], nums[left], nums[right]}
				if _, found := dupMaps[key]; !found {
					res = append(res, []int{nums[i], nums[left], nums[right]})
					dupMaps[[3]int{nums[i], nums[left], nums[right]}] = true
				}
				left++
				right--
			} else if nums[left]+nums[right] < subSum {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
