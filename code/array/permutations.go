package array

import "sort"

// backtracking
func permutation(nums []int, rs *[][]int, left, right int) {
	if left == right {
		ansItem := make([]int, 0)
		ansItem = append(ansItem, nums...)
		*rs = append(*rs, ansItem)
	} else {
		for i := left; i <= right; i++ {
			// Swapping done
			nums[i], nums[left] = nums[left], nums[i]

			// Recursion called
			permutation(nums, rs, left+1, right)

			// backtrack
			nums[i], nums[left] = nums[left], nums[i]
		}
	}
}

func FindPermutations(nums []int) [][]int {
	rs := make([][]int, 0)
	result := make([][]int, 0)
	sort.Ints(nums)
	permutation(nums, &rs, 0, len(nums)-1)
	dupMap := make(map[[8]int]int)
	for _, item := range rs {
		itemArr := [8]int{}
		for i, it := range item {
			itemArr[i] = it
		}
		if _, found := dupMap[itemArr]; !found {
			dupMap[itemArr] = 1
			result = append(result, append([]int{}, item...))
		}
	}
	return result
}

// concatenating subarrays
func permuteSubArrays(nums, subArray []int, res *[][]int) {
	if len(nums) == 0 {
		ansItem := make([]int, 0)
		ansItem = append(ansItem, subArray...)
		*res = append(*res, ansItem)
	} else {
		for i := 0; i < len(nums); i++ {
			item := nums[i]
			left := nums[0:i]
			right := nums[i+1:]
			rest := append(left, right...)
			permuteSubArrays(rest, append(subArray, item), res)
		}
	}
}

func FindPermutationsBySubArrays(nums []int) [][]int {
	rs := make([][]int, 0)
	subArray := make([]int, 0)
	permuteSubArrays(nums, subArray, &rs)
	return rs
}
