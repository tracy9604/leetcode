package DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode, prev, first, second *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left, prev, first, second)

	if prev != nil && prev.Val > root.Val {
		if first == nil {
			first = prev
		}
		second = root
	}
	prev = root
	inorder(root.Right, prev, first, second)
}

func RecoverTree(root *TreeNode) {
	var prev, first, second *TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)

		if prev != nil && prev.Val > node.Val {
			if first == nil {
				first = prev
			}
			second = node
		}
		prev = node
		dfs(node.Right)
	}

	dfs(root)

	tmp := first.Val
	first.Val = second.Val
	second.Val = tmp
}
