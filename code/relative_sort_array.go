package code

import "sort"

func RelativeSortArray(arr1 []int, arr2 []int) []int {
	arr1Map := make(map[int]int)
	result := make([]int, len(arr1))
	idx := 0
	for _, item := range arr1 {
		arr1Map[item]++
	}

	for _, item := range arr2 {
		count, _ := arr1Map[item]
		for i := 0; i < count; i++ {
			result[i+idx] = item
		}

		arr1Map[item] = 0
		idx += count
	}

	remainingIdx := idx
	for _, item := range arr1 {
		if count, _ := arr1Map[item]; count > 0 {
			result[idx] = item
			idx++
		}
	}

	sort.SliceStable(result, func(i, j int) bool {
		return i >= remainingIdx && j >= remainingIdx && result[i] <= result[j]
	})

	return result
}
