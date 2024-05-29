package BFS

func InvertBinaryTreeDFS(node *TreeNode) {
	if node == nil {
		return
	}
	node.Left, node.Right = node.Right, node.Left
	InvertBinaryTreeDFS(node.Left)
	InvertBinaryTreeDFS(node.Right)
}

func InvertBinaryTreeBFS(node *TreeNode) *TreeNode {
	if node == nil {
		return node
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, node)

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		top.Left, top.Right = top.Right, top.Left

		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}
	return node
}

func InvertBinaryTree(root *TreeNode) *TreeNode {
	InvertBinaryTreeDFS(root)
	return root
}
