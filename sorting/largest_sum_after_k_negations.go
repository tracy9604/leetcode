package sorting

func LargestSumAfterKNegations(nums []int, k int) int {
	for i := 0; i < k; i++ {
		idxMin := 0
		for j := 0; j < len(nums); j++ {
			if nums[idxMin] > nums[j] {
				idxMin = j
			}
		}
		nums[idxMin] = -1 * nums[idxMin]
	}

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}
