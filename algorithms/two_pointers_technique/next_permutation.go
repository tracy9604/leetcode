package two_pointers_technique

func FindNextPermutation(nums []int) {
	// find the first element from the right which is smaller than the element next to it
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		// find the smallest from the right which is greater than the first element - the second element
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}

		// swap the first element and the second element
		nums[i], nums[j] = nums[j], nums[i]
	}

	// reverse the elements after index i
	left := i + 1
	right := len(nums) - 1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
