package BFS

import "math"

func MaxDepthRecursion(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + int(math.Max(float64(MaxDepthRecursion(node.Left)), float64(MaxDepthRecursion(node.Right))))
}

func FindMaximumDepthBinaryTree(root *TreeNode) int {
	return MaxDepthRecursion(root)
}

func FindMaximumDepthBTBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)

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
		depth++
	}
	return depth
}
