package binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	res = append(res, root.Val)
	res = append(res, PreOrderTraversal(root.Left)...)
	res = append(res, PreOrderTraversal(root.Right)...)
	return res
}

func PreOrderTraversalStack(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	st := make([]*TreeNode, 0)

	curr := root

	for len(st) > 0 || curr != nil {
		for curr != nil {
			st = append(st, curr)
			top := st[len(st)-1]
			res = append(res, top.Val)
			curr = curr.Left
		}

		curr = st[len(st)-1]
		st = st[:len(st)-1]

		curr = curr.Right
	}
	return res
}
