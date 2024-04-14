package binary_search

func BinarySearchRecursion(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if nums[mid] == target {
		return mid
	}

	if target < nums[mid] {
		return BinarySearchRecursion(nums, left, mid, target)
	}

	return BinarySearchRecursion(nums, mid+1, right, target)
}

func BinarySearch(nums []int, target int) int {
	return BinarySearchRecursion(nums, 0, len(nums)-1, target)
}

func BinarySearchV2(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		}
	}
	return -1
}
