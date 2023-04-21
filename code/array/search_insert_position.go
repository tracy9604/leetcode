package array

func SearchInsertPosition(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		} else if i == len(nums)-1 {
			return i + 1
		}
	}
	return 0
}

func SearchInsertPositionLog(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[left] < target {
			left++
		} else {
			right--
		}
	}
	return left
}
