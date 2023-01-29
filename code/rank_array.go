package code

import "sort"

func ArrayRankTransform(arr []int) []int {
	result := make([]int, len(arr))
	arrMap := make(map[int]int)
	arrKeys := make([]int, 0)

	for _, item := range arr {
		if _, found := arrMap[item]; !found {
			arrMap[item]++
			arrKeys = append(arrKeys, item)
		}
	}

	sort.Ints(arrKeys)

	for idx, item := range arrKeys {
		arrMap[item] = idx
	}

	for idx, item := range arr {
		result[idx] = arrMap[item] + 1
	}

	return result
}
