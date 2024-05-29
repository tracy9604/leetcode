package DFS

func preorderDFS(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	ans := append([]int{node.Val}, preorderDFS(node.Left)...)
	ans = append(ans, preorderDFS(node.Right)...)
	return ans
}

func FlattenBinaryTreeToLinkedList(root *TreeNode) {
	if root == nil {
		return
	}

	data := preorderDFS(root)

	root.Left = nil
	prev := root
	for i := 1; i < len(data); i++ {
		node := &TreeNode{Val: data[i]}
		prev.Right = node
		prev = node
	}
}
