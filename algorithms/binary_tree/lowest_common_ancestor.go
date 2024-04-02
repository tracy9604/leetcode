package binary_tree

func findPathFromRootToNode(root, node *TreeNode, path []*TreeNode) (bool, []*TreeNode) {
	if root == nil {
		return false, path
	}
	path = append(path, root)

	if root.Val == node.Val {
		return true, path
	}

	isValidLeftPath := false
	isValidRightPath := false
	tmpPath := make([]*TreeNode, 0)

	if root.Left != nil {
		isValidLeftPath, tmpPath = findPathFromRootToNode(root.Left, node, path)
		if isValidLeftPath {
			path = tmpPath
		}
	}
	if root.Right != nil {
		isValidRightPath, tmpPath = findPathFromRootToNode(root.Right, node, path)
		if isValidRightPath {
			path = tmpPath
		}
	}

	if isValidLeftPath || isValidRightPath {
		return true, path
	}

	path = path[:len(path)-1]

	return false, path
}

func FindLowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	validPathNodeP, pathP := findPathFromRootToNode(root, p, make([]*TreeNode, 0))
	validPathNodeQ, pathQ := findPathFromRootToNode(root, q, make([]*TreeNode, 0))

	if !validPathNodeP || !validPathNodeQ {
		return nil
	}

	i := 0
	var ancestor *TreeNode
	for i < len(pathP) && i < len(pathQ) {
		if pathP[i].Val != pathQ[i].Val {
			break
		}

		ancestor = pathP[i]
		i++
	}

	return ancestor
}
