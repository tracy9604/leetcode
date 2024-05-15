package BFS

func popQueue(queue []*TreeNode) (*TreeNode, []*TreeNode) {
	if len(queue) == 0 {
		return nil, queue
	}
	if len(queue) == 1 {
		return queue[0], []*TreeNode{}
	}
	return queue[0], queue[1:]
}

func LevelOrderTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := make([]*TreeNode, 0)
	ans := make([][]int, 0)

	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		tmpAns := make([]int, 0)
		for i := 0; i < size; i++ {
			top, tmpQueue := popQueue(queue)
			queue = tmpQueue
			tmpAns = append(tmpAns, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}

		}
		if len(tmpAns) > 0 {
			ans = append(ans, tmpAns)
		}
	}
	return ans
}
