package arrays

func SortArrayByParity(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for nums[left]%2 == 0 && left < right {
			left++
		}
		for nums[right]%2 != 0 && left < right {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return nums
}
