package sorting

import "sort"

// ArrayPairSum read at link: https://massivealgorithms.blogspot.com/2017/04/leetcode-561-array-partition-i.html
func ArrayPairSum(nums []int) int {
	sort.Ints(nums)
	result := 0

	for i := 0; i < len(nums); i += 2 {
		result += nums[i]
	}
	return result
}
