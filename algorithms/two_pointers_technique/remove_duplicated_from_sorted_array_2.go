package two_pointers_technique

func RemoveDuplicatedFromSortedArray2(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	k := 0

	for i := 0; i < len(nums); i++ {
		if k < 2 || nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
