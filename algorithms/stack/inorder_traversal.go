package stack

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversalTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := append(InorderTraversalTree(root.Left), root.Val)
	return append(res, InorderTraversalTree(root.Right)...)
}

func popTreeNodeStack(stack []*TreeNode) (*TreeNode, []*TreeNode) {
	if len(stack) == 0 {
		return nil, stack
	}
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return top, stack
}

func DFSInOrderTraversalTree(node *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		top, tmpStack := popTreeNodeStack(stack)
		stack = tmpStack
		res = append(res, top.Val)
		node = top.Right
	}
	return res
}

func InorderTraversalTreeV2(root *TreeNode) []int {
	return DFSInOrderTraversalTree(root)
}
