package stack

func FindTotalWays(nums []int, idx, curSum, target int) int {
	if idx == len(nums) && curSum == target {
		return 1
	}

	if idx >= len(nums) {
		return 0
	}
	return FindTotalWays(nums, idx+1, curSum+nums[idx], target) + FindTotalWays(nums, idx+1, curSum-nums[idx], target)
}

func FindTargetSumWays(nums []int, target int) int {
	return FindTotalWays(nums, 0, 0, target)
}
