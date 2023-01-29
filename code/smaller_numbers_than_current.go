package code

import "sort"

func SmallerNumbersThanCurrent(nums []int) []int {
	// remove duplicate
	result := make([]int, len(nums))
	arr := make([]int, len(nums))
	copy(arr, nums)
	tmpMap := make(map[int]int)

	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		if _, found := tmpMap[arr[i]]; !found {
			tmpMap[arr[i]] = i
		}
	}

	for i := 0; i < len(nums); i++ {
		if val, found := tmpMap[nums[i]]; found {
			result[i] = val
		}
	}

	return result
}
