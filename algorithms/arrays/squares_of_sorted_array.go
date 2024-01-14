package arrays

func SortedSquares(nums []int) []int {
	k := 0
	for ; k < len(nums); k++ {
		if nums[k] > 0 {
			break
		}
	}

	left := k - 1 // last index of first half
	right := k    // first index of second half
	idx := 0
	temp := make([]int, len(nums))

	//
	for left >= 0 && right < len(nums) {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			temp[idx] = nums[left] * nums[left]
			left--
		} else {
			temp[idx] = nums[right] * nums[right]
			right++
		}
		idx++
	}

	for left >= 0 {
		temp[idx] = nums[left] * nums[left]
		idx++
		left--
	}

	for right < len(nums) {
		temp[idx] = nums[right] * nums[right]
		right++
		idx++
	}

	return temp
}
