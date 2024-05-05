package two_pointers_technique

func RemoveDuplicatedFromSortedList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyNode := &ListNode{Val: -101, Next: head}
	prev := dummyNode
	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Val != curr.Next.Val {
			prev = curr
			curr = curr.Next
			continue
		}

		for curr != nil && curr.Val == curr.Next.Val {
			curr = curr.Next
		}

		prev.Next = curr.Next
		curr = curr.Next
	}

	return dummyNode.Next
}
