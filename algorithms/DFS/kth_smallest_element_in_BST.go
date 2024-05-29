package DFS

func inorderTraversal(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	ans := append(inorderTraversal(node.Left), node.Val)
	ans = append(ans, inorderTraversal(node.Right)...)
	return ans
}

func FindKthElement(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}

	data := inorderTraversal(root)
	if k > len(data) {
		return -1
	}
	return data[k-1]
}
