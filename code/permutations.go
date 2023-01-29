package code

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
	permutation(nums, &rs, 0, len(nums)-1)
	return rs
}
