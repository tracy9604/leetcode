package BFS

func reverseList(list []int) {
	i, j := 0, len(list)-1
	for i < j {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}
}

func ZigZacLevelOrderTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	flag := 1 // 0: right, 1: left
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ans := make([][]int, 0)

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

		if flag == 0 {
			reverseList(tmpAns)
		}

		if len(tmpAns) > 0 {
			ans = append(ans, tmpAns)
		}

		flag ^= 1
	}

	return ans
}
