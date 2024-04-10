package recursion

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SearchInBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val < val {
		return SearchInBST(root.Right, val)
	}

	return SearchInBST(root.Left, val)
}
