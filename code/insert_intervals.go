package code

import "sort"

func InsertIntervals(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	intervals = append(intervals, newInterval)
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
