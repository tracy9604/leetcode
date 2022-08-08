package sorting

import "sort"

func SortByParity(nums []int) []int {
	result := make([]int, len(nums))
	evenIdx := 0
	oddIdx := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			result[evenIdx] = nums[i]
			evenIdx++
		} else {
			result[oddIdx] = nums[i]
			oddIdx--
		}
	}
	return result
}

func SortByParity2(nums []int) []int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i]%2 < nums[j]%2
	})
	return nums
}

func SortByParityII(nums []int) []int {
	even := 0
	odd := 1

	for even < len(nums) && odd < len(nums) {
		if nums[even]%2 != 0 {
			nums[even], nums[odd] = nums[odd], nums[even]
			odd += 2
		} else {
			even += 2
		}
	}

	return nums
}
