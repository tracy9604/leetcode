package sorting

func SortColor(nums []int) {
	for i := 0; i < len(nums); i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}

func SortColorV2(nums []int) {
	if len(nums) <= 1 {
		return
	}
	maxNum := nums[0]
	countMap := make(map[int]int)
	for _, item := range nums {
		if item > maxNum {
			maxNum = item
		}
		countMap[item]++
	}

	count := make([]int, maxNum+1)
	for i := 0; i <= maxNum; i++ {
		if val, found := countMap[i]; found && val > 0 {
			count[i] = val
		} else {
			count[i] = 0
		}
	}

	startIdx := 0
	for idx, c := range count {
		count[idx] = startIdx
		startIdx += c
	}

	sortedList := make([]int, len(nums))
	for _, item := range nums {
		sortedList[count[item]] = item
		count[item]++
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = sortedList[i]
	}
}
