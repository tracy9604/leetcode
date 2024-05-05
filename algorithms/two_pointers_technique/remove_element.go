package two_pointers_technique

func RemoveElement(nums []int, val int) int {
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}

func RemoveElementV2(nums []int, val int) int {
	left := 0
	right := len(nums)

	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return right
}
