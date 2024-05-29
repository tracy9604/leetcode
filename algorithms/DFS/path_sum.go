package DFS

func hasPathSumDFS(node *TreeNode, sum, target int) bool {
	if node == nil {
		return false
	}

	sum += node.Val
	if node.Left == nil && node.Right == nil {
		return sum == target
	}
	leftAns := hasPathSumDFS(node.Left, sum, target)
	rightAns := hasPathSumDFS(node.Right, sum, target)
	return leftAns || rightAns
}

func HasPathSum(root *TreeNode, targetSum int) bool {
	return hasPathSumDFS(root, 0, targetSum)
}
