package DFS

type pathStruct struct {
	data []int
}

func pathSum2DFS(node *TreeNode, sum, targetSum int, path []int) [][]int {
	if node == nil {
		return [][]int{}
	}
	path = append(path, node.Val)
	sum += node.Val
	if node.Left == nil && node.Right == nil && sum == targetSum {
		return [][]int{path}
	}

	newGroupOfNodes := make([]int, len(path))
	copy(newGroupOfNodes, path)

	leftAns := pathSum2DFS(node.Left, sum, targetSum, newGroupOfNodes)
	rightAns := pathSum2DFS(node.Right, sum, targetSum, newGroupOfNodes)

	return append(leftAns, rightAns...)
}

func PathSum2(root *TreeNode, targetSum int) [][]int {
	return pathSum2DFS(root, 0, targetSum, []int{})
}
