package sorting

func FindMissingNumber(nums []int) int {
	lenNums := len(nums)
	sum := lenNums * (lenNums + 1) / 2
	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
	}

	return sum
}
