package BFS

func CoinChange(coins []int, amount int) int {
	queue := make([]int, 0)
	queue = append(queue, amount)

	visited := make(map[int]bool)
	ans := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]

			if top == 0 {
				return ans
			}

			if _, found := visited[top]; found || top < 0 {
				continue
			}

			visited[top] = true

			for j := 0; j < len(coins); j++ {
				queue = append(queue, top-coins[j])
			}
		}
		ans++
	}
	return -1
}
