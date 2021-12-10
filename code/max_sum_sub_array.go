package code

func MaximumSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	return max
}
