package array

func MaximumSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	sum := 0

	// if sum is negative, reset sum to 0 because 0 always greater than negative numbers
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
