package contest_24_12

import "sort"

func AppendArray(nums []int) []int {
	sort.Ints(nums)
	arr := make([]int, 0)
	i := 0
	j := 1
	for len(arr) < len(nums) {
		arr = append(arr, nums[j], nums[i])
		i += 2
		j += 2
	}
	return arr
}
