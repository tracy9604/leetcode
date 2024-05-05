package two_pointers_technique

func RemoveDuplicatedFromSortedArray(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	j := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[j] = nums[i]
			j++
		}

	}
	nums[j] = nums[len(nums)-1]
	return j + 1
}
