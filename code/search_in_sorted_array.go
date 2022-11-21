package code

func SearchInRotatedSortedArray(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	return binarySearch(nums, low, high, target)
}

func binarySearch(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if nums[mid] == target {
		return mid
	}
	left := binarySearch(nums, low, mid-1, target)
	if left != -1 {
		return left
	}
	return binarySearch(nums, mid+1, high, target)
}
