package sorting

import "math"

func countingSort(nums []int, placeVal int) {
	counts := make([]int, 10)

	for _, item := range nums {
		current := item / placeVal
		counts[current%10]++
	}

	startingIdx := 0
	for idx, count := range counts {
		counts[idx] = startingIdx
		startingIdx += count
	}

	sortedArr := make([]int, len(nums))
	for _, item := range nums {
		current := item / placeVal
		sortedArr[counts[current%10]] = item
		counts[current%10]++
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = sortedArr[i]
	}
}

func radixSort(nums []int) {
	maxElem := nums[0]
	for _, item := range nums {
		if item > maxElem {
			maxElem = item
		}
	}

	placeVal := 1
	for maxElem/placeVal > 0 {
		countingSort(nums, placeVal)
		placeVal *= 10
	}
}

func FindMaximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	radixSort(nums)

	maxGap := math.MinInt
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] > maxGap {
			maxGap = nums[i+1] - nums[i]
		}
	}
	return maxGap
}
