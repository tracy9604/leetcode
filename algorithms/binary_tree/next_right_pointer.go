package binary_tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func popNodeQueue(queue []*Node) (*Node, []*Node) {
	if len(queue) == 0 {
		return nil, []*Node{}
	}
	if len(queue) == 1 {
		return queue[0], []*Node{}
	}
	front := queue[0]
	queue = queue[1:]
	return front, queue
}

func BFSNextRightPointers(node *Node) {
	if node == nil {
		return
	}

	queue := make([]*Node, 0)
	queue = append(queue, node)

	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]*Node, 0)

		for i := 0; i < size; i++ {
			front, tmpQueue := popNodeQueue(queue)
			queue = tmpQueue

			if len(tmp) > 0 {
				tmp[len(tmp)-1].Next = front
			}
			tmp = append(tmp, front)

			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		}
	}
}

func NextRightPointers(root *Node) *Node {
	BFSNextRightPointers(root)
	return root
}
