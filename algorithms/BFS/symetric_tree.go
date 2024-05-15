package BFS

func SymmetricTree(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	if (root.Left == nil && root.Right != nil) || (root.Left != nil && root.Right == nil) {
		return false
	}

	queue := make([][]*TreeNode, 0)
	queue = append(queue, []*TreeNode{root.Left, root.Right})

	for len(queue) > 0 {
		top1, top2, tmpQueue := pop(queue)
		queue = tmpQueue

		if top1.Val != top2.Val {
			return false
		}

		if checkDiff(top1.Left, top2.Right) || checkDiff(top1.Right, top2.Left) {
			return false
		}

		if top1.Left != nil && top2.Right != nil {
			queue = append(queue, []*TreeNode{top1.Left, top2.Right})
		}

		if top1.Right != nil && top2.Left != nil {
			queue = append(queue, []*TreeNode{top1.Right, top2.Left})
		}
	}
	return true

}
