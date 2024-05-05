package binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func validBSTUtil(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if min != nil && root.Val <= min.Val {
		return false
	}

	if max != nil && root.Val >= max.Val {
		return false
	}

	return validBSTUtil(root.Left, min, root) && validBSTUtil(root.Right, root, max)
}

func isValidBSTUtil(root *TreeNode) bool {
	st := make([]*TreeNode, 0)
	var pred *TreeNode

	for root != nil || len(st) > 0 {
		for root != nil {
			st = append(st, root)
			root = root.Left
		}

		top := st[len(st)-1]
		st = st[:len(st)-1]

		root = top
		if pred != nil && root.Val <= pred.Val {
			return false
		}
		pred = root
		root = root.Right
	}
	return true
}

func ValidBST(root *TreeNode) bool {
	return validBSTUtil(root, nil, nil)
}
