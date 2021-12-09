package code

func RemoveDuplicated(nums []int) int {
	counter := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i - 1] {
			nums[counter] = nums[i]
			counter++
		}
	}
	return counter
}
