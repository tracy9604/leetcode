package binary_tree

func searchInIdx(inorder []int, startIdx, endIdx, nodeVal int) int {
	i := 0
	for i = startIdx; i < endIdx+1; i++ {
		if inorder[i] == nodeVal {
			break
		}
	}
	return i
}

func buildTree1(inorder []int, postorder []int, startIdx, endIdx int, pIdx []int) *TreeNode {
	if startIdx > endIdx {
		return nil
	}

	node := &TreeNode{
		Val: postorder[pIdx[0]],
	}
	pIdx[0]--

	// node has no children, return
	if startIdx == endIdx {
		return node
	}

	inIdx := searchInIdx(inorder, startIdx, endIdx, node.Val)
	node.Right = buildTree1(inorder, postorder, inIdx+1, endIdx, pIdx)
	node.Left = buildTree1(inorder, postorder, startIdx, inIdx-1, pIdx)
	return node
}

func ConstructBinaryTree1(inorder []int, postorder []int) *TreeNode {
	pIdx := make([]int, 1)
	pIdx[0] = len(postorder) - 1
	return buildTree1(inorder, postorder, 0, len(inorder)-1, pIdx)
}
