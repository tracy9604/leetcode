package contests_06_02

import (
	"fmt"
	"math"
)

func countDays(days int, meetings [][]int) int {
	i, j := 0, len(meetings)-1

	prev := make([]int, 0)

	for i <= j {
		meeting1 := meetings[i]
		meeting2 := meetings[j]

		start := int(math.Min(float64(meeting1[0]), float64(meeting2[0])))
		end := int(math.Max(float64(meeting1[1]), float64(meeting2[1])))

		if len(prev) == 0 {
			prev = []int{start, end}
		} else {
			tmpStart := int(math.Min(float64(start), float64(prev[0])))
			tmpEnd := int(math.Max(float64(end), float64(prev[1])))
			prev = []int{tmpStart, tmpEnd}
		}
		i++
		j--
	}

	return days - (prev[1] - prev[0] + 1)
}

func TestCountDays() {
	days := 10
	meetings := [][]int{
		{5, 7},
		{1, 3},
		{9, 10},
	}
	fmt.Println(countDays(days, meetings))
}
