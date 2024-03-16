package queue

import (
	"math"
)

type square struct {
	num   int
	count int
}

func popIntQueue(queue []square) []square {
	if len(queue) == 0 {
		return queue
	}
	if len(queue) == 1 {
		return []square{}
	}
	return queue[1:]
}

func CheckPerfectSquares(n int) int {
	queue := make([]square, 0)
	visited := make(map[int]bool)

	start := square{num: n, count: 0}
	queue = append(queue, start)

	ans := math.MaxInt
	for len(queue) > 0 {
		i := 1
		curr := queue[0]
		queue = popIntQueue(queue)

		if curr.num == 0 {
			ans = int(math.Min(float64(ans), float64(curr.count)))
		}

		for i*i <= curr.num {
			path := curr.num - i*i
			_, found := visited[path]
			if path == 0 || (!found && path > 0) {
				queue = append(queue, square{num: path, count: curr.count + 1})
				visited[path] = true
			}
			i++
		}
	}
	return ans
}
