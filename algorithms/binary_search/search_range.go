package binary_search

func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	result := make([]int, 2)

	left := 0
	right := len(nums) - 1

	for left+1 < right {
		mid := left + (right-left)/2
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid
		}
	}
	if nums[left] == target {
		result[0] = left
	} else if nums[right] == target {
		result[0] = right
	} else {
		result[0] = -1
	}

	left = 0
	right = len(nums) - 1

	for left+1 < right {
		mid := left + (right-left)/2
		if target >= nums[mid] {
			left = mid
		} else {
			right = mid
		}
	}
	if nums[right] == target {
		result[1] = right
	} else if nums[left] == target {
		result[1] = left
	} else {
		result[1] = -1
	}
	return result
}
