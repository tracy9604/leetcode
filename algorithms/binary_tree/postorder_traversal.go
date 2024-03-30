package binary_tree

func PostOrderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	st1 := make([]*TreeNode, 0)
	st2 := make([]*TreeNode, 0)
	if root == nil {
		return res
	}
	st1 = append(st1, root)

	for len(st1) > 0 {
		top := st1[len(st1)-1]
		st1 = st1[:len(st1)-1]
		st2 = append(st2, top)

		if top.Left != nil {
			st1 = append(st1, top.Left)
		}
		if top.Right != nil {
			st1 = append(st1, top.Right)
		}
	}

	for len(st2) > 0 {
		top := st2[len(st2)-1]
		st2 = st2[:len(st2)-1]
		res = append(res, top.Val)
	}

	return res
}
