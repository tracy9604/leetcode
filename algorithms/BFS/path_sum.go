package BFS

func PathSumRecursion(node *TreeNode, sum, targetSum int) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil {
		sum += node.Val
		return sum == targetSum
	}

	sum += node.Val
	ansLeft := PathSumRecursion(node.Left, sum, targetSum)
	ansRight := PathSumRecursion(node.Right, sum, targetSum)
	return ansLeft || ansRight
}

func PathSum(root *TreeNode, targetSum int) bool {
	return PathSumRecursion(root, 0, targetSum)
}

func PathSumBFS(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	sums := []int{targetSum}

	for len(queue) > 0 {
		size := len(queue)
		tmpSums := make([]int, 0)

		for i := 0; i < size; i++ {
			top, tmpQueue := popQueue(queue)
			queue = tmpQueue

			v := sums[i] - top.Val
			if top.Left == nil && top.Right == nil && v == 0 {
				return true
			}
			if top.Left != nil {
				queue = append(queue, top.Left)
				tmpSums = append(tmpSums, v)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
				tmpSums = append(tmpSums, v)
			}
		}
		sums = tmpSums
	}
	return false
}
