package two_pointers_technique

func MoveZeroes(nums []int) {
	i := -1

	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
