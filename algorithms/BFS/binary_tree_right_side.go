package BFS

func FindRightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ans := make([]int, 0)

	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]int, 0)

		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]
			tmp = append(tmp, top.Val)

			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		if len(tmp) > 0 {
			ans = append(ans, tmp[len(tmp)-1])
		}
	}
	return ans
}
