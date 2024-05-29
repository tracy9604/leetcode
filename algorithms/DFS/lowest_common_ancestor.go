package DFS

func lowestCommonAncestor(node *TreeNode, p, q *TreeNode) *TreeNode {
	if node == nil || node == p || node == q {
		return node
	}

	left := lowestCommonAncestor(node.Left, p, q)
	right := lowestCommonAncestor(node.Right, p, q)

	if left != nil && right != nil {
		return node
	}
	if left != nil {
		return left
	}
	return right
}

func FindLowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return lowestCommonAncestor(root, p, q)
}
