package code

import "sort"

func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ {
		top := result[len(result)-1]

		if top[1] < intervals[i][0] {
			result = append(result, intervals[i])
		} else if top[1] < intervals[i][1] {
			top[1] = intervals[i][1]
			result = pop(result)
			result = append(result, top)
		}
	}
	return result
}

func pop(arr [][]int) [][]int {
	return arr[:len(arr)-1]
}
