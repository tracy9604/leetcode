package binary_search

func FindPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] >= nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func FindPeakElementV2(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] >= nums[mid-1] && nums[mid] >= nums[mid+1] {
			return mid
		}
		if nums[mid] < nums[mid-1] {
			right = mid
		} else if nums[mid] < nums[mid+1] {
			left = mid
		}
	}
	if nums[left] < nums[right] {
		return right
	}
	return left
}
