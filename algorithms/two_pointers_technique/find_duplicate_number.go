package two_pointers_technique

func FindDuplicateNumber(nums []int) int {
	slow, fast := nums[0], nums[0]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
