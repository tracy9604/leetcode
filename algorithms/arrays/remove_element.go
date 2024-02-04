package arrays

func RemoveElement(nums []int, val int) int {
	i, j, k := 0, len(nums)-1, 0
	for i < len(nums) && j >= i {
		if nums[i] == val {
			k++
			if nums[j] != val {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
			j--
			continue
		}
		i++
	}
	return len(nums) - k
}
