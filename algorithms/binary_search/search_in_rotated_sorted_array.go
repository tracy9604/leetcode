package binary_search

func findPilot(nums []int, left, right int) int {
	if left > right {
		return -1
	}

	if left == right {
		return left
	}

	mid := (left + right) / 2
	if mid > left && nums[mid-1] > nums[mid] {
		return mid - 1
	}
	if mid < right && nums[mid] > nums[mid+1] {
		return mid
	}
	if nums[left] >= nums[mid] {
		return findPilot(nums, left, mid-1)
	}
	return findPilot(nums, mid+1, right)
}

func binarySearchRotatedSortedArray(nums []int, left, right, target int) int {
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func SearchInRotatedSortedArray(nums []int, target int) int {
	pilot := findPilot(nums, 0, len(nums)-1)

	if pilot == -1 {
		pilot = 0
	}

	if nums[pilot] == target {
		return pilot
	}
	leftRs := binarySearchRotatedSortedArray(nums, 0, pilot, target)
	if leftRs == -1 {
		return binarySearchRotatedSortedArray(nums, pilot+1, len(nums)-1, target)
	}
	return leftRs
}

func SearchInRotatedSortedArrayV2Recursion(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if nums[mid] == target {
		return mid
	}

	if nums[left] < nums[mid] {
		leftRs := SearchInRotatedSortedArrayV2Recursion(nums, left, mid-1, target)
		if leftRs == -1 {
			return SearchInRotatedSortedArrayV2Recursion(nums, mid+1, right, target)
		}
		return leftRs
	} else {
		leftRs := SearchInRotatedSortedArrayV2Recursion(nums, mid+1, right, target)
		if leftRs == -1 {
			return SearchInRotatedSortedArrayV2Recursion(nums, left, mid-1, target)
		}
		return leftRs
	}
}

func SearchInRotatedSortedArrayV2(nums []int, target int) int {
	return SearchInRotatedSortedArrayV2Recursion(nums, 0, len(nums)-1, target)
}
