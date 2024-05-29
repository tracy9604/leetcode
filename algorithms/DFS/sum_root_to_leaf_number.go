package DFS

import (
	"strconv"
	"strings"
)

func sumDFS(node *TreeNode, path []string) int {
	if node == nil {
		return 0
	}

	path = append(path, strconv.Itoa(node.Val))
	if node.Left == nil && node.Right == nil {
		num, _ := strconv.Atoi(strings.Join(path, ""))
		return num
	}

	copyPath := make([]string, len(path))
	copy(copyPath, path)

	return sumDFS(node.Left, copyPath) + sumDFS(node.Right, copyPath)
}

func SumRootToLeafNumber(root *TreeNode) int {
	return sumDFS(root, []string{})
}
