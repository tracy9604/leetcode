package recursion

func ReserveLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	curr := head
	var prev, next *ListNode
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	head = prev
	return head
}

func ReserveLinkedListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	rest := ReserveLinkedListRecursion(head.Next)

	head.Next.Next = head

	head.Next = nil
	return rest
}
