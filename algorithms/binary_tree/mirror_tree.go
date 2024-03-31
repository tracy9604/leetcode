package binary_tree

func isMirror(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	// root1 is mirror of root2 if
	// root1 and root2 has same value
	// left subtree of root1 is right subtree of root2
	// right subtree of root1 is left subtree of root2

	if root1 != nil && root2 != nil && root1.Val == root2.Val && isMirror(root1.Left, root2.Right) && isMirror(root1.Right, root2.Left) {
		return true
	}
	return false
}

func SymmetricTree(root *TreeNode) bool {
	return isMirror(root, root)
}
