package BFS

type Node struct {
	Val       int
	Left      *Node
	Right     *Node
	Next      *Node
	Neighbors []*Node
}

func popNodeQueue(queue []*Node) (*Node, []*Node) {
	if len(queue) == 0 {
		return nil, queue
	}
	if len(queue) == 1 {
		return queue[0], []*Node{}
	}
	return queue[0], queue[1:]
}

func PopulatingNextRightPointers(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		var dummyNode *Node
		size := len(queue)

		for i := 0; i < size; i++ {
			top, tmpQueue := popNodeQueue(queue)
			queue = tmpQueue

			if dummyNode != nil {
				dummyNode.Next = top
			}
			dummyNode = top

			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}

		}
	}

	return root
}
