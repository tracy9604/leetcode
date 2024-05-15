package BFS

import "math"

func FindMinDepthBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return 1 + FindMinDepthBinaryTree(root.Right)
	}

	if root.Right == nil {
		return 1 + FindMinDepthBinaryTree(root.Left)
	}

	return 1 + int(math.Min(float64(FindMinDepthBinaryTree(root.Left)), float64(FindMinDepthBinaryTree(root.Right))))
}

func FindMinDepthBinaryTreeBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	depth := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			top, tmpQueue := popQueue(queue)
			queue = tmpQueue

			if top.Left == nil && top.Right == nil {
				return depth
			}
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
