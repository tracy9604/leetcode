package DFS

import (
	"strconv"
	"strings"
)

func binaryTreePaths(node *TreeNode, path []string) []string {
	if node == nil {
		return []string{}
	}
	path = append(path, strconv.Itoa(node.Val))
	if node.Left == nil && node.Right == nil {
		return []string{strings.Join(path, "->")}
	}
	return append(binaryTreePaths(node.Left, path), binaryTreePaths(node.Right, path)...)
}

func FindBinaryTreePath(root *TreeNode) []string {
	ans := binaryTreePaths(root, []string{})
	return ans
}
