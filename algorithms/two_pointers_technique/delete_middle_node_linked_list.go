package two_pointers_technique

func DeleteMiddleNodeLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummyNode := &ListNode{}
	dummyNode.Next = head
	slow, fast := dummyNode, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummyNode.Next
}
