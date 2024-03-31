package binary_tree

func TopDownPathSum(root *TreeNode, sum, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		sum += root.Val
		return sum == targetSum
	}

	sum += root.Val
	ansLeft := TopDownPathSum(root.Left, sum, targetSum)
	ansRight := TopDownPathSum(root.Right, sum, targetSum)
	return ansLeft || ansRight
}

func PathSum(root *TreeNode, targetSum int) bool {
	return TopDownPathSum(root, 0, targetSum)
}
