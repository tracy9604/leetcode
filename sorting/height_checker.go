package sorting

import (
	"sort"
)

func HeightChecker(heights []int) int {
	tmpHeights := make([]int, 0)
	tmpHeights = append(tmpHeights, heights...)
	sort.Ints(heights)

	count := 0
	for idx, item := range tmpHeights {
		if item != heights[idx] {
			count++
		}
	}

	return count
}

func HeightChecker2(heights []int) int {
	heightsMap := make(map[int]int)
	for i := 0; i < len(heights); i++ {
		heightsMap[heights[i]]++
	}

	result := 0
	curHeight := 0
	for i := 0; i < len(heights); i++ {
		for heightsMap[curHeight] == 0 {
			curHeight++
		}

		if curHeight != heights[i] {
			result++
		}
		heightsMap[curHeight]--
	}

	return result
}

func HeightChecker3(heights []int) int {
	heightsMap := make(map[int]int)
	max := 0

	for _, item := range heights {
		if item > max {
			max = item
		}
		heightsMap[item]++
	}

	for i := 1; i <= max; i++ {
		heightsMap[i] += heightsMap[i-1]
	}

	result := 0
	for i := 0; i < len(heights); i++ {
		if heights[heightsMap[heights[i]]-1] != heights[i] {
			result++
		}
		heightsMap[heights[i]]--
	}

	return result
}
