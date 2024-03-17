package stack

type Node struct {
	Val       int
	Neighbors []*Node
}

func DFSCloneGraph(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if val, found := visited[node]; found {
		return val
	}
	newNode := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
	visited[node] = newNode
	for _, item := range node.Neighbors {
		visited[node].Neighbors = append(visited[node].Neighbors, DFSCloneGraph(item, visited))
	}
	return newNode
}

func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]*Node)
	return DFSCloneGraph(node, visited)
}
