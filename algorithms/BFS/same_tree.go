package BFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CheckSameTreeDFS(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil {
		return p.Val == q.Val && CheckSameTreeDFS(p.Left, q.Left) && CheckSameTreeDFS(p.Right, q.Right)
	}

	return false
}

func pop(queue [][]*TreeNode) (*TreeNode, *TreeNode, [][]*TreeNode) {
	if len(queue) == 0 {
		return nil, nil, queue
	}
	if len(queue) == 1 {
		return queue[0][0], queue[0][1], [][]*TreeNode{}
	}
	return queue[0][0], queue[0][1], queue[1:]
}

func checkDiff(node1, node2 *TreeNode) bool {
	return (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil)
}

func CheckSameTreeBFS(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	queue := make([][]*TreeNode, 0)
	queue = append(queue, []*TreeNode{p, q})

	for len(queue) > 0 {
		top1, top2, tmpQueue := pop(queue)
		queue = tmpQueue

		if top1.Val != top2.Val {
			return false
		}

		if checkDiff(top1.Left, top2.Left) || checkDiff(top1.Right, top2.Right) {
			return false
		}

		if top1.Left != nil && top2.Left != nil {
			queue = append(queue, []*TreeNode{top1.Left, top2.Left})
		}
		if top1.Right != nil && top2.Right != nil {
			queue = append(queue, []*TreeNode{top1.Right, top2.Right})
		}
	}
	return true
}
