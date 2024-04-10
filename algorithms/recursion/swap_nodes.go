package recursion

type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	remaining := head.Next.Next
	newHead := head.Next
	head.Next.Next = head
	head.Next = SwapNodes(remaining)
	return newHead
}

func SwapNodesIteratively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := head
	curr := head.Next

	head = curr

	for {
		next := curr.Next
		curr.Next = prev

		if next == nil || next.Next == nil {
			prev.Next = next
			break
		}

		prev.Next = next.Next
		prev = next
		curr = prev.Next
	}

	return head
}
