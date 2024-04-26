package linked_list

func ReverseLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func ReverseLinkedListV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head
	next := curr.Next

	for curr.Next != nil {
		curr.Next = next.Next
		next.Next = head
		head = next
		next = curr.Next
	}
	return head
}
