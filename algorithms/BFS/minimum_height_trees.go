package BFS

func FindMinimumHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make(map[int][]int)
	degree := make([]int, n)
	for _, edge := range edges {
		nodeA := edge[0]
		nodeB := edge[1]

		graph[nodeA] = append(graph[nodeA], nodeB)
		graph[nodeB] = append(graph[nodeB], nodeA)

		degree[nodeA]++
		degree[nodeB]++
	}

	queue := make([]int, 0)

	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}

	minHeights := make([]int, 0)

	for len(queue) > 0 {
		minHeights = make([]int, 0)
		size := len(queue)
		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]

			minHeights = append(minHeights, top)

			for _, neighbor := range graph[top] {
				degree[neighbor]--
				if degree[neighbor] == 1 {
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return minHeights
}

func minimumHeightTreesBFS(root int, edgesMap map[int][]int) int {
	queue := make([]int, 0)
	queue = append(queue, root)
	height := 0
	visited := make(map[int]bool)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]

			visited[top] = true

			if neighbors, found := edgesMap[top]; found && len(neighbors) > 0 {
				for _, neighbor := range neighbors {
					if _, check := visited[neighbor]; !check {
						queue = append(queue, neighbor)
					}
				}
			}
		}
		height++
	}
	return height
}
