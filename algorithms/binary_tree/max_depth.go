package binary_tree

import "math"

func MaxDepth(root *TreeNode) int {
	return BottomUpMaxDepth(root)
}

func TopDownMaxDepth(root *TreeNode, answer, depth int) {
	if root == nil {
		return
	}

	// if node is leaf node
	if root.Left == nil && root.Right == nil {
		answer = int(math.Max(float64(answer), float64(depth)))
	}

	TopDownMaxDepth(root.Left, answer, depth+1)
	TopDownMaxDepth(root.Right, answer, depth+1)
}

func BottomUpMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxLeft := BottomUpMaxDepth(root.Left)
	maxRight := BottomUpMaxDepth(root.Right)
	return int(math.Max(float64(maxLeft), float64(maxRight))) + 1
}
