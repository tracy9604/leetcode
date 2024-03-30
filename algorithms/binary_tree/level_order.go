package binary_tree

func LevelOrderTraversal(root *TreeNode) [][]int {
	res := make([][]int, 0)

	height := findHeight(root)
	for i := 1; i <= height; i++ {
		res = append(res, getCurrentLevel(root, i))
	}
	return res
}

func getCurrentLevel(node *TreeNode, level int) []int {
	if node == nil {
		return []int{}
	}
	if level == 1 {
		return []int{node.Val}
	}
	return append(getCurrentLevel(node.Left, level-1), getCurrentLevel(node.Right, level-1)...)
}

func findHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	lHeight := findHeight(node.Left)
	rHeight := findHeight(node.Right)
	if lHeight > rHeight {
		return lHeight + 1
	}
	return rHeight + 1
}

func popQueue(queue []*TreeNode) (*TreeNode, []*TreeNode) {
	if len(queue) == 0 {
		return nil, []*TreeNode{}
	}
	if len(queue) == 1 {
		return queue[0], []*TreeNode{}
	}
	front := queue[0]
	queue = queue[1:]
	return front, queue
}

func BFSLevelOrderTraversal(root *TreeNode) [][]int {
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)

	if root == nil {
		return res
	}

	queue = append(queue, root)

	for len(queue) > 0 {
		tmpRes := make([]int, 0)
		size := len(queue)
		for i := 0; i < size; i++ {
			front, tmpQueue := popQueue(queue)
			queue = tmpQueue

			tmpRes = append(tmpRes, front.Val)

			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		}

		if len(tmpRes) > 0 {
			res = append(res, tmpRes)
		}

	}
	return res
}
