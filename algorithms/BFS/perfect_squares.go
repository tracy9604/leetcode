package BFS

import "math"

func PerfectSquares(n int) int {
	if math.Sqrt(float64(n))-math.Floor(math.Sqrt(float64(n))) == 0 {
		return 1
	}

	if n <= 3 {
		return n
	}

	res := n

	for x := 1; x <= n; x++ {
		temp := x * x
		if temp > n {
			break
		}
		res = int(math.Min(float64(res), float64(1+PerfectSquares(n-temp))))
	}

	return res
}

func PerfectSquaresBFS(n int) int {
	visited := make(map[int]bool)
	queue := make([][]int, 0)

	ans := math.MaxInt
	queue = append(queue, []int{n, 0})

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if top[0] == 0 {
			ans = int(math.Min(float64(ans), float64(top[1])))
		}

		for x := 1; x*x <= top[0]; x++ {
			path := top[0] - x*x
			if _, found := visited[path]; path >= 0 && (!found || path == 0) {
				visited[path] = true
				queue = append(queue, []int{path, top[1] + 1})
			}
		}
	}
	return ans
}
