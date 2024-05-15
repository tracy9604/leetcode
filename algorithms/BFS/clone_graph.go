package BFS

func CloneGraphRecursion(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return node
	}

	if val, found := visited[node]; found {
		return val
	}

	newNode := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
	visited[node] = newNode
	for _, neighbor := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, CloneGraphRecursion(neighbor, visited))
	}

	return newNode
}

func CloneGraphBFS(node *Node) *Node {
	if node == nil {
		return node
	}

	queue := make([]*Node, 0)
	queue = append(queue, node)
	nodes := make([]*Node, 0)
	copyNodes := make(map[*Node]*Node)
	visited := make(map[*Node]bool)

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			if _, found := visited[curr]; found {
				continue
			}

			visited[curr] = true
			nodes = append(nodes, curr)
			copyNodes[curr] = &Node{Val: curr.Val}

			queue = append(queue, curr.Neighbors...)
		}
	}

	for _, nodeItem := range nodes {
		copyNodes[nodeItem].Neighbors = make([]*Node, 0)
		for _, cNode := range nodeItem.Neighbors {
			copyNodes[nodeItem].Neighbors = append(copyNodes[nodeItem].Neighbors, copyNodes[cNode])
		}
	}

	return copyNodes[node]
}

func CloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	return CloneGraphRecursion(node, map[*Node]*Node{})
}
