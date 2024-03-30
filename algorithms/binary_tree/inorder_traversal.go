package binary_tree

func InorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	st := make([]*TreeNode, 0)
	curr := root

	for len(st) > 0 || curr != nil {
		for curr != nil {
			st = append(st, curr)
			curr = curr.Left
		}

		top := st[len(st)-1]
		st = st[:len(st)-1]

		res = append(res, top.Val)
		curr = top.Right
	}
	return res
}
