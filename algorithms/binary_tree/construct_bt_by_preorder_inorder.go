package binary_tree

func buildTree2(inorder, preorder []int, startIdx, endIdx int, pIdx []int) *TreeNode {
	if startIdx > endIdx || pIdx[0] >= len(preorder) {
		return nil
	}

	node := &TreeNode{
		Val: preorder[pIdx[0]],
	}
	pIdx[0]++

	if startIdx == endIdx {
		return node
	}

	inIdx := searchInIdx(inorder, startIdx, endIdx, node.Val)
	node.Left = buildTree2(inorder, preorder, startIdx, inIdx-1, pIdx)
	node.Right = buildTree2(inorder, preorder, inIdx+1, endIdx, pIdx)

	return node
}

func ConstructBinaryTree2(preorder []int, inorder []int) *TreeNode {
	return buildTree2(inorder, preorder, 0, len(inorder)-1, []int{0})
}
