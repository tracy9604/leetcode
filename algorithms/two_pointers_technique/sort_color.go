package two_pointers_technique

func SortColorByCountingSort(nums []int) {
	if len(nums) == 0 {
		return
	}

	maxNumber := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxNumber {
			maxNumber = nums[i]
		}
	}

	counts := make([]int, maxNumber+1)
	for _, item := range nums {
		counts[item]++
	}

	startingIndex := 0
	for idx, count := range counts {
		counts[idx] = startingIndex
		startingIndex += count
	}

	sortedArray := make([]int, len(nums))
	for _, item := range nums {
		sortedArray[counts[item]] = item
		counts[item]++
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = sortedArray[i]
	}
}
