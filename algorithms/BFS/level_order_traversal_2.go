package BFS

func LevelOrderTraversal2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ans := make([][]int, 0)

	for len(queue) > 0 {
		size := len(queue)
		tmpAns := make([]int, 0)
		for i := 0; i < size; i++ {
			top, tmpQueue := popQueue(queue)
			queue = tmpQueue

			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		if len(tmpAns) > 0 {
			ans = append([][]int{tmpAns}, ans...)
		}
	}

	if len(ans) < 2 {
		return ans
	}

	return ans
}
